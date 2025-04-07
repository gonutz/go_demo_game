package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/gonutz/ds"
	"github.com/hajimehoshi/go-mp3"
	"github.com/jfreymuth/oggvorbis"
)

const (
	// 4096 samples is about 93 ms at 44100 Hz.
	soundWriteAheadSamples = 4096
	soundWriteAheadSize    = 4 * soundWriteAheadSamples
)

type soundHandle int

const invalidSoundHandle soundHandle = 0

type soundSystem struct {
	dsound *ds.DirectSound
	// mixBuffer is our hardware sound buffer that gets played in a loop. We
	// regularly update its contents at the position that will be played next.
	mixBuffer     *ds.Buffer
	mixBufferSize int
	// writeAheadBuffer and writeAheadMixBuffer are really temporary buffers
	// used in the main update loop. We keep them here to not allocate them
	// anew every frame.
	writeAheadBuffer    [soundWriteAheadSamples]soundSample
	writeAheadMixBuffer [soundWriteAheadSamples]mixSample
	// lastWritePos is the offset into the mixBuffer where we last wrote to.
	// This way we can calculate how many samples have been played since the
	// last update.
	lastWritePos int
	// loadedSounds caches the raw sound data for all sound files loaded from
	// disk.
	loadedSounds map[string][]byte
	// playingSounds keeps the currently playing sound states. Once a sound is
	// finished playing, it is removed from this queue.
	playingSounds []soundState
	// nextHandle is an ever increasing ID number for all the sounds that are
	// played over time.
	nextHandle soundHandle
	queue      []consecutiveSounds
}

type soundState struct {
	handle    soundHandle
	samples   []soundSample
	pos       float64
	lastSpeed float64
	speed     float64
	looping   bool
	queued    bool
}

type consecutiveSounds [2]soundHandle

func (s *soundState) isOver() bool {
	return !s.looping && s.pos >= float64(len(s.samples)-1)
}

// soundSample is the final raw data that gets send to the sound card. We use a
// sample rate of 44100 Hz, stereo sound, so 2 samples (left and right), with 2
// bytes per channel (int16).
type soundSample struct {
	channels [2]int16
}

// mixSample is used for our mixing of multiple sounds. Sounds waves simply add
// on top of each other. If we did this in the int16 space, we would soon be
// out of range. We use int32 for the temporary summation of sound samples and
// convert back down to int16 when we get ready to send it to the sound card.
type mixSample struct {
	channels [2]int32
}

func initSoundSystem(window ds.HWND) (*soundSystem, error) {
	dsound, err := ds.Create(nil)
	if err != nil {
		return nil, err
	}

	// We use the cooperation level "normal" which means that we are restricted
	// to using 44100 Hz, 2 channel, int16 samples. That is what we set our
	// sound back buffer to.
	if err := dsound.SetCooperativeLevel(window, ds.SCL_NORMAL); err != nil {
		dsound.Release()
		return nil, err
	}

	soundFormat := ds.WAVEFORMATEX{
		FormatTag:     ds.WAVE_FORMAT_PCM,
		Channels:      2,
		SamplesPerSec: 44100,
		BitsPerSample: 16,
	}
	soundFormat.BlockAlign =
		(soundFormat.Channels * soundFormat.BitsPerSample) / 8
	soundFormat.AvgBytesPerSec =
		soundFormat.SamplesPerSec * uint32(soundFormat.BlockAlign)

	// Reserve 2 seconds worth of audio buffer.
	bufferSize := 2 * soundFormat.AvgBytesPerSec

	// The sound buffer should always hold whole 4-byte samples (2 channels, 2
	// bytes per sample), so it must be divisible by 4.
	for bufferSize%4 != 0 {
		bufferSize++
	}

	buffer, err := dsound.CreateSoundBuffer(ds.BUFFERDESC{
		Flags:       ds.BCAPS_GETCURRENTPOSITION2 | ds.BCAPS_GLOBALFOCUS,
		BufferBytes: bufferSize,
		WfxFormat:   &soundFormat,
	})
	if err != nil {
		dsound.Release()
		return nil, err
	}

	// Initialze the output buffer to silence (all 0).
	soundMem, err := buffer.Lock(0, bufferSize, 0)
	if err != nil {
		buffer.Release()
		dsound.Release()
		return nil, err
	}
	soundMem.Write(0, make([]byte, soundMem.Size()))
	if err := buffer.Unlock(soundMem); err != nil {
		buffer.Release()
		dsound.Release()
		return nil, err
	}

	// Start playing the buffer in an infinite loop. We will write into it at
	// its current play position every frame, updating the audible sound as we
	// go.
	if err := buffer.Play(0, ds.BPLAY_LOOPING); err != nil {
		buffer.Release()
		dsound.Release()
		return nil, err
	}

	return &soundSystem{
		dsound:        dsound,
		mixBuffer:     buffer,
		mixBufferSize: int(bufferSize),
		loadedSounds:  map[string][]byte{},
		nextHandle:    1,
	}, nil
}

func (s *soundSystem) close() {
	s.mixBuffer.Stop()
	s.mixBuffer.Release()
	s.dsound.Release()
}

func (s *soundSystem) stop(handle soundHandle) error {
	for i := range s.playingSounds {
		if handle == s.playingSounds[i].handle {
			s.playingSounds = append(s.playingSounds[:i], s.playingSounds[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("cannot stop unknown sound handle")
}

func (s *soundSystem) setSpeed(handle soundHandle, speed float64) error {
	for i := range s.playingSounds {
		if handle == s.playingSounds[i].handle {
			s.playingSounds[i].speed = speed
			return nil
		}
	}
	return fmt.Errorf("cannot set speed on unknown sound handle")
}

func (s *soundSystem) update() error {
	for i := range s.writeAheadMixBuffer {
		for c := range s.writeAheadMixBuffer[i].channels {
			s.writeAheadMixBuffer[i].channels[c] = 0
		}
	}

	mem, err := s.mixBuffer.Lock(0, soundWriteAheadSize, ds.BLOCK_FROMWRITECURSOR)
	if err != nil {
		return err
	}

	_, write, err := s.mixBuffer.GetCurrentPosition()
	if err != nil {
		return err
	}
	writePos := int(write)

	// Calculate how many samples were played since the last update. Combining
	// the number of bytes played with the known last sound speed we can update
	// our current sound position and continue playing the sound from there at
	// the updated current sound speed.
	playedSamples := s.writeSampleDist(s.lastWritePos, writePos)

	for i := range s.playingSounds {
		sound := &s.playingSounds[i]

		if sound.queued {
			continue
		}

		sound.pos += float64(playedSamples) * sound.lastSpeed
		if sound.looping {
			sound.pos = wrapSoundPos(sound.pos, len(sound.samples))
		}

		for i := range s.writeAheadMixBuffer {
			pos := sound.pos + float64(i)*sound.speed
			if sound.looping {
				pos = wrapSoundPos(pos, len(sound.samples))
			}
			j := round(pos)
			if 0 <= j && j < len(sound.samples) {
				for c := range s.writeAheadMixBuffer[i].channels {
					s.writeAheadMixBuffer[i].channels[c] +=
						int32(sound.samples[j].channels[c])
				}
			}
		}
	}

	for i := range s.writeAheadBuffer {
		for c := range s.writeAheadBuffer[i].channels {
			x := s.writeAheadMixBuffer[i].channels[c]
			if x > 32767 {
				x = 32767
			}
			if x < -32768 {
				x = -32768
			}
			s.writeAheadBuffer[i].channels[c] = int16(x)
		}
	}
	mem.WriteRaw(
		0, unsafe.Pointer(&s.writeAheadBuffer[0]), len(s.writeAheadBuffer)*4)

	if err := s.mixBuffer.Unlock(mem); err != nil {
		return err
	}

	// Remove all sounds that are over.
	n := 0
	for i := range s.playingSounds {
		if s.playingSounds[i].isOver() {
			queueN := 0
			for _, q := range s.queue {
				if q[0] == s.playingSounds[i].handle {
					if follow := s.soundFromHandle(q[1]); follow != nil {
						follow.queued = false
					}
				} else {
					s.queue[queueN] = q
					queueN++
				}
			}
		} else {
			s.playingSounds[n] = s.playingSounds[i]
			n++
		}
	}
	s.playingSounds = s.playingSounds[:n]

	for i := range s.playingSounds {
		s.playingSounds[i].lastSpeed = s.playingSounds[i].speed
	}

	s.lastWritePos = writePos

	return nil
}

func (s *soundSystem) soundFromHandle(handle soundHandle) *soundState {
	for i := range s.playingSounds {
		if handle == s.playingSounds[i].handle {
			return &s.playingSounds[i]
		}
	}
	return nil
}

func wrapSoundPos(pos float64, sampleCount int) float64 {
	n := float64(sampleCount) - 1
	for pos < 0 {
		pos += n
	}
	for pos > n {
		pos -= n
	}
	return pos
}

func round(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
}

func (s *soundSystem) writeSampleDist(a, b int) int {
	d := b - a
	if d < 0 {
		d = s.mixBufferSize - a + b
	}
	if d%4 != 0 {
		panic("why does the sound card play partial samples?")
	}
	return d / 4
}

func (s *soundSystem) play(path string) (soundHandle, error) {
	return s.playLoopingAndQueued(path, false, false)
}

func (s *soundSystem) loop(path string) (soundHandle, error) {
	return s.playLoopingAndQueued(path, true, false)
}

func (s *soundSystem) queueLoopAfter(atEndOf soundHandle, path string) (soundHandle, error) {
	handle, err := s.playLoopingAndQueued(path, true, true)
	if err != nil {
		return invalidSoundHandle, nil
	}
	s.queue = append(s.queue, consecutiveSounds{atEndOf, handle})
	return handle, nil
}

func (s *soundSystem) preload(path string) error {
	_, err := s.loadRawSamples(path)
	return err
}

func (s *soundSystem) playLoopingAndQueued(path string, looping, queued bool) (soundHandle, error) {
	raw, err := s.loadRawSamples(path)
	if err != nil {
		return invalidSoundHandle, err
	}

	// We read raw bytes above but we know that a single sound sample consists
	// of two int16, one for the left and one for the right channel. This makes
	// 4 bytes, so we cast the raw sound data to an array of 4-byte items (we
	// chose uint32).
	// We can index this array to get samples to pass to the sound card.
	samples := unsafe.Slice((*soundSample)(unsafe.Pointer(&raw[0])), len(raw)/4)

	handle := s.nextHandle
	s.nextHandle++

	s.playingSounds = append(s.playingSounds, soundState{
		handle:  handle,
		samples: samples,
		speed:   1,
		looping: looping,
		queued:  queued,
	})

	return handle, nil
}

func (s *soundSystem) loadRawSamples(path string) ([]byte, error) {
	if samples, ok := s.loadedSounds[path]; ok {
		return samples, nil
	}

	soundFile, err := assetFiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rawSoundData []byte

	if strings.HasSuffix(path, ".raw") {
		rawSoundData = soundFile
	} else if strings.HasSuffix(path, ".ogg") {
		data, format, err := oggvorbis.ReadAll(bytes.NewReader(soundFile))
		if err != nil {
			return nil, err
		}
		if format.SampleRate != 44100 {
			return nil, fmt.Errorf("we expect ogg files to be 44100 Hz")
		}
		if format.Channels != 2 {
			return nil, fmt.Errorf("we expect ogg files to have 2 channels")
		}
		rawSoundData = make([]byte, len(data)*2)
		for i := range data {
			j := 2 * i
			sample := int16(data[i] * 32767)
			*(*int16)(unsafe.Pointer(&rawSoundData[j])) = sample
		}
	} else if strings.HasSuffix(path, ".mp3") {
		decoder, err := mp3.NewDecoder(bytes.NewReader(soundFile))
		if err != nil {
			return nil, err
		}

		if decoder.SampleRate() != 44100 {
			return nil, fmt.Errorf("we expect mp3 files to be 44100 Hz")
		}

		rawSoundData, err = io.ReadAll(decoder)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unknown file extension for %q", path)
	}

	s.loadedSounds[path] = rawSoundData

	return s.loadedSounds[path], nil
}

package ds

import (
	"syscall"
	"unsafe"
)

// Buffer is used to manage sound buffers.
// To create a Buffer use DirectSound.CreateSoundBuffer.
// Not all methods of Buffer are valid for primary buffers. For example,
// SetCurrentPosition will fail.
type Buffer struct {
	vtbl *bufferVtbl
}

type bufferVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetCaps            uintptr
	GetCurrentPosition uintptr
	GetFormat          uintptr
	GetVolume          uintptr
	GetPan             uintptr
	GetFrequency       uintptr
	GetStatus          uintptr
	Initialize         uintptr
	Lock               uintptr
	Play               uintptr
	SetCurrentPosition uintptr
	SetFormat          uintptr
	SetVolume          uintptr
	SetPan             uintptr
	SetFrequency       uintptr
	Stop               uintptr
	Unlock             uintptr
	Restore            uintptr
	SetFX              uintptr
	GetObjectInPath    uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Buffer) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.AddRef,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *Buffer) Release() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// GetCaps retrieves the capabilities of the buffer object.
func (obj *Buffer) GetCaps() (caps BCAPS, err Error) {
	caps.Size = uint32(unsafe.Sizeof(caps))
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetCaps,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&caps)),
	)
	err = toErr(ret)
	return
}

// GetCurrentPosition retrieves the position of the play and write cursors in
// the sound buffer.
func (obj *Buffer) GetCurrentPosition() (playCursor, writeCursor uint32, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetCurrentPosition,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&playCursor)),
		uintptr(unsafe.Pointer(&writeCursor)),
	)
	err = toErr(ret)
	return
}

// GetFormat retrieves a description of the format of the sound data in the
// buffer, or the buffer size needed to retrieve the format description.
func (obj *Buffer) GetFormat() (format WAVEFORMATEX, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetFormat,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&format)),
		unsafe.Sizeof(format),
		0,
	)
	err = toErr(ret)
	return
}

// GetFormatExtensible retrieves a description of the format of the sound data
// in the buffer, or the buffer size needed to retrieve the format description.
func (obj *Buffer) GetFormatExtensible() (format WAVEFORMATEXTENSIBLE, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetFormat,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&format)),
		unsafe.Sizeof(format),
		0,
	)
	err = toErr(ret)
	return
}

// GetVolume retrieves the attenuation of the sound.
func (obj *Buffer) GetVolume() (volume int32, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetVolume,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&volume)),
	)
	err = toErr(ret)
	return
}

// GetPan retrieves the relative volume of the left and right audio channels.
func (obj *Buffer) GetPan() (pan int32, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetPan,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&pan)),
	)
	err = toErr(ret)
	return
}

// GetFrequency retrieves the frequency, in samples per second, at which the
// buffer is playing.
func (obj *Buffer) GetFrequency() (freq uint32, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetFrequency,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&freq)),
	)
	err = toErr(ret)
	return
}

// GetStatus retrieves the status of the sound buffer.
func (obj *Buffer) GetStatus() (status uint32, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetStatus,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&status)),
	)
	err = toErr(ret)
	return
}

// Initialize initializes a sound buffer object if it has not yet been
// initialized.
func (obj *Buffer) Initialize(ds *DirectSound, desc *BUFFERDESC) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Initialize,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ds)),
		uintptr(unsafe.Pointer(desc)),
	)
	return toErr(ret)
}

// Lock readies all or part of the buffer for a data write and returns pointers
// to which data can be written.
func (obj *Buffer) Lock(offset, bytes, flags uint32) (mem BufferMemory, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Lock,
		uintptr(unsafe.Pointer(obj)),
		uintptr(offset),
		uintptr(bytes),
		uintptr(unsafe.Pointer(&mem.dataPtr1)),
		uintptr(unsafe.Pointer(&mem.byteCount1)),
		uintptr(unsafe.Pointer(&mem.dataPtr2)),
		uintptr(unsafe.Pointer(&mem.byteCount2)),
		uintptr(flags),
	)
	err = toErr(ret)
	return
}

// Unlock releases a locked sound buffer.
func (obj *Buffer) Unlock(mem BufferMemory) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Unlock,
		uintptr(unsafe.Pointer(obj)),
		mem.dataPtr1,
		uintptr(mem.written1),
		mem.dataPtr2,
		uintptr(mem.written2),
	)
	return toErr(ret)
}

// BufferMemory wraps the conceptual ring buffer of a locked DirectSound buffer.
// This buffer is write-only.
type BufferMemory struct {
	dataPtr1, dataPtr2     uintptr
	byteCount1, byteCount2 uint32
	written1, written2     uint32
}

// Size returns the number of bytes in the locked buffer part. You can Write up
// to Size() bytes into this BufferMemory.
func (m *BufferMemory) Size() uint32 {
	return m.byteCount1 + m.byteCount2
}

// WriteRaw does the same as Write but accepts any data type in the form of a
// pointer and the total data size in bytes.
func (m *BufferMemory) WriteRaw(at uint32, data unsafe.Pointer, sizeInBytes int) {
	samples := unsafe.Slice((*byte)(data), sizeInBytes)
	m.Write(at, samples)
}

// Write copies the given data to the locked buffer memory.
// at is the relative offset in the locked area to start writing the data to.
// Use at=0 and data of length Size() to fill the whole buffer.
// Note that Write handles the ring buffer specifics of the DirectSound buffer
// internally.
func (m *BufferMemory) Write(at uint32, data []byte) {
	size := m.Size()
	if at >= size {
		return
	}
	maxDataLen := size - at
	if uint32(len(data)) > maxDataLen {
		data = data[:maxDataLen]
	}
	if at < m.byteCount1 {
		copyCount := uint32(len(data))
		if at+copyCount > m.byteCount1 {
			copyCount = m.byteCount1 - at
		}
		toCopy := data[:copyCount]
		// copy toCopy to dataPtr1+at
		dest := m.dataPtr1 + uintptr(at)
		for _, b := range toCopy {
			*((*byte)(unsafe.Pointer(dest))) = b
			dest++
		}
		// the rest of the data must go to the start of part 2
		data = data[copyCount:]
		at = 0
	}
	// at this point, data contains only those bytes that must go into part 2
	// of the ring buffer
	if len(data) > 0 {
		// copy data to dataPtr2+at
		dest := m.dataPtr2 + uintptr(at)
		for _, b := range data {
			*((*byte)(unsafe.Pointer(dest))) = b
			dest++
		}
	}
}

// Play causes the sound buffer to play, starting at the play cursor.
func (obj *Buffer) Play(priority, flags uint32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Play,
		uintptr(unsafe.Pointer(obj)),
		0,
		uintptr(priority),
		uintptr(flags),
	)
	return toErr(ret)
}

// SetCurrentPosition sets the position of the play cursor, which is the point
// at which the next byte of data is read from the buffer.
func (obj *Buffer) SetCurrentPosition(newPlayCursor uint32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetCurrentPosition,
		uintptr(unsafe.Pointer(obj)),
		uintptr(newPlayCursor),
	)
	return toErr(ret)
}

// SetFormat sets the format of the primary buffer. Whenever this application
// has the input focus, DirectSound will set the primary buffer to the specified
// format.
func (obj *Buffer) SetFormat(format WAVEFORMATEX) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetFormat,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&format)),
	)
	return toErr(ret)
}

// SetVolume sets the attenuation of the sound.
func (obj *Buffer) SetVolume(volume int32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetVolume,
		uintptr(unsafe.Pointer(obj)),
		uintptr(volume),
	)
	return toErr(ret)
}

// SetPan sets the relative volume of the left and right channels.
func (obj *Buffer) SetPan(pan int32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetPan,
		uintptr(unsafe.Pointer(obj)),
		uintptr(pan),
	)
	return toErr(ret)
}

// SetFrequency sets the frequency at which the audio samples are played.
func (obj *Buffer) SetFrequency(freq uint32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetFrequency,
		uintptr(unsafe.Pointer(obj)),
		uintptr(freq),
	)
	return toErr(ret)
}

// Stop causes the sound buffer to stop playing.
func (obj *Buffer) Stop() Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Stop,
		uintptr(unsafe.Pointer(obj)),
	)
	return toErr(ret)
}

// Restore restores the memory allocation for a lost sound buffer.
func (obj *Buffer) Restore() Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Restore,
		uintptr(unsafe.Pointer(obj)),
	)
	return toErr(ret)
}

// SetFX enables effects on a buffer.
// For this function to succeed, CoInitialize must have been called.
// Additionally, the buffer must not be playing or locked.
func (obj *Buffer) SetFX(
	effectsCount uint32,
	desc *EFFECTDESC,
) (resultCodes []uint32, err Error) {
	resultCodes = make([]uint32, effectsCount)
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Restore,
		uintptr(unsafe.Pointer(obj)),
		uintptr(effectsCount),
		uintptr(unsafe.Pointer(desc)),
		uintptr(unsafe.Pointer(&resultCodes[0])),
	)
	err = toErr(ret)
	return
}

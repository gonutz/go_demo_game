package ds

// VECTOR describes a 3D vector.
type VECTOR struct {
	X float32
	Y float32
	Z float32
}

// BUFFER3D contains all information necessary to uniquely describe the
// location, orientation, and motion of a 3D sound buffer.
type BUFFER3D struct {
	Size              uint32
	Position          VECTOR
	Velocity          VECTOR
	InsideConeAngle   uint32
	OutsideConeAngle  uint32
	ConeOrientation   VECTOR
	ConeOutsideVolume int32
	MinDistance       float32
	MaxDistance       float32
	Mode              uint32
}

// LISTENER3D describes the 3D world parameters and position of the listener.
type LISTENER3D struct {
	Size           uint32
	Position       VECTOR
	Velocity       VECTOR
	OrientFront    VECTOR
	OrientTop      VECTOR
	DistanceFactor float32
	RolloffFactor  float32
	DopplerFactor  float32
}

// BCAPS describes the capabilities of a DirectSound buffer object.
type BCAPS struct {
	Size               uint32
	Flags              uint32
	BufferBytes        uint32
	UnlockTransferRate uint32
	PlayCpuOverhead    uint32
}

type (
	// HANDLE describes a handle.
	HANDLE uintptr
	// HWND is a window handle.
	HWND HANDLE
)

// BPOSITIONNOTIFY describes a notification position. It is used by
// Notify.SetNotificationPositions.
type BPOSITIONNOTIFY struct {
	Offset      uint32
	EventNotify HANDLE
}

// BUFFERDESC describes the characteristics of a buffer.
type BUFFERDESC struct {
	Size        uint32
	Flags       uint32
	BufferBytes uint32
	Reserved    uint32
	WfxFormat   *WAVEFORMATEX
	Algorithm3D GUID
}

// WAVEFORMATEX defines the format of waveform-audio data.
type WAVEFORMATEX struct {
	FormatTag      uint16
	Channels       uint16 // number of channels (i.e. mono, stereo...)
	SamplesPerSec  uint32
	AvgBytesPerSec uint32 // for buffer estimation
	BlockAlign     uint16 // block size of data
	BitsPerSample  uint16 // number of bits per sample of mono data
	Size           uint16 // size of extra information (after Size) in bytes
}

// WAVEFORMATEXTENSIBLE defines the format of waveform-audio data for formats
// having more than two channels or higher resolutions than allowed by
// WAVEFORMATEX.
type WAVEFORMATEXTENSIBLE struct {
	Format              WAVEFORMATEX
	wValidBitsPerSample uint16
	wSamplesPerBlock    uint16
	wReserved           uint16
	dwChannelMask       uint32
	SubFormat           uint32
}

// GUID is a global unique identifier.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// CAPS describes the capabilities of a device.
type CAPS struct {
	Size                         uint32
	Flags                        uint32
	MinSecondarySampleRate       uint32
	MaxSecondarySampleRate       uint32
	PrimaryBuffers               uint32
	MaxHwMixingAllBuffers        uint32
	MaxHwMixingStaticBuffers     uint32
	MaxHwMixingStreamingBuffers  uint32
	FreeHwMixingAllBuffers       uint32
	FreeHwMixingStaticBuffers    uint32
	FreeHwMixingStreamingBuffers uint32
	MaxHw3DAllBuffers            uint32
	MaxHw3DStaticBuffers         uint32
	MaxHw3DStreamingBuffers      uint32
	FreeHw3DAllBuffers           uint32
	FreeHw3DStaticBuffers        uint32
	FreeHw3DStreamingBuffers     uint32
	TotalHwMemBytes              uint32
	FreeHwMemBytes               uint32
	MaxContigFreeHwMemBytes      uint32
	UnlockTransferRateHwBuffers  uint32
	PlayCpuOverheadSwBuffers     uint32
	Reserved1                    uint32
	Reserved2                    uint32
}

// CBCAPS describes the capabilities of a capture buffer.
type CBCAPS struct {
	Size        uint32
	Flags       uint32
	BufferBytes uint32
	Reserved    uint32
}

// CBUFFERDESC describes a capture buffer.
type CBUFFERDESC struct {
	Size        uint32
	Flags       uint32
	BufferBytes uint32
	Reserved    uint32
	WfxFormat   *WAVEFORMATEX
}

// CCAPS describes the capabilities of a capture device.
type CCAPS struct {
	Size     uint32
	Flags    uint32
	Formats  uint32
	Channels uint32
}

// CEFFECTDESC contains parameters for an effect associated with a capture
// buffer.
type CEFFECTDESC struct {
	Size              uint32
	Flags             uint32
	GuidDSCFXClass    GUID
	GuidDSCFXInstance GUID
	Reserved1         uint32
	Reserved2         uint32
}

// CFXAec contains parameters for acoustic echo cancellation in a capture
// buffer.
type CFXAec struct {
	Enable    int32
	NoiseFill int32
	Mode      uint32
}

// CFXNoiseSuppress contains parameters for noise suppression in a capture
// buffer.
type CFXNoiseSuppress struct {
	Enable int32
}

// EFFECTDESC describes an effect associated with a buffer.
type EFFECTDESC struct {
	Size          uint32
	Flags         uint32
	GuidDSFXClass GUID
	Reserved1     *uint32
	Reserved2     *uint32
}

// FXI3DL2Reverb contains parameters for an I3DL2 (Interactive 3D Audio Level 2)
// reverberation effect.
type FXI3DL2Reverb struct {
	Room              int32   // [-10000, 0]     default: -1000 mB
	RoomHF            int32   // [-10000, 0]     default: 0 mB
	RoomRolloffFactor float32 // [0.0, 10.0]     default: 0.0
	DecayTime         float32 // [0.1, 20.0]     default: 1.49s
	DecayHFRatio      float32 // [0.1, 2.0]      default: 0.83
	Reflections       int32   // [-10000, 1000]  default: -2602 mB
	ReflectionsDelay  float32 // [0.0, 0.3]      default: 0.007 s
	Reverb            int32   // [-10000, 2000]  default: 200 mB
	ReverbDelay       float32 // [0.0, 0.1]      default: 0.011 s
	Diffusion         float32 // [0.0, 100.0]    default: 100.0 %
	Density           float32 // [0.0, 100.0]    default: 100.0 %
	HFReference       float32 // [20.0, 20000.0] default: 5000.0 Hz
}

// FXChorus contains parameters for a chorus effect.
type FXChorus struct {
	WetDryMix float32
	Depth     float32
	Feedback  float32
	Frequency float32
	Waveform  int32 // LFO shape; FXCHORUS_WAVE_xxx
	Delay     float32
	Phase     int32
}

// FXCompressor contains parameters for a compression effect.
type FXCompressor struct {
	Gain      float32
	Attack    float32
	Release   float32
	Threshold float32
	Ratio     float32
	Predelay  float32
}

// FXDistortion contains parameters for a distortion effect.
type FXDistortion struct {
	Gain                  float32
	Edge                  float32
	PostEQCenterFrequency float32
	PostEQBandwidth       float32
	PreLowpassCutoff      float32
}

// FXEcho contains parameters for an echo effect.
type FXEcho struct {
	WetDryMix  float32
	Feedback   float32
	LeftDelay  float32
	RightDelay float32
	PanDelay   int32
}

// FXFlanger contains parameters for a flange effect.
type FXFlanger struct {
	WetDryMix float32
	Depth     float32
	Feedback  float32
	Frequency float32
	Waveform  int32
	Delay     float32
	Phase     int32
}

// FXGargle contains parameters for an amplitude modulation effect.
type FXGargle struct {
	RateHz    uint32 // rate of modulation in hz
	WaveShape uint32 // FXGARGLE_WAVE_xxx
}

// FXParamEq contains parameters for a parametric equalizer effect.
type FXParamEq struct {
	Center    float32
	Bandwidth float32
	Gain      float32
}

// FXWavesReverb contains parameters for a Waves reverberation effect.
type FXWavesReverb struct {
	InGain          float32 // [-96.0,0.0]    default: 0.0 dB
	ReverbMix       float32 // [-96.0,0.0]    default: 0.0 db
	ReverbTime      float32 // [0.001,3000.0] default: 1000.0 ms
	HighFreqRTRatio float32 // [0.001,0.999]  default: 0.001
}

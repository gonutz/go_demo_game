/*
Package ds provides a wrapper for Microsoft's DirectSound API in pure Go.
It can only be used on Windows.

When running a DirectSound application you need to have dsound.dll installed on
the system, which fortunately has been deployed with all Windows versions since
XP. This means if you have Go installed, you also have the DLL installed. This
also means that your application can be deployed without the DirectX DLLs and if
you have no other dependencies you can just give the executable file to the
users.

NOTE that the prefix DS is left out of all identifiers. In cases where this
makes the name start with "3D", the 3D part is appended to the name instead,
e.g. DS3DBUFFER becomes BUFFER3D.
*/
package ds

import (
	"syscall"
	"unsafe"
)

var (
	dll                = syscall.NewLazyDLL("dsound.dll")
	directSoundCreate8 = dll.NewProc("DirectSoundCreate8")
)

// Create creates and initializes an object that supports the DirectSound
// interface.
// Use nil as the guid for the default device.
// The application must call the DirectSound.SetCooperativeLevel method
// immediately after creating a device object.
func Create(guid *GUID) (*DirectSound, Error) {
	var obj *DirectSound
	ret, _, _ := directSoundCreate8.Call(
		uintptr(unsafe.Pointer(guid)),
		uintptr(unsafe.Pointer(&obj)),
		0,
	)
	return obj, toErr(ret)
}

// DirectSound is used to create buffer objects, manage devices, and set up the
// environment. Use Create to create an instance of this type.
type DirectSound struct {
	vtbl *directSoundVtbl
}

type directSoundVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	CreateSoundBuffer    uintptr
	GetCaps              uintptr
	DuplicateSoundBuffer uintptr
	SetCooperativeLevel  uintptr
	Compact              uintptr
	GetSpeakerConfig     uintptr
	SetSpeakerConfig     uintptr
	Initialize           uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *DirectSound) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.AddRef,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *DirectSound) Release() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// CreateSoundBuffer creates a sound buffer object to manage audio samples.
// The Size of the BUFFERDESC is set automatically.
func (obj *DirectSound) CreateSoundBuffer(desc BUFFERDESC) (buf *Buffer, err Error) {
	desc.Size = uint32(unsafe.Sizeof(desc))
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.CreateSoundBuffer,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&desc)),
		uintptr(unsafe.Pointer(&buf)),
		0,
	)
	err = toErr(ret)
	return
}

// GetCaps retrieves the capabilities of the hardware device that is
// represented by the device object.
func (obj *DirectSound) GetCaps() (caps CAPS, err Error) {
	caps.Size = uint32(unsafe.Sizeof(caps))
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetCaps,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&caps)),
	)
	err = toErr(ret)
	return
}

// DuplicateSoundBuffer creates a new secondary buffer that shares the original
// buffer's memory.
func (obj *DirectSound) DuplicateSoundBuffer(orig *Buffer) (copy *Buffer, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.DuplicateSoundBuffer,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(orig)),
		uintptr(unsafe.Pointer(&copy)),
	)
	err = toErr(ret)
	return
}

// SetCooperativeLevel sets the cooperative level of the application for this
// sound device.
func (obj *DirectSound) SetCooperativeLevel(window HWND, level uint32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetCooperativeLevel,
		uintptr(unsafe.Pointer(obj)),
		uintptr(window),
		uintptr(level),
	)
	return toErr(ret)
}

// Compact has no effect.
func (obj *DirectSound) Compact() Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Compact,
		uintptr(unsafe.Pointer(obj)),
	)
	return toErr(ret)
}

// GetSpeakerConfig retrieves the speaker configuration.
// The config returned can be a packed uint32 containing both configuration and
// geometry information. Use the SPEAKER_CONFIG and SPEAKER_GEOMETRY functions
// to unpack the uint32.
func (obj *DirectSound) GetSpeakerConfig() (config uint32, err Error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetSpeakerConfig,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&config)),
	)
	err = toErr(ret)
	return
}

// SetSpeakerConfig specifies the speaker configuration of the device.
func (obj *DirectSound) SetSpeakerConfig(config uint32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetSpeakerConfig,
		uintptr(unsafe.Pointer(obj)),
		uintptr(config),
	)
	return toErr(ret)
}

// Initialize initializes a device object that was created by using the
// CoCreateInstance function.
func (obj *DirectSound) Initialize(device *GUID) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Initialize,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(device)),
	)
	return toErr(ret)
}

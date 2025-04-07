package di8

import (
	"syscall"
	"unsafe"
)

var (
	dll                = syscall.NewLazyDLL("dinput8.dll")
	directInput8Create = dll.NewProc("DirectInput8Create")
)

// Create returns the basic DirectInput8 object that you need to query and
// create input devices.
func Create(windowInstance HINSTANCE) (*DirectInput, error) {
	var obj *DirectInput
	ret, _, _ := directInput8Create.Call(
		uintptr(windowInstance),
		VERSION,
		uintptr(unsafe.Pointer(&IID_IDirectInput8W)),
		uintptr(unsafe.Pointer(&obj)),
		0,
	)
	return obj, toErr(ret)
}

// DirectInput lets you query and create input devices.
type DirectInput struct {
	vtbl *directInputVtbl
}

type directInputVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	CreateDevice    uintptr
	EnumDevices     uintptr
	GetDeviceStatus uintptr
	RunControlPanel uintptr
	Initialize      uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *DirectInput) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.AddRef,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *DirectInput) Release() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// CreateDevice creates a mouse, keyboard or game controller device. The guid
// can be retrieved from EnumDevice. It is the DEVICEINSTANCE.GuidInstance that
// is passed to the enumeration callback.
func (obj *DirectInput) CreateDevice(guid GUID) (device *Device, err error) {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.CreateDevice,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&guid)),
		uintptr(unsafe.Pointer(&device)),
		0,
	)
	return device, toErr(ret)
}

// EnumDevices looks for all devices of the given type and calls the given
// callback with each of them.
//
// Return ENUM_CONTINUE or ENUM_STOP from the callback. As long as you return
// ENUM_CONTINUE, enumeration contrinues, once you return ENUM_STOP,
// enumeration stops.
//
// devType can be one of:
// - DEVCLASS_ALL: all devices.
// - DEVCLASS_DEVICE: devices that fall into none of the other classes.
// - DEVCLASS_GAMECTRL: game controllers.
// - DEVCLASS_KEYBOARD: keyboards.
// - DEVCLASS_POINTER: mouse and screen pointer devices.
//
// flags can be a combination of:
// - EDFL_ALLDEVICES: all installed devices.
// - EDFL_ATTACHEDONLY: all installed and attached devices.
// - EDFL_FORCEFEEDBACK: devices that support force feedback (rumbling).
// - EDFL_INCLUDEALIASES: include devices that are aliases for other devices.
// - EDFL_INCLUDEHIDDEN: include hidden devices.
// - EDFL_INCLUDEPHANTOMS: include placeholder devices.
//
// context will be passed to the callback.
func (obj *DirectInput) EnumDevices(
	devType uint32,
	callback func(instance *DEVICEINSTANCE, context uintptr) uintptr,
	context uintptr,
	flags uint32) error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.EnumDevices,
		uintptr(unsafe.Pointer(obj)),
		uintptr(devType),
		syscall.NewCallback(callback),
		context,
		uintptr(flags),
	)
	return toErr(ret)
}

// RunControlPanel runs Control Panel to enable the user to install a new input
// device or modify configurations.
func (obj *DirectInput) RunControlPanel(owner HWND) error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.RunControlPanel,
		uintptr(unsafe.Pointer(obj)),
		uintptr(owner),
		0,
	)
	return toErr(ret)
}

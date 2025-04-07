package di8

import (
	"syscall"
	"unsafe"
)

// Device represents a mouse, keyboard or game controller.
type Device struct {
	vtbl *deviceVtbl
}

type deviceVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetCapabilities      uintptr
	EnumObjects          uintptr
	GetProperty          uintptr
	SetProperty          uintptr
	Acquire              uintptr
	Unacquire            uintptr
	GetDeviceState       uintptr
	GetDeviceData        uintptr
	SetDataFormat        uintptr
	SetEventNotification uintptr
	SetCooperativeLevel  uintptr
	GetObjectInfo        uintptr
	GetDeviceInfo        uintptr
	RunControlPanel      uintptr
	Initialize           uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on an
// object.
func (obj *Device) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.AddRef,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *Device) Release() uint32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return uint32(ret)
}

// Acquire prepares the Device to be queried for input. Before calling Acquire,
// you must call SetDataFormat() and set the buffer size property with
// SetProperty(di8.PROP_BUFFERSIZE). When you are done with the Device, call
// Unacquire().
func (obj *Device) Acquire() Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Acquire,
		uintptr(unsafe.Pointer(obj)),
	)
	return toErr(ret)
}

// Unacquire stops using the Device.
func (obj *Device) Unacquire() Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Unacquire,
		uintptr(unsafe.Pointer(obj)),
	)
	return toErr(ret)
}

// SetCooperativeLevel allows the device to be queried exclusively and even
// when the window is not in focus. flags can be a combination of:
// - SCL_BACKGROUND: read the device even when the window is not active.
// - SCL_EXCLUSIVE: only one window at a time can have exclusive access.
// - SCL_FOREGROUND: the device is automatically unacquired when the window
// becomes inactive.
// - SCL_NONEXCLUSIVE: non-exclusive access is always granted.
// - SCL_NOWINKEY: disables the Windows logo key.
func (obj *Device) SetCooperativeLevel(window HWND, flags uint32) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetCooperativeLevel,
		uintptr(unsafe.Pointer(obj)),
		uintptr(window),
		uintptr(flags),
	)
	return toErr(ret)
}

// SetDataFormat defines the data format that will be used for GetDeviceState.
// Use these predefined formats with these device states:
// - SetDataFormat(&di8. Keyboard): di8.KEYBOARDSTATE
// - SetDataFormat(&di8.Mouse): di8.MOUSESTATE
// - SetDataFormat(&di8.Mouse2): di8.MOUSESTATE2
// - SetDataFormat(&di8.Joystick): di8.JOYSTATE
// - SetDataFormat(&di8.Joystick2): di8.JOYSTATE2
func (obj *Device) SetDataFormat(format *DATAFORMAT) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetDataFormat,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(format)),
	)
	return toErr(ret)
}

// GetDeviceData retrieves the device events that occurred since the last call
// to GetDeviceData. data provides a buffer for the events, the returned number
// indicates how many were actually filled by GetDeviceData.
func (obj *Device) GetDeviceData(data []DEVICEOBJECTDATA, flags uint32) (int, Error) {
	var dataPtr uintptr
	if len(data) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&data[0]))
	}
	count := uint32(len(data))
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetDeviceData,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Sizeof(DEVICEOBJECTDATA{})),
		dataPtr,
		uintptr(unsafe.Pointer(&count)),
		uintptr(flags),
	)
	return int(count), toErr(ret)
}

// EnumObjects calls callback for every controll object, like buttons, sliders
// and povs, on the device.
func (obj *Device) EnumObjects(
	callback func(object *DEVICEOBJECTINSTANCE, context uintptr) uintptr,
	context uintptr,
	flags uint32,
) error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.EnumObjects,
		uintptr(unsafe.Pointer(obj)),
		syscall.NewCallback(callback),
		context,
		uintptr(flags),
	)
	return toErr(ret)
}

// DeviceState is the base type for KEYBOARDSTATE, MOUSESTATE, MOUSESTATE2,
// JOYSTATE and JOYSTATE2, which can be used as arguments to GetDeviceState.
type DeviceState interface {
	ptr() uintptr
	size() int
}

// GetDeviceState fills state with the current state of the device's controlls.
// Make sure you call SetDataFormat before using GetDeviceState and that the
// state type corresponds to the format.
func (obj *Device) GetDeviceState(state DeviceState) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetDeviceState,
		uintptr(unsafe.Pointer(obj)),
		uintptr(state.size()),
		state.ptr(),
	)
	return toErr(ret)
}

// Property is the base type for PROPCPOINTS, PROPDWORD, PROPRANGE, PROPCAL,
// PROPCALPOV, PROPGUIDANDPATH, PROPSTRING and PROPPOINTER. Create them with
// the NewProp*. These can be passed to SetProperty.
type Property interface {
	propHeader() *PROPHEADER
}

// SetProperty sets one of the PROP_* properties for the device. Predefined
// property types are: PROPCPOINTS, PROPDWORD, PROPRANGE, PROPCAL, PROPCALPOV,
// PROPGUIDANDPATH, PROPSTRING and PROPPOINTER. Create them with the NewProp*
// functions.
func (obj *Device) SetProperty(guid *GUID, prop Property) Error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.SetProperty,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(guid)),
		uintptr(unsafe.Pointer(prop.propHeader())),
	)
	return toErr(ret)
}

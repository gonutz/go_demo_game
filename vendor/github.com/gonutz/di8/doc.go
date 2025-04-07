/*
di8 provides a way to access Windows' DirectInput API in Go.

Call di8.Create() to create the basic DirectInput object.

Use DirectInput.EnumDevices() to list all the mouse, keyboard and game
controller devices.

Use DirectInput.CreateDevice() to create Device objects for all devices that
you want to use.

To be able to query data from a Device, call Device.SetDataFormat() first, then
Device.SetProperty(di8.PROP_BUFFERSIZE) and lastly Device.Acquire(). When you
are done with it, call Device.Unacquire().

There are two ways to query a Device: get the current state or get events when
the state changes.
To get the current state, call Device.GetDeviceState.
To get state change events, call Device.GetDeviceData.

Call Release() on all objects when you are done using them.
*/
package di8

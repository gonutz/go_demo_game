# ds
This library is a pure Go wrapper for Microsoft's DirectSound8 API.

# Installation
Get and build the library with:

    go get -u github.com/gonutz/ds

To run a DirectSound application you need to have `dsound.dll` installed on your system. Luckily, all Windows versions starting with Windows XP have this pre-installed. This means that you will not need to ship any additional DLL with your application when using this library.

# Usage
All DirectSound8 interfaces are translated to Go types and their methods are translated to functions on the types so the usage is very close to the C++ API.

There are some differences in the names in Go, since the package is named `ds`, all names in that package drop the `DS` and `8` parts because they would be redundant. The changes are:

- Interfaces drop the `IDirectSound` prefix and the `8` suffix, e.g. `IDirectSound3DBuffer8` becomes `ds.Buffer`. The only exception is `IDirectSound8` which in Go becomes `DirectSound`.
- Constants and enumerations drop the `DS` prefix, otherwise they are the same and keep the upper case convention so users of DirectSound can easily find what they are looking for. For example `DSSPEAKER_STEREO` becomes `ds.SPEAKER_STEREO`. Names that start with `DS3D` start with an underscore in Go, e.g. `DS3DMODE_NORMAL` becomes `ds._3DMODE_NORMAL`.
- Structs, like constants, only drop the `DS` prefix, they too keep the upper case naming convention, so `DSBCAPS` becomes `ds.BCAPS`. There are two exceptions: `DS3DBUFFER` becomes `BUFFER3D` and `DS3DLISTENER` becomes `LISTENER3D` in Go.
- Error constants also drop the `DS` prefix so `DSERR_OUTOFMEMORY` becomes `ds.ERR_OUTOFMEMORY`. However, the interface functions do not return these constants, they return Go `error`s instead of `HRESULT`s.
- Instead of returning `HRESULT`, functions return `ds.Error` which implements the standard Go error interface and has an additional function `Code() int32` which returns the error code. This code can be checked against the defined error constants. If a function succeeds it returns `nil` (and not `DS_OK`) as the ds.Error.

Note that DirectSound needs a window instance for setting it up. This means that you need some way to create a native window and get the handle to pass it to ds. Libraries to help you do that include the [SDL2 Go wrapper](https://github.com/veandco/go-sdl2) and [Allen Dang's w32](https://github.com/AllenDang/w32). You could also use other Windows wrapper libraries, like the [walk library](https://github.com/lxn/walk), or just write a little CGo code to set up a window yourself. This library does not provide window creation or event handling functionality, only the DirectSound8 wrapper.

All calls to DirectSound must happen from the same thread that creates the DirectSound object so make sure to add this code in your main package:

    func init() {
	    runtime.LockOSThread()
	}

There are some additional convenience functions. `Buffer` has a `Lock` function which returns a `void*` pointer in the C API and would thus return `uintptr`s in Go. You can use these pointers to read and write various types from/to that memory. However, using `uintptr` or `unsafe.Pointer` is not idiomatic Go so the `Lock` function returns a wrapper around the `uintptr` instead, providing a `Size` and `Write` function which takes a slice of `[]byte` and handles copying the data for you. See the documentation of these functions for further information.

# Documentation
See the [GoDoc](https://godoc.org/github.com/gonutz/ds) for the Go API. The functions are only documented very generally, to get more detailed information about the DirectSound API see the [MSDN documentation](https://msdn.microsoft.com/en-us/library/windows/desktop/ee416975(v=vs.85).aspx).

# Status
The API is currently incomplete, only the `IDirectSound8` and `IDirectSoundBuffer8` interfaces have been translated so far. If you need the other interfaces or are missing any functions, please write an issue or even create a pull request.

# Help improve this library

Only real world use and feedback can improve the usability of this library, so please use it, fork it, send pull requests and create issues to help improve this library.

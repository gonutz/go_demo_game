package di8

import "strconv"

// Error is returned by all di8 functions. It encapsulates the error code
// returned by DirectInput. If a function succeeds it will return nil as the
// Error and if it fails you can retrieve the error code using the Code()
// function. You can check the result against the predefined error codes, like
// ERR_INVALIDPARAM, ERR_OUTOFMEMORY, etc.
type Error interface {
	error
	// Code returns the DirectInput error code for a function. Call this
	// function only if the Error is not nil, if the error code is DI_OK or any
	// other code that signifies success, a function will return nil as the
	// Error instead of a non-nil error with that code in it. This way,
	// functions behave in a standard Go way, returning nil as the error in case
	// of success and only returning non-nil errors if something went wrong.
	Code() int32
}

func toErr(result uintptr) Error {
	res := hResultError(result)
	if res >= 0 {
		return nil
	}
	return res
}

type hResultError int32

func (r hResultError) Code() int32 {
	return int32(r)
}

func (e hResultError) Error() string {
	// These casts are needed to make comparisons work correctly.
	switch int64(int32(e)) {
	case ERR_ACQUIRED:
		return "ERR_ACQUIRED: The operation cannot be performed while the device is acquired."
	case ERR_ALREADYINITIALIZED:
		return "ERR_ALREADYINITIALIZED: This object is already initialized"
	case ERR_BADDRIVERVER:
		return "ERR_BADDRIVERVER: The object could not be created due to an incompatible driver version or mismatched or incomplete driver components."
	case ERR_BETADIRECTINPUTVERSION:
		return "ERR_BETADIRECTINPUTVERSION: The object could not be created due to an incompatible driver version or mismatched or incomplete driver components."
	case ERR_DEVICEFULL:
		return "ERR_DEVICEFULL: The device is full."
	case ERR_DEVICENOTREG:
		return "ERR_DEVICENOTREG: The device or device instance is not registered with DirectInput. This value is equal to the REGDB_E_CLASSNOTREG standard COM return value."
	case ERR_EFFECTPLAYING:
		return "ERR_EFFECTPLAYING: The parameters were updated in memory but were not downloaded to the device because the device does not support updating an effect while it is still playing."
	case ERR_GENERIC:
		return "ERR_GENERIC: An undetermined error occurred inside the DirectInput subsystem. This value is equal to the E_FAIL standard COM return value."
	case ERR_HANDLEEXISTS:
		return "ERR_HANDLEEXISTS: The device already has an event notification associated with it. This value is equal to the E_ACCESSDENIED standard COM return value."
	case ERR_HASEFFECTS:
		return "ERR_HASEFFECTS: The device cannot be reinitialized because effects are attached to it."
	case ERR_INCOMPLETEEFFECT:
		return "ERR_INCOMPLETEEFFECT: The effect could not be downloaded because essential information is missing. For example, no axes have been associated with the effect, or no type-specific information has been supplied."
	case ERR_INPUTLOST:
		return "ERR_INPUTLOST: Access to the input device has been lost. It must be reacquired."
	case ERR_INVALIDPARAM:
		return "ERR_INVALIDPARAM: An invalid parameter was passed to the returning function, or the object was not in a state that permitted the function to be called. This value is equal to the E_INVALIDARG standard COM return value."
	case ERR_MAPFILEFAIL:
		return "ERR_MAPFILEFAIL: An error has occurred either reading the vendor-supplied action-mapping file for the device or reading or writing the user configuration mapping file for the device."
	case ERR_MOREDATA:
		return "ERR_MOREDATA: Not all the requested information fit into the buffer."
	case ERR_NOAGGREGATION:
		return "ERR_NOAGGREGATION: This object does not support aggregation."
	case ERR_NOINTERFACE:
		return "ERR_NOINTERFACE: The object does not support the specified interface. This value is equal to the E_NOINTERFACE standard COM return value."
	case ERR_NOTACQUIRED:
		return "ERR_NOTACQUIRED: The operation cannot be performed unless the device is acquired."
	case ERR_NOTBUFFERED:
		return "ERR_NOTBUFFERED: The device is not buffered. Set the DIPROP_BUFFERSIZE property to enable buffering."
	case ERR_NOTDOWNLOADED:
		return "ERR_NOTDOWNLOADED: The effect is not downloaded."
	case ERR_NOTEXCLUSIVEACQUIRED:
		return "ERR_NOTEXCLUSIVEACQUIRED: The operation cannot be performed unless the device is acquired in DISCL_EXCLUSIVE mode."
	case ERR_NOTFOUND:
		return "ERR_NOTFOUND, ERR_OBJECTNOTFOUND: The requested object does not exist."
	case ERR_NOTINITIALIZED:
		return "ERR_NOTINITIALIZED: This object has not been initialized."
	case BUFFEROVERFLOW:
		return "DI_BUFFEROVERFLOW: the device data was truncated."
	default:
		return "Error code " + strconv.Itoa(int(e))
	}
}

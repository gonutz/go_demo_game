package ds

import "strconv"

// Error is returned by all DirectSound functions. It encapsulates the error
// code returned by DirectSound. If a function succeeds it will return nil as
// the Error and if it fails you can retrieve the error code using the Code()
// function. You can check the result against the predefined error codes (like
// ERR_INVALIDCALL, E_OUTOFMEMORY etc).
type Error interface {
	error
	// Code returns the DirectSound error code for a function. Call this
	// function only if the Error is not nil, if the error code is OK or any
	// other code that signifies success, a function will return nil as the
	// Error instead of a non-nil error with that code in it. This way,
	// functions behave in a standard Go way, returning nil as the error in case
	// of success and only returning non-nil errors if something went wrong.
	Code() int32
}

func toErr(result uintptr) Error {
	res := hResultError(result) // cast to signed int
	if res >= 0 {
		return nil
	}
	return res
}

type hResultError int32

func (r hResultError) Code() int32 { return int32(r) }

func (r hResultError) Error() string {
	switch r {
	case ERR_ALLOCATED:
		return "resource already alloceated"
	case ERR_CONTROLUNAVAIL:
		return "control not available"
	case ERR_INVALIDPARAM:
		return "invalid parameter"
	case ERR_INVALIDCALL:
		return "invalid call"
	case ERR_GENERIC:
		return "undetermined error"
	case ERR_PRIOLEVELNEEDED:
		return "insufficient priority level"
	case ERR_OUTOFMEMORY:
		return "out of memory"
	case ERR_BADFORMAT:
		return "bad format"
	case ERR_UNSUPPORTED:
		return "unsupported"
	case ERR_NODRIVER:
		return "no driver"
	case ERR_ALREADYINITIALIZED:
		return "already initialized"
	case ERR_NOAGGREGATION:
		return "no aggregation"
	case ERR_BUFFERLOST:
		return "buffer lost"
	case ERR_OTHERAPPHASPRIO:
		return "other app has priority"
	case ERR_UNINITIALIZED:
		return "unintialized"
	case ERR_NOINTERFACE:
		return "COM interface unavailable"
	case ERR_ACCESSDENIED:
		return "access denied"
	case ERR_BUFFERTOOSMALL:
		return "buffer too short"
	case ERR_DS8_REQUIRED:
		return "DirectSound8 required"
	case ERR_SENDLOOP:
		return "circular loop of send effects detected"
	case ERR_BADSENDBUFFERGUID:
		return "GUID in audiopath file does not match a valid MIXIN buffer"
	case ERR_OBJECTNOTFOUND:
		return "object not found"
	case ERR_FXUNAVAILABLE:
		return "effect unavailable"
	default:
		return "unknown error code " + strconv.Itoa(int(r))
	}
}

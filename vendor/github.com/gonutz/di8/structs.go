package di8

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const max_path = 260

func toString(s []uint16) string {
	for i, c := range s {
		if c == 0 {
			return string(utf16.Decode(s[:i]))
		}
	}
	return string(utf16.Decode(s))
}

type HWND uintptr

type HINSTANCE uintptr

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

type DEVICEINSTANCE struct {
	Size         uint32
	GuidInstance GUID
	GuidProduct  GUID
	DevType      uint32
	InstanceName [max_path]uint16
	ProductName  [max_path]uint16
	GuidFFDriver GUID
	UsagePage    uint16
	Usage        uint16
}

func (d *DEVICEINSTANCE) GetInstanceName() string {
	return toString(d.InstanceName[:])
}

func (d *DEVICEINSTANCE) GetProductName() string {
	return toString(d.ProductName[:])
}

type DATAFORMAT struct {
	Size     uint32
	ObjSize  uint32
	Flags    uint32
	DataSize uint32
	NumObjs  uint32
	Rgodf    *OBJECTDATAFORMAT
}

type OBJECTDATAFORMAT struct {
	Guid  *GUID
	Ofs   uint32
	Type  uint32
	Flags uint32
}

type PROPHEADER struct {
	Size       uint32
	HeaderSize uint32
	Obj        uint32
	How        uint32
}

var _ Property = &PROPHEADER{}

// We define propHeader on the base type so that all types that embed
// PROPHEADER will fulfill the Property interface.
func (p *PROPHEADER) propHeader() *PROPHEADER {
	return p
}

func NewPropCPoints(obj, how uint32, points []CPOINT) *PROPCPOINTS {
	var p PROPCPOINTS
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	if len(points) > MAXCPOINTSNUM {
		points = points[:MAXCPOINTSNUM]
	}
	p.CPointsNum = uint32(len(points))
	for i := range points {
		p.Points[i] = points[i]
	}
	return &p
}

type PROPCPOINTS struct {
	PROPHEADER
	CPointsNum uint32
	Points     [MAXCPOINTSNUM]CPOINT
}

var _ Property = &PROPCPOINTS{}

type CPOINT struct {
	P   int32
	Log uint32
}

func NewPropDWord(obj, how, data uint32) *PROPDWORD {
	var p PROPDWORD
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	p.Data = data
	return &p
}

type PROPDWORD struct {
	PROPHEADER
	Data uint32
}

var _ Property = &PROPDWORD{}

func NewPropRange(obj, how uint32, min, max int32) *PROPRANGE {
	var p PROPRANGE
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	p.Min = min
	p.Max = max
	return &p
}

type PROPRANGE struct {
	PROPHEADER
	Min int32
	Max int32
}

var _ Property = &PROPRANGE{}

func NewPropCal(obj, how uint32, min, center, max int32) *PROPCAL {
	var p PROPCAL
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	p.Min = min
	p.Center = center
	p.Max = max
	return &p
}

type PROPCAL struct {
	PROPHEADER
	Min    int32
	Center int32
	Max    int32
}

var _ Property = &PROPCAL{}

func NewPropCalPOV(obj, how uint32, min, max [5]int32) *PROPCALPOV {
	var p PROPCALPOV
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	p.Min = min
	p.Max = max
	return &p
}

type PROPCALPOV struct {
	PROPHEADER
	Min [5]int32
	Max [5]int32
}

var _ Property = &PROPCALPOV{}

func NewPropGuidAndPath(obj, how uint32, guid GUID, path string) *PROPGUIDANDPATH {
	var p PROPGUIDANDPATH
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	p.GuidClass = guid
	str, err := syscall.UTF16FromString(path)
	if err == nil {
		if len(str) > max_path {
			str = str[:max_path]
			str[max_path-1] = 0
		}
		for i := range str {
			p.Path[i] = str[i]
		}
	}
	return &p
}

type PROPGUIDANDPATH struct {
	PROPHEADER
	GuidClass GUID
	Path      [max_path]uint16
}

var _ Property = &PROPGUIDANDPATH{}

func (p *PROPGUIDANDPATH) GetPath() string {
	return toString(p.Path[:])
}

func NewPropString(obj, how uint32, s string) *PROPSTRING {
	var p PROPSTRING
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	str, err := syscall.UTF16FromString(s)
	if err == nil {
		if len(str) > max_path {
			str = str[:max_path]
			str[max_path-1] = 0
		}
		for i := range str {
			p.String[i] = str[i]
		}
	}
	return &p
}

type PROPSTRING struct {
	PROPHEADER
	String [max_path]uint16
}

var _ Property = &PROPSTRING{}

func (p *PROPSTRING) GetString() string {
	return toString(p.String[:])
}

func NewPropPointer(obj, how uint32, pointer uintptr) *PROPPOINTER {
	var p PROPPOINTER
	p.Obj = obj
	p.How = how
	p.Size = uint32(unsafe.Sizeof(p))
	p.HeaderSize = uint32(unsafe.Sizeof(p.PROPHEADER))
	p.Data = pointer
	return &p
}

type PROPPOINTER struct {
	PROPHEADER
	Data uintptr
}

var _ Property = &PROPPOINTER{}

type DEVICEOBJECTDATA struct {
	Ofs       uint32
	Data      uint32
	TimeStamp uint32
	Sequence  uint32
	AppData   uintptr
}

type MOUSESTATE struct {
	X       int32
	Y       int32
	Z       int32
	Buttons [4]byte
}

var _ DeviceState = &MOUSESTATE{}

func (s *MOUSESTATE) ptr() uintptr { return uintptr(unsafe.Pointer(s)) }
func (s *MOUSESTATE) size() int    { return int(unsafe.Sizeof(*s)) }

type MOUSESTATE2 struct {
	X       int32
	Y       int32
	Z       int32
	Buttons [8]byte
}

var _ DeviceState = &MOUSESTATE2{}

func (s *MOUSESTATE2) ptr() uintptr { return uintptr(unsafe.Pointer(s)) }
func (s *MOUSESTATE2) size() int    { return int(unsafe.Sizeof(*s)) }

type KEYBOARDSTATE [256]byte

var _ DeviceState = &KEYBOARDSTATE{}

func (s *KEYBOARDSTATE) ptr() uintptr { return uintptr(unsafe.Pointer(&s[0])) }
func (s *KEYBOARDSTATE) size() int    { return len(*s) }

type JOYSTATE struct {
	X       int32
	Y       int32
	Z       int32
	Rx      int32
	Ry      int32
	Rz      int32
	Slider  [2]int32
	POV     [4]uint32
	Buttons [32]byte
}

var _ DeviceState = &JOYSTATE{}

func (s *JOYSTATE) ptr() uintptr { return uintptr(unsafe.Pointer(s)) }
func (s *JOYSTATE) size() int    { return int(unsafe.Sizeof(*s)) }

type JOYSTATE2 struct {
	X       int32
	Y       int32
	Z       int32
	Rx      int32
	Ry      int32
	Rz      int32
	Slider  [2]int32
	POV     [4]uint32
	Buttons [128]byte
	VX      int32
	VY      int32
	VZ      int32
	VRx     int32
	VRy     int32
	VRz     int32
	VSlider [2]int32
	AX      int32
	AY      int32
	AZ      int32
	ARx     int32
	ARy     int32
	ARz     int32
	ASlider [2]int32
	FX      int32
	FY      int32
	FZ      int32
	FRx     int32
	FRy     int32
	FRz     int32
	FSlider [2]int32
}

var _ DeviceState = &JOYSTATE2{}

func (s *JOYSTATE2) ptr() uintptr { return uintptr(unsafe.Pointer(s)) }
func (s *JOYSTATE2) size() int    { return int(unsafe.Sizeof(*s)) }

type DEVICEOBJECTINSTANCE struct {
	Size              uint32
	GuidType          GUID
	Ofs               uint32
	Type              uint32
	Flags             uint32
	Name              [max_path]uint16
	FFMaxForce        uint32
	FFForceResolution uint32
	CollectionNumber  uint16
	DesignatorIndex   uint16
	UsagePage         uint16
	Usage             uint16
	Dimension         uint32
	Exponent          uint16
	ReportId          uint16
}

func (d *DEVICEOBJECTINSTANCE) GetName() string {
	return toString(d.Name[:])
}

package main

import (
	"github.com/gonutz/di8"
	"github.com/gonutz/w32/v2"
)

const (
	axisMin = 0.35
	axisMax = 0.95
)

type inputSystem struct {
	dinput         *di8.DirectInput
	joystickDevice *di8.Device
	xboxController xboxControllerState
	joystick       joystickState
}

type xboxControllerState struct {
	connected bool
	// buttons is a bitmask with buttons A, B, X, Y, Back, Start, LB, RB, left
	// axis, right axis.
	// the main button in the center that lets you select your controller index
	// is not being reported to us. It seems to be handled specially by
	// Windows.
	buttons uint16
	// Axes are in the range [-1..1].
	leftXAxis  float32
	leftYAxis  float32
	rightXAxis float32
	rightYAxis float32
	// dpad is in 100 degrees, 0 is north, 4500 north-east, 9000 east, ...
	// 31500 is north-west. Values > 36000 mean the DPad is in idle state.
	dpad uint32
	// Triggers are 0 when released and 1 when pressed all the way down.
	leftTrigger  float32
	rightTrigger float32
}

func (s *xboxControllerState) buttonADown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_A != 0
}

func (s *xboxControllerState) buttonBDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_B != 0
}

func (s *xboxControllerState) buttonXDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_X != 0
}

func (s *xboxControllerState) buttonYDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_Y != 0
}

func (s *xboxControllerState) buttonStartDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_START != 0
}

func (s *xboxControllerState) buttonBackDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_BACK != 0
}

func (s *xboxControllerState) buttonLBDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_LEFT_SHOULDER != 0
}

func (s *xboxControllerState) buttonRBDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_RIGHT_SHOULDER != 0
}

func (s *xboxControllerState) leftAxisDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_LEFT_THUMB != 0
}

func (s *xboxControllerState) rightAxisDown() bool {
	return s.buttons&w32.XINPUT_GAMEPAD_RIGHT_THUMB != 0
}

// joystickState represents the state of our very specific, known joystick.
type joystickState struct {
	xAxis      float32
	yAxis      float32
	buttonDown [8]bool
	// dpad is in 100 degrees, 0 is north, 4500 north-east, 9000 east, ...
	// 31500 is north-west. Values > 36000 mean the DPad is in idle state.
	dpad uint32
	// wheel is 0 when the wheel is rotated all the way back towards the user.
	// wheel is 1 when the wheel is rotated all the way away from the user.
	wheel float32
}

func initInputSystem() (*inputSystem, error) {
	dinput, err := di8.Create(di8.HINSTANCE(w32.GetModuleHandle("")))
	if err != nil {
		return nil, err
	}

	s := &inputSystem{
		dinput: dinput,
	}
	s.connectJoystick()
	return s, nil
}

func (s *inputSystem) close() {
	s.closeJoystick()
	s.dinput.Release()
}

func (s *inputSystem) connectJoystick() {
	if s.joystickDevice != nil {
		return // We are already connected with the joystick.
	}

	var (
		joystickFound bool
		joystickGuid  di8.GUID
	)
	s.dinput.EnumDevices(
		di8.DEVCLASS_GAMECTRL,
		func(device *di8.DEVICEINSTANCE, _ uintptr) uintptr {
			if device.GetProductName() == "Generic   USB  Joystick  " {
				joystickFound = true
				joystickGuid = device.GuidInstance
				return di8.ENUM_STOP
			}
			return di8.ENUM_CONTINUE
		},
		0,
		di8.EDFL_ATTACHEDONLY,
	)

	if !joystickFound {
		return
	}

	if joy, err := s.dinput.CreateDevice(joystickGuid); err == nil {
		if joy.SetDataFormat(&di8.Joystick2) != nil {
			joy.Release()
		} else if joy.SetProperty(
			di8.PROP_BUFFERSIZE,
			di8.NewPropDWord(0, di8.PH_DEVICE, 32),
		) != nil {
			joy.Release()
		} else if joy.Acquire() != nil {
			joy.Release()
		} else {
			s.joystickDevice = joy
		}
	}
}

func (s *inputSystem) closeJoystick() {
	if s.joystickDevice == nil {
		return
	}

	s.joystickDevice.Unacquire()
	s.joystickDevice.Release()
	s.joystickDevice = nil
}

func (s *inputSystem) update() {
	// Reset the controller in case it got lost, we will fill in the data
	// below and overwrite them if it is still connected.
	s.xboxController.connected = false
	s.xboxController.buttons = 0
	s.xboxController.leftXAxis = 0
	s.xboxController.leftYAxis = 0
	s.xboxController.rightXAxis = 0
	s.xboxController.rightYAxis = 0
	s.xboxController.dpad = 0xFFFF
	s.xboxController.leftTrigger = 0
	s.xboxController.rightTrigger = 0

	// We query the first XBox controller that we find.
	for i := 0; i < 4; i++ {
		state, err := w32.XInputGetState(i)
		if err == nil {
			s.xboxController.connected = true
			s.xboxController.buttons = state.Gamepad.Buttons
			s.xboxController.leftXAxis = clampAxis(float32(state.Gamepad.ThumbLX) / 32768)
			s.xboxController.leftYAxis = clampAxis(-float32(state.Gamepad.ThumbLY) / 32768)
			s.xboxController.rightXAxis = clampAxis(float32(state.Gamepad.ThumbRX) / 32768)
			s.xboxController.rightYAxis = clampAxis(-float32(state.Gamepad.ThumbRY) / 32768)
			up := state.Gamepad.Buttons&w32.XINPUT_GAMEPAD_DPAD_UP != 0
			right := state.Gamepad.Buttons&w32.XINPUT_GAMEPAD_DPAD_RIGHT != 0
			down := state.Gamepad.Buttons&w32.XINPUT_GAMEPAD_DPAD_DOWN != 0
			left := state.Gamepad.Buttons&w32.XINPUT_GAMEPAD_DPAD_LEFT != 0
			s.xboxController.dpad = dpadTo100Degrees(up, right, down, left)
			s.xboxController.leftTrigger = float32(state.Gamepad.LeftTrigger) / 255
			s.xboxController.rightTrigger = float32(state.Gamepad.RightTrigger) / 255
			break
		}
	}

	if s.joystickDevice != nil {
		var joyState di8.JOYSTATE2
		disconnected := s.joystickDevice.GetDeviceState(&joyState) != nil
		if disconnected {
			s.closeJoystick()
		} else {
			s.joystick.xAxis = clampAxis(float32(joyState.X-32768) / 32768)
			s.joystick.yAxis = clampAxis(float32(joyState.Y-32768) / 32768)
			for i := range s.joystick.buttonDown {
				s.joystick.buttonDown[i] = joyState.Buttons[i] != 0
			}
			s.joystick.dpad = joyState.POV[0]
			s.joystick.wheel = 1 - float32(joyState.Rz)/0xFFFF
		}
	}
}

func clampAxis(rel float32) float32 {
	if -axisMin <= rel && rel <= axisMin {
		return 0
	}
	if rel > axisMax {
		return 1
	}
	if rel < -axisMax {
		return -1
	}
	return rel
}

func relativeAxis(pos float32) float32 {
	var rel float32
	if pos > 0 {
		rel = (pos - axisMin) / (axisMax - axisMin)
		if rel > 1 {
			rel = 1
		}
	} else if pos < 0 {
		rel = -(pos - -axisMin) / (-axisMax - -axisMin)
		if rel < -1 {
			rel = -1
		}
	}
	return rel
}

func dpadTo100Degrees(up, right, down, left bool) uint32 {
	// We set 4 bits, 1 for each direction, and check that value.
	// x = binary bits: LDRU.

	x := 0
	if up {
		x += 1
	}
	if right {
		x += 2
	}
	if down {
		x += 4
	}
	if left {
		x += 8
	}

	return [16]uint32{
		0:  0xFFFF,
		1:  0,     // up
		2:  9000,  // right
		3:  4500,  // up right
		4:  18000, // down
		5:  0xFFFF,
		6:  13500, // right down
		7:  0xFFFF,
		8:  27000, // left
		9:  31500, // left up
		10: 0xFFFF,
		11: 0xFFFF,
		12: 22500, // left down
		13: 0xFFFF,
		14: 0xFFFF,
		15: 0xFFFF,
	}[x]
}

package main

import (
	"embed"
	"math"
	"math/rand/v2"
	"runtime"
	"syscall"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gonutz/d3d9"
	m "github.com/gonutz/d3dmath/column_major/d3dmath"
	"github.com/gonutz/ds"
	"github.com/gonutz/dxc"
	"github.com/gonutz/ease"
	"github.com/gonutz/obj"
	"github.com/gonutz/w32/v2"
)

//go:embed assets
var assetFiles embed.FS

const fullscreen = true

const fieldOfView = 50

const (
	gameStateFadingIn = iota
	gameStateXBoxControllerFlyingIn
	gameStateXBoxController
	gameStateTransitionToJoystick
	gameStateJoystickRotating
	gameStateJoystickShrinking
	gameStatePlayingLevel
)

var desiredButtonStates = []uint16{
	w32.XINPUT_GAMEPAD_A,
	0,
	w32.XINPUT_GAMEPAD_B,
	0,
	w32.XINPUT_GAMEPAD_B,
	0,
	w32.XINPUT_GAMEPAD_B,
	0,
	w32.XINPUT_GAMEPAD_A,
	0,
	w32.XINPUT_GAMEPAD_B,
	0,
	w32.XINPUT_GAMEPAD_B,
	0,
	w32.XINPUT_GAMEPAD_X,
	0,
	w32.XINPUT_GAMEPAD_Y,
	0,
	w32.XINPUT_GAMEPAD_X,
	0,
	w32.XINPUT_GAMEPAD_START,
	0,
	w32.XINPUT_GAMEPAD_A,
	0,
	w32.XINPUT_GAMEPAD_RIGHT_THUMB,
	0,
}

var floorHeights = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func floorHeightAt(x, z float32) int {
	if x < 0 || z > 0 {
		return 999
	}
	worldW, worldH := len(floorHeights[0]), len(floorHeights)
	tx, ty := int(x), int(-z)
	if 0 <= tx && tx < worldW &&
		0 <= ty && ty < worldH {
		return floorHeights[ty][tx]
	}
	return 999
}

// This function computes our desired sound distortion (the speed at which we
// play the sound), depending on the controller input x, which is in the range
// [-1..1]. It will return a speed of 1 at roughly 0.5, so when the controller
// is moved about half way right.
func makeSoundSpeed(x float64) float64 {
	x *= 0.9
	y := 9 * x * x * x
	return y
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func norm01(x float64) float64 {
	for x < 0 {
		x += 1
	}
	for x > 1 {
		x -= 1
	}
	return x
}

func main() {
	runtime.LockOSThread()

	// These are the state variables used throughout the different states of
	// the game.
	gameState := gameStateFadingIn
	fadeInColor := -100
	const backgroundGray = 200
	xboxBlinkTimer := 0
	joystickBlinkTimer := 0
	controllerFlyTime := 0.0
	const finalControllerZ = 2.0
	const finalControllerXRotation = 0.12
	const joystickScaleSpeed = 0.0061
	controllerYRotation := float32(0)
	controllerXRotation := float32(0)
	specularStrength := float32(0.5)
	specularExponent := float32(16)
	var lastButtonState uint16
	lastButtonStates := make([]uint16, len(desiredButtonStates))
	gamepadScale := float32(1)
	joystickScale := float32(0)
	const joystickYRotationSpeed = 0.0025
	joystickYRotation := float32(0)
	var lastJoystickState joystickState
	var lastXBoxState xboxControllerState
	levelColor := float32(30)
	jokerPos := m.Vec3{9.4, 0, -7.6}
	jokerRot := float32(0.57)
	const jokerBaseRot = -0.25
	jokerSpeed := 0.0
	const jokerAcceleration = 0.004
	const maxJokerSpeed = 0.04
	const minJokerSpeed = -maxJokerSpeed / 2
	jokerLimbRot := 0.0
	const jokerSpeedLimbRatio = 0.55
	cameraCornerPositions := []m.Vec3{
		{9, 5.5, -0.5},
		{17.5, 5.5, -0.5},
		{17.5, 5.5, -9},
		{17.5, 5.5, -17.5},
		{9, 5.5, -17.5},
		{0.5, 5.5, -17.5},
		{0.5, 5.5, -9},
		{0.5, 5.5, -0.5},
	}
	cameraTargetCorner := cameraCornerPositions[5]
	cameraPos := cameraTargetCorner
	cameraInCorner := true
	jokerSpeedY := float32(0)
	const gravity = -0.005
	const jokerJumpSpeed = 0.115
	wasOnGround := true
	stepCoolDown := 0

	pushButtonState := func(s uint16) {
		copy(lastButtonStates, lastButtonStates[1:])
		lastButtonStates[len(lastButtonStates)-1] = s
	}

	lightDir := m.Vec4{1, -1, 1, 0}

	var err error

	input, err := initInputSystem()
	check(err)
	defer input.close()

	var lastMouseX, lastMouseY int
	var rotationAboutY, rotationAboutX float32
	rotationAboutX = 0.1
	translation := float32(4)

	className, _ := syscall.UTF16PtrFromString("game_window_class")
	w32.RegisterClassEx(&w32.WNDCLASSEX{
		Cursor: w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		WndProc: syscall.NewCallback(func(window w32.HWND, msg uint32, w, l uintptr) uintptr {
			switch msg {
			case w32.WM_MOUSEWHEEL:
				delta := float32(int16((w&0xFFFF0000)>>16)) / 120
				translation -= delta
				return 0
			case w32.WM_LBUTTONUP:
				w32.SetCapture(0)
				return 0
			case w32.WM_LBUTTONDOWN:
				w32.SetCapture(window)
				return 0
			case w32.WM_MOUSEMOVE:
				x := int(int16(l & 0x0000FFFF))
				y := int(int16((l & 0xFFFF0000) >> 16))

				if w&w32.MK_LBUTTON != 0 {
					dx, dy := x-lastMouseX, y-lastMouseY
					rotationAboutY += float32(dx) / 1000
					rotationAboutX += float32(dy) / 1000
					if rotationAboutX < -0.25 {
						rotationAboutX = -0.25
					}
					if rotationAboutX > 0.25 {
						rotationAboutX = 0.25
					}
				}

				lastMouseX, lastMouseY = x, y

				return 0
			case w32.WM_KEYDOWN, w32.WM_KEYUP:
				if w == w32.VK_ESCAPE {
					w32.PostQuitMessage(0)
				}
				return 0
			case w32.WM_DEVICECHANGE:
				if w == w32.DBT_DEVNODES_CHANGED {
					input.connectJoystick()
				}
				return 0
			case w32.WM_DESTROY:
				w32.PostQuitMessage(0)
				return 0
			default:
				return w32.DefWindowProc(window, msg, w, l)
			}
		}),
		ClassName: className,
	})

	title, _ := syscall.UTF16PtrFromString("The Game")
	window := w32.CreateWindow(
		className,
		title,
		w32.WS_OVERLAPPEDWINDOW,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 640, 480,
		0, 0, 0, nil,
	)

	sound, err := initSoundSystem(ds.HWND(window))
	check(err)
	defer sound.close()

	check(sound.preload("assets/music_intro.ogg"))
	check(sound.preload("assets/music_loop.ogg"))
	check(sound.preload("assets/blip.ogg"))
	check(sound.preload("assets/step.ogg"))

	instructions, err := sound.loop("assets/instructions.ogg")
	check(err)
	sound.setSpeed(instructions, 0)

	objectVertexShaderCode, err := dxc.Compile([]byte(`
float4x4 mvp: register(c0);
float4x4 normalTransform: register(c4);

struct input {
	float4 position: POSITION;
	float3 normal: NORMAL;
	float2 uv: TEXCOORD0;
};

struct output {
	float4 position: POSITION;
	float3 normal: NORMAL;
	float2 uv: TEXCOORD0;
	float4 worldPosition: TEXCOORD1;
};

void main(in input IN, out output OUT) {
	OUT.position = mul(IN.position, mvp);
	OUT.normal = mul(float4(IN.normal, 1), normalTransform).xyz;
	OUT.uv = IN.uv;
	OUT.worldPosition = OUT.position;
}
	`), "main", "vs_3_0", dxc.WARNINGS_ARE_ERRORS, 0)
	check(err)

	objectPixelShaderCode, err := dxc.Compile([]byte(`
float4 colorFactor: register(c0);
float4 lightDirection: register(c1);
// lightParameters is (specular strength, specular exponent, ambient strength).
float4 lightParameters: register(c2);

sampler img;

struct input {
	float3 normal: NORMAL;
	float2 uv: TEXCOORD0;
	float4 worldPosition: TEXCOORD1;
};

struct output {
	float4 color: COLOR0;
};

void main(in input IN, out output OUT) {
	// For an explanation of this lighting model, see
	// https://learnopengl.com/Lighting/Basic-Lighting
	float4 lightColor = float4(1, 1, 1, 1);
	float4 objectColor = tex2D(img, IN.uv);
	float3 norm = normalize(IN.normal);
	float3 pos = IN.worldPosition.xyz / IN.worldPosition.w;

	float ambientStrength = lightParameters.z;
	float4 ambient = ambientStrength * lightColor;

	float3 lightDir = -normalize(lightDirection.xyz);
	float diff = max(0, dot(norm, lightDir));
	float4 diffuse = diff * lightColor;

	float specularStrength = lightParameters.x;
	float3 viewPos = float3(0, 0, 0);
	float3 viewDir = normalize(viewPos - pos);
	float3 reflectDir = reflect(-lightDir, norm);
	float spec = pow(max(0, dot(viewDir, reflectDir)), lightParameters.y);
	float4 specular = specularStrength * spec * lightColor;

	OUT.color = min(1, ambient + diffuse + specular) * objectColor * colorFactor;
}
	`), "main", "ps_3_0", dxc.WARNINGS_ARE_ERRORS, 0)
	check(err)

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	createFlags := uint32(d3d9.CREATE_SOFTWARE_VERTEXPROCESSING)
	caps, err := d3d.GetDeviceCaps(d3d9.ADAPTER_DEFAULT, d3d9.DEVTYPE_HAL)
	if err == nil &&
		caps.DevCaps&d3d9.DEVCAPS_HWTRANSFORMANDLIGHT != 0 {
		createFlags = d3d9.CREATE_HARDWARE_VERTEXPROCESSING
	}

	pp := d3d9.PRESENT_PARAMETERS{
		Windowed:      1,
		HDeviceWindow: d3d9.HWND(window),
		// SWAPEFFECT_COPY lets Present define the rectangle.
		SwapEffect: d3d9.SWAPEFFECT_COPY,
		// We use 2048 by 2048 which gets scaled to the monitor resolution on
		// Present.
		BackBufferWidth:        2048,
		BackBufferHeight:       2048,
		BackBufferFormat:       d3d9.FMT_UNKNOWN,
		BackBufferCount:        1,
		EnableAutoDepthStencil: 1,
		AutoDepthStencilFormat: d3d9.FMT_D24X8,
	}

	device, _, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(window),
		createFlags,
		pp,
	)
	check(err)
	defer device.Release()

	objectVertexShader, err := device.CreateVertexShaderFromBytes(objectVertexShaderCode)
	check(err)
	defer objectVertexShader.Release()

	objectPixelShader, err := device.CreatePixelShaderFromBytes(objectPixelShaderCode)
	check(err)
	defer objectPixelShader.Release()

	texturedVertex, err := device.CreateVertexDeclaration([]d3d9.VERTEXELEMENT{
		{Offset: 0, Type: d3d9.DECLTYPE_FLOAT3, Usage: d3d9.DECLUSAGE_POSITION},
		{Offset: 3 * 4, Type: d3d9.DECLTYPE_FLOAT3, Usage: d3d9.DECLUSAGE_NORMAL},
		{Offset: 6 * 4, Type: d3d9.DECLTYPE_FLOAT2, Usage: d3d9.DECLUSAGE_TEXCOORD},
		d3d9.DeclEnd(),
	})
	check(err)
	defer texturedVertex.Release()

	xboxControllerTexture, err := loadTexture(device, "assets/xbox_controller.jpg")
	check(err)
	defer xboxControllerTexture.Release()

	joystickTexture, err := loadTexture(device, "assets/joystick.jpg")
	check(err)
	defer joystickTexture.Release()

	jokerTexture, err := loadTexture(device, "assets/joker.jpg")
	check(err)
	defer jokerTexture.Release()

	levelTexture, err := loadTexture(device, "assets/level.png")
	check(err)
	defer levelTexture.Release()

	jokerModel, err := loadObj("assets/joker.obj")
	check(err)

	levelModel, err := loadObj("assets/level.obj")
	check(err)

	controllerModel, err := loadObj("assets/xbox_controller.obj")
	check(err)

	joystickModel, err := loadObj("assets/joystick.obj")
	check(err)

	vertices := make([]float32, 0, 1024*1024*4)

	addFace := func(obj *obj.File, f obj.FaceVertex, box *aabb) {
		v := obj.Vertices[f.VertexIndex][:3]
		x, y, z := v[0], v[1], v[2]

		if x < box.x.min {
			box.x.min = x
		}
		if y < box.y.min {
			box.y.min = y
		}
		if z < box.z.min {
			box.z.min = z
		}
		if x > box.x.max {
			box.x.max = x
		}
		if y > box.y.max {
			box.y.max = y
		}
		if z > box.z.max {
			box.z.max = z
		}

		vertices = append(vertices, v...)
		vertices = append(vertices, obj.Normals[f.NormalIndex][:3]...)
		if f.TexCoordIndex < 0 {
			vertices = append(vertices, 0, 0)
		} else {
			vertices = append(vertices, obj.TexCoords[f.TexCoordIndex][:2]...)
		}
	}

	addModel := func(obj *obj.File) model {
		var m model
		for _, o := range obj.Objects {
			faces := obj.Faces[o.StartFace:o.EndFace]
			part := modelPart{
				name:        o.Name,
				firstVertex: len(vertices),
				box:         emptyAABB,
			}

			for _, face := range faces {
				for i := 2; i < len(face); i++ {
					addFace(obj, face[0], &part.box)
					addFace(obj, face[i-1], &part.box)
					addFace(obj, face[i], &part.box)
				}
			}

			part.endVertex = len(vertices)
			m = append(m, part)
		}
		return m
	}

	controller3D := addModel(controllerModel)
	joystick3D := addModel(joystickModel)
	joker3D := addModel(jokerModel)
	level3D := addModel(levelModel)

	float32sPerTexturedVertex := 8
	objectBufferSize := uint(len(vertices) * float32sPerTexturedVertex)
	objectBufferStride := uint(float32sPerTexturedVertex * 4)

	objectBuffer, err := device.CreateVertexBuffer(
		objectBufferSize, d3d9.USAGE_WRITEONLY, 0, d3d9.POOL_DEFAULT, 0,
	)
	check(err)
	defer objectBuffer.Release()

	mem, err := objectBuffer.Lock(0, objectBufferSize, d3d9.LOCK_DISCARD)
	check(err)
	mem.SetFloat32s(0, vertices)
	check(objectBuffer.Unlock())

	check(device.SetRenderState(d3d9.RS_CULLMODE, uint32(d3d9.CULL_CCW)))

	drawXBoxController := func(modelTransform m.Mat4) {
		bounds := w32.GetClientRect(window)
		aspect := float32(bounds.Right) / float32(bounds.Bottom)

		check(device.SetVertexDeclaration(texturedVertex))
		check(device.SetVertexShader(objectVertexShader))
		check(device.SetPixelShader(objectPixelShader))
		check(device.SetStreamSource(0, objectBuffer, 0, objectBufferStride))

		colorFactor := m.Vec4{1, 1, 1, 1}
		if gameState == gameStateXBoxController && !input.xboxController.connected {
			xboxBlinkTimer++
			f := float32(math.Sin(float64(xboxBlinkTimer)/10)) + 1
			colorFactor = m.Vec4{1.2 * f, f, f, 1}
		} else {
			xboxBlinkTimer = 0
		}

		check(device.SetPixelShaderConstantF(0, colorFactor[:]))
		check(device.SetPixelShaderConstantF(1, lightDir[:]))
		check(device.SetPixelShaderConstantF(2, []float32{
			specularStrength,
			specularExponent,
			0.1,
			0,
		}))

		// Draw the XBox controller.
		check(device.SetTexture(0, xboxControllerTexture))
		for _, o := range controller3D {
			custom := m.Identity4()

			if o.name == "buttonA" && input.xboxController.buttonADown() ||
				o.name == "buttonB" && input.xboxController.buttonBDown() ||
				o.name == "buttonX" && input.xboxController.buttonXDown() ||
				o.name == "buttonY" && input.xboxController.buttonYDown() {
				custom = m.Translate(0, -0.025, 0)
			}

			if o.name == "buttonRB" && input.xboxController.buttonRBDown() ||
				o.name == "buttonLB" && input.xboxController.buttonLBDown() {
				custom = m.Translate(0, 0, -0.015)
			}

			if o.name == "buttonStart" && input.xboxController.buttonStartDown() ||
				o.name == "buttonBack" && input.xboxController.buttonBackDown() {
				custom = m.Translate(0, -0.025, 0)
			}

			if o.name == "leftAxis" || o.name == "rightAxis" {
				rotationAxis := m.Vec3{
					relativeAxis(input.xboxController.leftYAxis),
					0,
					relativeAxis(input.xboxController.leftXAxis),
				}
				if o.name == "rightAxis" {
					rotationAxis = m.Vec3{
						relativeAxis(input.xboxController.rightYAxis),
						0,
						relativeAxis(input.xboxController.rightXAxis),
					}
				}

				// Rotate about the bottom of the stick.
				x := (o.box.x.min + o.box.x.max) / 2
				y := o.box.y.min + (o.box.y.max-o.box.y.min)*-0.5
				z := (o.box.z.min + o.box.z.max) / 2

				var dy float32
				if o.name == "leftAxis" && input.xboxController.leftAxisDown() ||
					o.name == "rightAxis" && input.xboxController.rightAxisDown() {
					dy = (o.box.y.max - o.box.y.min) * -0.1
				}

				custom = m.Mul4(
					m.Translate(-x, -y, -z),
					m.RotateRightHandAbout(rotationAxis, rotationAxis.Norm()/20),
					m.Translate(x, y+dy, z),
				)
			}

			if o.name == "dpad" && 0 <= input.xboxController.dpad &&
				input.xboxController.dpad <= 31500 {
				base := m.Vec4{-1, 0, 0, 1}
				turns := float32(input.xboxController.dpad) / 36000
				rot := m.RotateLeftHandAbout(m.Vec3{0, 1, 0}, turns)
				rotationAxis := base.MulMat(rot).DropW()

				// Rotate about the bottom of the stick.
				x := (o.box.x.min + o.box.x.max) / 2
				y := o.box.y.min + (o.box.y.max-o.box.y.min)*-0.5
				z := (o.box.z.min + o.box.z.max) / 2

				custom = m.Mul4(
					m.Translate(-x, -y, -z),
					m.RotateRightHandAbout(rotationAxis, 0.03),
					m.Translate(x, y, z),
					m.Translate(0, (o.box.y.max-o.box.y.min)*-0.2, 0),
				)
			}

			if o.name == "buttonLT" || o.name == "buttonRT" {
				value := input.xboxController.leftTrigger
				if o.name == "buttonRT" {
					value = input.xboxController.rightTrigger
				}

				zRange := o.box.z.max - o.box.z.min
				x := (o.box.x.min + o.box.x.max) / 2
				y := o.box.y.max
				z := o.box.z.min
				custom = m.Mul4(
					m.Translate(-x, -y, -z),
					m.RotateLeftHandX(value/20),
					m.Translate(x, y+value*zRange*-0.1, z+value*zRange*0.05),
				)
			}

			finalModelTransform := m.Mul4(
				custom,
				modelTransform,
			)

			normalTransform := finalModelTransform
			normalTransform[3] = 0
			normalTransform[7] = 0
			normalTransform[11] = 0
			normalTransform[12] = 0
			normalTransform[13] = 0
			normalTransform[14] = 0
			normalTransform[15] = 0

			mvp := m.Mul4(
				finalModelTransform,
				m.Perspective(m.DegToRad*80, aspect, 0.1, 1000.0),
			)

			check(device.SetVertexShaderConstantF(0, mvp[:]))
			check(device.SetVertexShaderConstantF(4, normalTransform[:]))

			vertices := vertices[o.firstVertex:o.endVertex]
			triangleCount := uint(len(vertices) / (3 * float32sPerTexturedVertex))
			offset := uint(o.firstVertex / float32sPerTexturedVertex)
			check(device.DrawPrimitive(d3d9.PT_TRIANGLELIST, offset, triangleCount))
		}
	}

	drawJoystick := func(modelTransform m.Mat4) {
		bounds := w32.GetClientRect(window)
		aspect := float32(bounds.Right) / float32(bounds.Bottom)

		check(device.SetVertexDeclaration(texturedVertex))
		check(device.SetVertexShader(objectVertexShader))
		check(device.SetPixelShader(objectPixelShader))
		check(device.SetStreamSource(0, objectBuffer, 0, objectBufferStride))

		colorFactor := m.Vec4{1, 1, 1, 1}
		if input.joystickDevice == nil {
			joystickBlinkTimer++
			f := float32(math.Sin(float64(joystickBlinkTimer)/10)) + 1
			colorFactor = m.Vec4{1.2 * f, f, f, 1}
		} else {
			joystickBlinkTimer = 0
		}

		check(device.SetPixelShaderConstantF(0, colorFactor[:]))
		check(device.SetPixelShaderConstantF(1, []float32{1, -1, 3, 1}))
		check(device.SetPixelShaderConstantF(2, []float32{0.7, 128, 0.1, 0}))

		// Draw the joystick.
		check(device.SetTexture(0, joystickTexture))
		for _, o := range joystick3D {
			custom := m.Identity4()

			if o.name == "stick" {
				rotationAxis := m.Vec3{
					relativeAxis(input.joystick.yAxis),
					0,
					relativeAxis(input.joystick.xAxis),
				}

				// Rotate about the bottom of the stick.
				x := (o.box.x.min + o.box.x.max) / 2
				y := o.box.y.min
				z := (o.box.z.min + o.box.z.max) / 2

				custom = m.Mul4(
					m.Translate(-x, -y, -z),
					m.RotateRightHandAbout(rotationAxis, rotationAxis.Norm()/20),
					m.Translate(x, y, z),
				)
			}

			finalModelTransform := m.Mul4(
				custom,
				modelTransform,
			)

			normalTransform := finalModelTransform
			normalTransform[3] = 0
			normalTransform[7] = 0
			normalTransform[11] = 0
			normalTransform[12] = 0
			normalTransform[13] = 0
			normalTransform[14] = 0
			normalTransform[15] = 0

			mvp := m.Mul4(
				finalModelTransform,
				m.Perspective(m.DegToRad*80, aspect, 0.1, 1000.0),
			)

			check(device.SetVertexShaderConstantF(0, mvp[:]))
			check(device.SetVertexShaderConstantF(4, normalTransform[:]))

			vertices := vertices[o.firstVertex:o.endVertex]
			triangleCount := uint(len(vertices) / (3 * float32sPerTexturedVertex))
			offset := uint(o.firstVertex / float32sPerTexturedVertex)
			check(device.DrawPrimitive(d3d9.PT_TRIANGLELIST, offset, triangleCount))
		}
	}

	updateSound := func() {
		speed := 0.0
		if gameState == gameStateXBoxController {
			x := input.xboxController.leftXAxis
			speed = makeSoundSpeed(float64(relativeAxis(x)))
		}
		sound.setSpeed(instructions, speed)

		check(sound.update())
	}

	render := func() {
		if gameState == gameStateFadingIn {
			var c uint8
			if fadeInColor > 0 {
				c = uint8(fadeInColor)
			}
			check(device.Clear(
				nil,
				d3d9.CLEAR_TARGET|d3d9.CLEAR_ZBUFFER,
				d3d9.ColorRGB(c, c, c),
				1,
				0,
			))
			check(device.Present(nil, nil, 0, nil))
			fadeInColor++
			if fadeInColor >= backgroundGray {
				gameState = gameStateXBoxControllerFlyingIn
			}
		} else if gameState == gameStateXBoxControllerFlyingIn {
			check(device.Clear(
				nil,
				d3d9.CLEAR_TARGET|d3d9.CLEAR_ZBUFFER,
				d3d9.ColorRGB(backgroundGray, backgroundGray, backgroundGray),
				1,
				0,
			))

			check(device.BeginScene())
			scale := float32(controllerFlyTime * controllerFlyTime)
			rotation := controllerFlyTime * (10 + finalControllerXRotation)
			dz := float32((1 - controllerFlyTime) * 100)
			modelTransform := m.Mul4(
				m.Scale(scale, scale, scale),
				m.RotateRightHandX(float32(rotation)),
				m.Translate(0, 0, finalControllerZ+dz),
			)
			drawXBoxController(modelTransform)
			check(device.EndScene())
			check(device.Present(nil, nil, 0, nil))

			controllerFlyTime += 0.0025
			if controllerFlyTime >= 1 {
				controllerFlyTime = 1
				gameState = gameStateXBoxController
			}
		} else if gameState == gameStateXBoxController {
			check(device.Clear(
				nil,
				d3d9.CLEAR_TARGET|d3d9.CLEAR_ZBUFFER,
				d3d9.ColorRGB(backgroundGray, backgroundGray, backgroundGray),
				1,
				0,
			))

			check(device.BeginScene())
			modelTransform := m.Mul4(
				m.RotateRightHandX(finalControllerXRotation),
				m.RotateRightHandX(controllerXRotation),
				m.RotateLeftHandY(controllerYRotation),
				m.Translate(0, 0, finalControllerZ),
			)
			drawXBoxController(modelTransform)
			check(device.EndScene())
			check(device.Present(nil, nil, 0, nil))

			controllerXRotation += input.xboxController.rightYAxis / 200
			if controllerXRotation > 0.1 {
				controllerXRotation = 0.1
			}
			if controllerXRotation < -0.1 {
				controllerXRotation = -0.1
			}

			controllerYRotation -= float32(
				ease.InQuint(float64(input.xboxController.rightXAxis)) / 100)
			if controllerYRotation > 1 {
				controllerYRotation--
			}
			if controllerYRotation < -1 {
				controllerYRotation++
			}

			if input.xboxController.buttonADown() {
				specularStrength -= 0.01
				if specularStrength < 0.05 {
					specularStrength = 0.05
				}
			}
			if input.xboxController.buttonBDown() {
				specularStrength += 0.01
				if specularStrength > 0.95 {
					specularStrength = 0.95
				}
			}
			if input.xboxController.buttonXDown() {
				specularExponent /= 1.05
				if specularExponent < 2 {
					specularExponent = 2
				}
			}
			if input.xboxController.buttonYDown() {
				specularExponent *= 1.05
				if specularExponent > 128 {
					specularExponent = 128
				}
			}
			if input.xboxController.buttonStartDown() {
				specularStrength = 0.5
				specularExponent = 16
			}
			if input.xboxController.buttonRBDown() {
				lightDir = m.Vec4{1, -1, 1, 0}
			}
			if input.xboxController.buttonBackDown() {
				controllerXRotation = 0
				controllerYRotation = 0
			}
			if input.xboxController.dpad < 0xFFFF {
				degress := float64(input.xboxController.dpad) / 100
				dz, dx := math.Sincos(m.DegToRad * (90 - degress))
				lightDir = m.Vec4{float32(-dx), -2, float32(-dz), 0}
			}

			if input.xboxController.buttons != lastButtonState {
				lastButtonState = input.xboxController.buttons
				pushButtonState(input.xboxController.buttons)
				equal := func() bool {
					for i := range desiredButtonStates {
						if desiredButtonStates[i] != lastButtonStates[i] {
							return false
						}
					}
					return true
				}()
				if equal {
					gameState = gameStateTransitionToJoystick
					sound.stop(instructions)

					intro, err := sound.play("assets/music_intro.ogg")
					check(err)
					_, err = sound.queueLoopAfter(intro, "assets/music_loop.ogg")
					check(err)
				}
			}
		} else if gameState == gameStateTransitionToJoystick {
			check(device.Clear(
				nil,
				d3d9.CLEAR_TARGET|d3d9.CLEAR_ZBUFFER,
				d3d9.ColorRGB(backgroundGray, backgroundGray, backgroundGray),
				1,
				0,
			))

			check(device.BeginScene())

			xboxControllerTransform := m.Mul4(
				m.ScaleUniform(gamepadScale),
				m.RotateRightHandX(finalControllerXRotation),
				m.RotateRightHandX(controllerXRotation),
				m.RotateLeftHandY(controllerYRotation),
				m.Translate(0, 0, finalControllerZ),
			)
			drawXBoxController(xboxControllerTransform)

			joystickTransform := m.Mul4(
				m.ScaleUniform(0.5),
				m.ScaleUniform(joystickScale),
				m.RotateRightHandX(0.05),
				m.RotateRightHandY(joystickYRotation),
				m.Translate(0, -0.5, finalControllerZ),
			)
			drawJoystick(joystickTransform)

			check(device.EndScene())
			check(device.Present(nil, nil, 0, nil))

			joystickYRotation += joystickYRotationSpeed

			gamepadScale -= joystickScaleSpeed
			if gamepadScale <= 0 {
				gamepadScale = 0

				joystickScale += joystickScaleSpeed
				if joystickScale >= 1 {
					joystickScale = 1
					gameState = gameStateJoystickRotating
				}
			}
		} else if gameState == gameStateJoystickRotating {
			check(device.Clear(
				nil,
				d3d9.CLEAR_TARGET|d3d9.CLEAR_ZBUFFER,
				d3d9.ColorRGB(backgroundGray, backgroundGray, backgroundGray),
				1,
				0,
			))

			check(device.BeginScene())
			joystickTransform := m.Mul4(
				m.ScaleUniform(0.5),
				m.ScaleUniform(joystickScale),
				m.RotateRightHandX(0.05),
				m.RotateRightHandY(joystickYRotation),
				m.Translate(0, -0.5, finalControllerZ),
			)
			drawJoystick(joystickTransform)

			check(device.EndScene())
			check(device.Present(nil, nil, 0, nil))

			joystickYRotation += joystickYRotationSpeed

			if input.joystick.buttonDown != [8]bool{} {
				gameState = gameStateJoystickShrinking
			}
		} else if gameState == gameStateJoystickShrinking {
			check(device.Clear(
				nil,
				d3d9.CLEAR_TARGET|d3d9.CLEAR_ZBUFFER,
				d3d9.ColorRGB(backgroundGray, backgroundGray, backgroundGray),
				1,
				0,
			))

			check(device.BeginScene())
			joystickTransform := m.Mul4(
				m.ScaleUniform(0.5),
				m.ScaleUniform(joystickScale),
				m.RotateRightHandX(0.05),
				m.RotateRightHandY(joystickYRotation),
				m.Translate(0, -0.5, finalControllerZ),
			)
			drawJoystick(joystickTransform)

			check(device.EndScene())
			check(device.Present(nil, nil, 0, nil))

			joystickYRotation += joystickYRotationSpeed
			joystickScale -= joystickScaleSpeed

			if joystickScale <= 0 {
				gameState = gameStatePlayingLevel
			}
		} else if gameState == gameStatePlayingLevel {
			check(device.Clear(
				nil,
				d3d9.CLEAR_TARGET|d3d9.CLEAR_ZBUFFER,
				d3d9.ColorRGB(backgroundGray, backgroundGray, backgroundGray),
				1,
				0,
			))

			check(device.BeginScene())

			view := m.LookAt(cameraPos, jokerPos, m.Vec3{0, 1, 0})

			bounds := w32.GetClientRect(window)
			aspect := float32(bounds.Right) / float32(bounds.Bottom)

			check(device.SetVertexDeclaration(texturedVertex))
			check(device.SetVertexShader(objectVertexShader))
			check(device.SetPixelShader(objectPixelShader))
			check(device.SetStreamSource(0, objectBuffer, 0, objectBufferStride))
			lightColor := []float32{levelColor, levelColor, levelColor, 1}
			check(device.SetPixelShaderConstantF(0, lightColor))
			check(device.SetPixelShaderConstantF(1, []float32{-0.7, -4, 1, 1}))
			check(device.SetPixelShaderConstantF(2, []float32{0.1, 2, 0.6, 0}))

			check(device.SetTexture(0, levelTexture))
			for _, o := range level3D {
				normalTransform := m.Identity4()

				mvp := m.Mul4(
					view,
					m.Perspective(m.DegToRad*fieldOfView, aspect, 0.1, 1000.0),
				)

				check(device.SetVertexShaderConstantF(0, mvp[:]))
				check(device.SetVertexShaderConstantF(4, normalTransform[:]))

				vertices := vertices[o.firstVertex:o.endVertex]
				triangleCount := uint(len(vertices) / (3 * float32sPerTexturedVertex))
				offset := uint(o.firstVertex / float32sPerTexturedVertex)
				check(device.DrawPrimitive(d3d9.PT_TRIANGLELIST, offset, triangleCount))
			}

			// Draw the joker.
			check(device.SetPixelShaderConstantF(1, []float32{0, -1, 1, 1}))
			check(device.SetPixelShaderConstantF(2, []float32{0.7, 128, 0.2, 0}))
			check(device.SetTexture(0, jokerTexture))
			for _, o := range joker3D {
				custom := m.Identity4()

				if o.name == "leftLeg" || o.name == "rightLeg" ||
					o.name == "leftArm" || o.name == "rightArm" ||
					o.name == "leftHand" || o.name == "rightHand" {

					rot := jokerLimbRot
					if o.name == "leftLeg" ||
						o.name == "rightArm" || o.name == "rightHand" {
						rot = -rot
					}

					ref := jokerModel.FindObject("refArmJoint")
					if o.name == "leftLeg" || o.name == "rightLeg" {
						ref = jokerModel.FindObject("refLegJoint")
					}

					joint := jokerModel.Vertices[ref.StartVertex]

					x, y, z := joint[0], joint[1], joint[2]

					custom = m.Mul4(
						m.Translate(-x, -y, -z),
						m.RotateLeftHandX(0.16*float32(math.Sin(m.TurnsToRad*rot))),
						m.Translate(x, y, z),
					)
				}

				model := m.Mul4(
					custom,
					m.RotateRightHandY(jokerRot-jokerBaseRot),
					m.TranslateV(jokerPos),
				)

				normalTransform := model
				normalTransform[3] = 0
				normalTransform[7] = 0
				normalTransform[11] = 0
				normalTransform[12] = 0
				normalTransform[13] = 0
				normalTransform[14] = 0
				normalTransform[15] = 0

				mvp := m.Mul4(
					model,
					view,
					m.Perspective(m.DegToRad*fieldOfView, aspect, 0.1, 1000.0),
				)

				check(device.SetVertexShaderConstantF(0, mvp[:]))
				check(device.SetVertexShaderConstantF(4, normalTransform[:]))

				vertices := vertices[o.firstVertex:o.endVertex]
				triangleCount := uint(len(vertices) / (3 * float32sPerTexturedVertex))
				offset := uint(o.firstVertex / float32sPerTexturedVertex)
				check(device.DrawPrimitive(d3d9.PT_TRIANGLELIST, offset, triangleCount))
			}

			check(device.EndScene())
			check(device.Present(nil, nil, 0, nil))

			joyX := relativeAxis(input.joystick.xAxis)
			joyY := relativeAxis(input.joystick.yAxis)

			xboxX := relativeAxis(input.xboxController.leftXAxis)
			xboxY := relativeAxis(input.xboxController.leftYAxis)

			xAxis := joyX
			yAxis := joyY

			if abs(xboxY) > abs(joyY) {
				yAxis = xboxY
			}
			if abs(xboxX) > abs(joyX) {
				xAxis = xboxX
			}

			targetJokerSpeed := float64(-yAxis) * 0.05

			if jokerSpeed < targetJokerSpeed {
				jokerSpeed += jokerAcceleration
				if jokerSpeed > targetJokerSpeed {
					jokerSpeed = targetJokerSpeed
				}
			}

			if jokerSpeed > targetJokerSpeed {
				jokerSpeed -= jokerAcceleration
				if jokerSpeed < targetJokerSpeed {
					jokerSpeed = targetJokerSpeed
				}
			}

			lastLimbRot := jokerLimbRot

			if yAxis == 0 {
				if jokerSpeed > 0 {
					jokerSpeed -= jokerAcceleration
					if jokerSpeed < 0 {
						jokerSpeed = 0
					}
				}
				if jokerSpeed < 0 {
					jokerSpeed += jokerAcceleration
					if jokerSpeed > 0 {
						jokerSpeed = 0
					}
				}

				// Limb rotations of 0.0, 0.5 and 1.0 are all OK, as they are
				// all the standing position.
				if jokerLimbRot < 0.25 {
					// Go from (0.0, 0.25) down to 0.0.
					jokerLimbRot -= maxJokerSpeed * jokerSpeedLimbRatio
					if jokerLimbRot < 0 {
						jokerLimbRot = 0
					}
				} else if 0.25 < jokerLimbRot && jokerLimbRot < 0.5 {
					// Go from (0.25,  0.5) up to 0.5.
					jokerLimbRot += maxJokerSpeed * jokerSpeedLimbRatio
					if jokerLimbRot >= 0.5 {
						jokerLimbRot = 0
					}
				} else if 0.5 < jokerLimbRot && jokerLimbRot < 0.75 {
					// Go from (0.5,  0.75) down to 0.5.
					jokerLimbRot -= maxJokerSpeed * jokerSpeedLimbRatio
					if jokerLimbRot <= 0.5 {
						jokerLimbRot = 0
					}
				} else if 0.75 < jokerLimbRot {
					// Go from (0.75,  1.0) up to 1.0.
					jokerLimbRot += maxJokerSpeed * jokerSpeedLimbRatio
					if jokerLimbRot >= 1 {
						jokerLimbRot = 0
					}
				} else {
					jokerLimbRot = 0
				}
			}

			floorHeightsAt := func(x, z float32) [4]float32 {
				const collisionMargin = 0.25
				x0 := x - collisionMargin
				x1 := x + collisionMargin
				z0 := z - collisionMargin
				z1 := z + collisionMargin
				return [4]float32{
					float32(floorHeightAt(x0, z0)),
					float32(floorHeightAt(x0, z1)),
					float32(floorHeightAt(x1, z0)),
					float32(floorHeightAt(x1, z1)),
				}
			}

			collides := func(x, y, z float32) bool {
				heights := floorHeightsAt(x, z)
				for _, h := range heights {
					if h > y {
						return true
					}
				}
				return false
			}

			jokerRot += -xAxis * 0.006

			if jokerSpeed != 0 {
				if yAxis != 0 {
					jokerLimbRot += jokerSpeed * jokerSpeedLimbRatio
				}

				sin, cos := math.Sincos(float64(m.TurnsToRad * jokerRot))
				dx := float32(jokerSpeed * cos)
				dz := float32(jokerSpeed * sin)

				collidesX := collides(jokerPos[0]+dx, jokerPos[1], jokerPos[2])
				collidesZ := collides(jokerPos[0], jokerPos[1], jokerPos[2]+dz)
				if !collidesZ {
					jokerPos[2] += dz
				}
				if !collidesX {
					jokerPos[0] += dx
				}

			}

			wantsToJump :=
				!lastJoystickState.buttonDown[0] && input.joystick.buttonDown[0] ||
					!lastXBoxState.buttonADown() && input.xboxController.buttonADown()

			if !lastJoystickState.buttonDown[1] && input.joystick.buttonDown[1] ||
				!lastXBoxState.buttonYDown() && input.xboxController.buttonYDown() {
				cameraInCorner = !cameraInCorner
			}

			var targetCameraPos m.Vec3

			if cameraInCorner {
				cornerIndex := int(input.joystick.dpad) / 4500
				if cornerIndex >= len(cameraCornerPositions) {
					cornerIndex = int(input.xboxController.dpad) / 4500
				}
				if cornerIndex < len(cameraCornerPositions) {
					cameraTargetCorner = cameraCornerPositions[cornerIndex]
				}
				targetCameraPos = cameraTargetCorner
			} else {
				dirZ, dirX := math.Sincos(float64(m.TurnsToRad * jokerRot))
				maxCamX := float32(len(floorHeights[0]) - 1)
				minCamZ := -float32(len(floorHeights) - 1)
				targetCameraPos = m.Vec3{
					max(1, min(maxCamX, jokerPos[0]-5*float32(dirX))),
					4,
					min(-1, max(minCamZ, jokerPos[2]-5*float32(dirZ))),
				}
			}

			cameraPos = cameraPos.MulScalar(0.95).Add(targetCameraPos.MulScalar(0.05))

			lastJoystickState = input.joystick
			lastXBoxState = input.xboxController

			playStep := func() {
				if stepCoolDown > 0 {
					return
				}
				s, err := sound.play("assets/step.ogg")
				check(err)
				sound.setSpeed(s, 0.75+1.5*rand.Float64())
				stepCoolDown = 10
			}
			if stepCoolDown > 0 {
				stepCoolDown--
			}

			onGround := false
			jokerSpeedY += gravity
			jokerPos[1] += jokerSpeedY
			if collides(jokerPos[0], jokerPos[1], jokerPos[2]) {
				onGround = true
				jokerPos[1] = float32(int(jokerPos[1]))
				jokerSpeedY = 0

				if collides(jokerPos[0], jokerPos[1], jokerPos[2]) {
					jokerPos[1] = float32(int(jokerPos[1]) + 1)
				}

				if wantsToJump {
					jokerSpeedY = jokerJumpSpeed
					s, err := sound.play("assets/blip.ogg")
					check(err)
					sound.setSpeed(s, 1+0.5*rand.Float64())
				}
			}

			if onGround && !wasOnGround {
				playStep()
			}
			wasOnGround = onGround

			jokerLimbRot = norm01(jokerLimbRot)

			if onGround &&
				(lastLimbRot < 0.25 && jokerLimbRot >= 0.25 ||
					lastLimbRot < 0.75 && jokerLimbRot >= 0.75) {
				playStep()
			}

			levelColor = max(1, levelColor*0.95)
		}
	}

	if fullscreen {
		style := w32.GetWindowLong(window, w32.GWL_STYLE)
		var monitorInfo w32.MONITORINFO
		monitor := w32.MonitorFromWindow(window, w32.MONITOR_DEFAULTTOPRIMARY)
		var windowed w32.WINDOWPLACEMENT
		if w32.GetWindowPlacement(window, &windowed) &&
			w32.GetMonitorInfo(monitor, &monitorInfo) {
			w32.SetWindowLong(
				window,
				w32.GWL_STYLE,
				style & ^w32.WS_OVERLAPPEDWINDOW,
			)
			w32.SetWindowPos(
				window,
				0,
				int(monitorInfo.RcMonitor.Left),
				int(monitorInfo.RcMonitor.Top),
				int(monitorInfo.RcMonitor.Right-monitorInfo.RcMonitor.Left),
				int(monitorInfo.RcMonitor.Bottom-monitorInfo.RcMonitor.Top),
				w32.SWP_NOOWNERZORDER|w32.SWP_FRAMECHANGED,
			)
		}
		w32.ShowCursor(false)
	}

	w32.ShowWindow(window, syscall.SW_SHOWNORMAL)

	msg := w32.MSG{Message: w32.WM_QUIT + 1}
	for msg.Message != w32.WM_QUIT {
		if w32.PeekMessage(&msg, 0, 0, 0, w32.PM_REMOVE) {
			if msg.Message == w32.WM_QUIT {
				break
			}
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		} else {
			input.update()
			updateSound()
			render()
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

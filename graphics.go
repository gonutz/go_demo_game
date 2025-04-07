package main

import (
	"bytes"
	"image"
	"image/draw"
	"math"

	"github.com/gonutz/d3d9"
	"github.com/gonutz/obj"
)

type model []modelPart

// modelPart is a 3D modelPart with some meta data.
type modelPart struct {
	name        string
	firstVertex int
	endVertex   int
	box         aabb
}

type minMax struct {
	min float32
	max float32
}

var emptyMinMax = minMax{
	min: float32(math.Inf(1)),
	max: float32(math.Inf(-1)),
}

type aabb struct {
	x, y, z minMax
}

var emptyAABB = aabb{
	x: emptyMinMax,
	y: emptyMinMax,
	z: emptyMinMax,
}

func color(c uint32) float32 {
	return math.Float32frombits(c)
}

func readImage(data []byte) (*image.RGBA, error) {
	rgba, err := decodeRGBA(data)
	if err != nil {
		return nil, err
	}

	// Swap red and blue channels.
	for i := 0; i < len(rgba.Pix); i += 4 {
		rgba.Pix[i], rgba.Pix[i+2] = rgba.Pix[i+2], rgba.Pix[i]
	}

	return rgba, nil
}

func decodeRGBA(data []byte) (*image.RGBA, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if rgba, ok := img.(*image.RGBA); ok {
		return rgba, nil
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
	return rgba, nil
}

func loadTexture(device *d3d9.Device, path string) (*d3d9.Texture, error) {
	data, err := assetFiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	img, err := readImage(data)
	if err != nil {
		return nil, err
	}

	texture, err := device.CreateTexture(
		uint(img.Bounds().Dx()),
		uint(img.Bounds().Dy()),
		1,
		0,
		d3d9.FMT_A8R8G8B8,
		d3d9.POOL_MANAGED,
		0,
	)
	if err != nil {
		return nil, err
	}

	r, err := texture.LockRect(0, nil, d3d9.LOCK_DISCARD)
	if err != nil {
		return nil, err
	}
	r.SetAllBytes(img.Pix, img.Stride)
	err = texture.UnlockRect(0)
	if err != nil {
		return nil, err
	}

	return texture, nil
}

func loadObj(path string) (*obj.File, error) {
	data, err := assetFiles.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return obj.Decode(bytes.NewReader(data))

}

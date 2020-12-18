package image

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

type Image struct {
	i image.Image
	c image.Config
}

func Default() *Image {
	ins := &Image{}
	return ins
}

func (img *Image) ReadFile(path string) error {
	fb, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	img.c, img.i, err = img.parse(fb)
	if err != nil {
		return err
	}
	return nil
}

func (img *Image) parse(b []byte) (c image.Config, i image.Image, err error) {
	c, f, err := image.DecodeConfig(bytes.NewReader(b))
	if err != nil {
		return
	}
	switch f {
	case "jpeg":
		i, err = jpeg.Decode(bytes.NewReader(b))
	case "png":
		i, err = png.Decode(bytes.NewReader(b))
	}
	return
}

func (img *Image) HollowOut(r *Rect) {
	w, h := img.c.Width, img.c.Height
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			if !((r.X <= i && i <= r.X+r.W && r.Y+r.CircleRadius <= j && j <= r.Y+r.H-r.CircleRadius) ||
				(r.X+r.CircleRadius <= i && i <= r.X+r.W-r.CircleRadius && r.Y <= j && j <= r.Y+r.H) ||
				(r.X < i && i < r.X+r.CircleRadius && r.Y < j && j < r.Y+r.CircleRadius &&
					Pow(r.X+r.CircleRadius-i)+Pow(r.Y+r.CircleRadius-j) <= Pow(r.CircleRadius)) ||
				(r.X+r.W-r.CircleRadius < i && i < r.X+r.W && r.Y < j && j < r.Y+r.CircleRadius &&
					Pow(r.CircleRadius+i-r.X-r.W)+Pow(r.Y+r.CircleRadius-j) <= Pow(r.CircleRadius)) ||
				(r.X < i && i < r.X+r.CircleRadius && r.Y+r.H-r.CircleRadius < j && j < r.Y+r.H &&
					Pow(r.X+r.CircleRadius-i)+Pow(r.CircleRadius+j-r.Y-r.H) <= Pow(r.CircleRadius)) ||
				(r.X+r.W-r.CircleRadius < i && i < r.X+r.W && r.Y+r.H-r.CircleRadius < j && j < r.Y+r.H &&
					Pow(r.CircleRadius+i-r.X-r.W)+Pow(r.CircleRadius+j-r.Y-r.H) <= Pow(r.CircleRadius))) {
				p := img.i.At(i, j)
				r, g, b, a := p.RGBA()
				p2 := color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(a >> 8),
				}
				rgba.SetRGBA(i, j, p2)
			}
		}
	}
	img.i = rgba
}

func (img *Image) Save(format, path string) error {
	ofile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer ofile.Close()

	switch format {
	case "jpeg":
		err = jpeg.Encode(ofile, img.i, nil)
	case "png":
		err = png.Encode(ofile, img.i)
	}
	return err
}

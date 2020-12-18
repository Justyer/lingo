package image

import (
	"testing"
)

func TestImage(t *testing.T) {
	img := Default()
	err := img.ReadFile("in.jpg")
	if err != nil {
		t.Fatal(err)
	}
	img.HollowOut(&Rect{
		X:            (720 - 16*42) / 2,
		Y:            150,
		W:            16*42 - 1,
		H:            9*42 - 1,
		CircleRadius: 20,
	})
	img.Save("png", "out.png")
}

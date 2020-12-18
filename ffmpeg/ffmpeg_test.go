package ffmpeg

import (
	"fmt"
	"testing"
)

func TestFFmpeg(t *testing.T) {
	ff := Default()
	err := ff.Exec(`-v error -i download.mp4 -filter_complex scale=720:-1 out.mp4`)
	fmt.Println("err", err)
}

func TestFilter(t *testing.T) {
	ff := Default()
	ff.logMute = true
	i0 := &MediaFile{FileName: "src.mp4"}
	i1 := &MediaFile{FileName: "logo.png"}
	ff.AddInput(i0)
	ff.AddInput(i1)
	f0 := ff.AddFilter("scale_f", "scale=-1:720", i0.V())
	ff.AddFilter("ovl", "overlay=0:0", f0.S(), i1.V())
	ff.SetOutputFile("dest.mp4")
	err := ff.Do()
	if err != nil {
		t.Error(err)
	}
}

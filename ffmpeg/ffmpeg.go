package ffmpeg

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type FFmpeg struct {
	Cmd     string
	Inputs  []*MediaFile
	Output  string
	Filters []*Filter

	notCover bool
	logMute  bool
}

func Default() *FFmpeg {
	ins := &FFmpeg{
		Cmd: "ffmpeg",
	}
	return ins
}

func (ff *FFmpeg) AddInput(input *MediaFile) {
	input.index = len(ff.Inputs)
	ff.Inputs = append(ff.Inputs, input)
}

func (ff *FFmpeg) SetOutputFile(output string) {
	ff.Output = output
}

func (ff *FFmpeg) AddFilter(alias, content string, fls ...string) (f *Filter) {
	f = &Filter{
		Alias:   alias,
		Content: content,
		Inputs:  fls,
	}
	ff.Filters = append(ff.Filters, f)
	return
}

func (ff *FFmpeg) Do() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	var param []string

	if !ff.notCover {
		param = append(param, "-y")
	}

	// 处理input
	for _, input := range ff.Inputs {
		if input.StartTime != "" && input.Duration != 0 {
			param = append(param, "-ss", input.StartTime, "-t", strconv.Itoa(input.Duration))
		}
		param = append(param, "-i", input.FileName)
	}

	// 处理filter
	if len(ff.Filters) > 0 {
		param = append(param, "-filter_complex")

		var fls []string
		flen := len(ff.Filters)
		for i := 0; i < flen; i++ {
			var flr string
			if i+1 != flen {
				flr = fmt.Sprintf("%s%s[%s]", strings.Join(ff.Filters[i].Inputs, ""), ff.Filters[i].Content, ff.Filters[i].Alias)
			} else {
				flr = fmt.Sprintf("%s%s", strings.Join(ff.Filters[i].Inputs, ""), ff.Filters[i].Content)
			}
			fls = append(fls, flr)
		}
		param = append(param, strings.Join(fls, ","))
	}

	// 处理output
	param = append(param, ff.Output)

	return ff.execute(param)
}

func (ff *FFmpeg) Exec(cmd string) error {
	return ff.execute(strings.Split(cmd, " "))
}

func (ff *FFmpeg) execute(args []string) (err error) {
	cmd := exec.Command(ff.Cmd, args...)
	fmt.Println(cmd.String())
	var stdBuf, errBuf bytes.Buffer
	cmd.Stdout = &stdBuf
	cmd.Stderr = &errBuf
	err = cmd.Run()
	if err != nil {
		err = errors.New(string(errBuf.Bytes()))
		return
	}
	if !ff.logMute {
		log.Println(string(errBuf.Bytes()))
	}
	return
}

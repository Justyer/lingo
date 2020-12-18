package ffmpeg

import "fmt"

// 滤镜
type Filter struct {
	Alias   string   // 别名
	Content string   // 参数
	Inputs  []string // 输入的流
}

func (f *Filter) S() string {
	return fmt.Sprintf("[%s]", f.Alias)
}

// 媒体文件
type MediaFile struct {
	index     int    // 输入顺序
	StartTime string // 开始时间
	Duration  int    // 截取时长
	FileName  string // 文件位置
}

// 获取音频流
func (m *MediaFile) A() string {
	return fmt.Sprintf("[%d:a]", m.index)
}

// 获取视频流
func (m *MediaFile) V() string {
	return fmt.Sprintf("[%d:v]", m.index)
}

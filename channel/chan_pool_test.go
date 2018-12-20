package chl

import (
	"fmt"
	"testing"
)

func TestChanPool(t *testing.T) {

	// 实例化
	p := NewChanPool()

	// 设置channel缓冲区大小
	p.SetPipeChanBufferCount(1000)

	// 添加
	p.Push(1)
	p.Push(2)
	p.Push(3)

	// 弹出
	fmt.Println(p.Pop())

	// channel数量
	fmt.Println(p.Len())

	// channelpool状态
	p.Status()

	// 关闭channel
	p.Close()

}

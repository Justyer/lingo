package chl

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {

	// 实例化
	p := NewChannel()

	// 设置channel缓冲区大小
	p.SetChanBufferCount(1000)

	// 初始化channel信息
	p.InitChannel()

	// 添加
	p.Push(1)
	p.Push(2)
	p.Push(3)

	// 弹出
	fmt.Println(p.Pop())

	// channel剩余数量
	fmt.Println(p.Len())

	// channel容量
	fmt.Println(p.Cap())

	// 弹出所有channel中数据
	for c := range p.Chan() {
		fmt.Println(c)
	}

	// 关闭channel
	p.Close()

}

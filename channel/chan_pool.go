package chl

import (
	"fmt"
	"sync"
)

type ChanPool struct {
	chans           []*Channel
	pipeBufferCount int
	lock            *sync.Mutex
}

func NewChanPool() *ChanPool {
	return &ChanPool{
		lock: &sync.Mutex{},
	}
}

func (self *ChanPool) SetPipeChanBufferCount(c int) {
	self.pipeBufferCount = c
}

func (self *ChanPool) newAndPush(obj interface{}) {
	p := NewChannel()
	p.SetChanBufferCount(self.pipeBufferCount)
	p.InitChannel()
	self.chans = append(self.chans, p)
	p.Push(obj)
}

func (self *ChanPool) Push(obj interface{}) {
	switch {
	case self.Len() == 0:
		// 如果chans中没有可用channel，就新建一个，push进去
		self.newAndPush(obj)
	case self.Len() > 0:
		self.lock.Lock()
		defer self.lock.Unlock()
		for i := 0; i < self.Len(); i++ {
			if self.chans[i].Len() < self.chans[i].Cap() {
				// 如果chans中有可用的channel，并且未满，就push进去
				self.chans[i].Push(obj)
				return
			}
		}
		// 如果chans中有可用的channel，但是满了，就新建一个，push进去
		self.newAndPush(obj)
	}
}

func (self *ChanPool) Pop() interface{} {
	switch {
	case self.Len() == 0:
		// 如果chans中没有可用channel，返回nil
		return nil
	case self.Len() > 0:
		self.lock.Lock()
		defer self.lock.Unlock()
		for i := 0; i < self.Len(); i++ {
			if self.chans[i].Len() > 0 {
				// 如果chans中有可用channel，pop一个
				return self.chans[i].Pop()
			}
		}
	}
	// 如果chans中有可用的channel，但都为空，则返回nil
	return nil
}

func (self *ChanPool) Len() int {
	return len(self.chans)
}

func (self *ChanPool) Status() {
	for i := 0; i < self.Len(); i++ {
		p := self.chans[i]
		fmt.Printf("chan_%d: %#v, %d, %d\n", i, p, p.Len(), p.Cap())
	}
}

func (self *ChanPool) Close() {
	for i := 0; i < self.Len(); i++ {
		self.chans[i].Close()
	}
}

package chl

type Channel struct {
	chl         chan interface{}
	bufferCount int
}

func NewChannel() *Channel {
	return &Channel{}
}

func (self *Channel) SetChanBufferCount(c int) {
	self.bufferCount = c
}

func (self *Channel) InitChannel() {
	self.chl = make(chan interface{}, self.bufferCount)
}

func (self *Channel) Push(obj interface{}) {
	self.chl <- obj
}

func (self *Channel) Pop() interface{} {
	if self.Len() == 0 {
		return nil
	}
	return <-self.chl
}

func (self *Channel) Len() int {
	return len(self.chl)
}

func (self *Channel) Cap() int {
	return cap(self.chl)
}

func (self *Channel) Chan() chan interface{} {
	return self.chl
}

func (self *Channel) Close() {
	close(self.chl)
}

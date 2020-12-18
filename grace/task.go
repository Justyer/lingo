package grace

import (
	"os"
	"os/signal"
	"syscall"
)

type BaseTask struct {
	Resume bool
}

func (bt *BaseTask) Grace() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	bt.Resume = true
	go func() {
		<-sigChan
		bt.Resume = false
	}()
}

func (bt *BaseTask) GraceCallback(f func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	bt.Resume = true
	go func() {
		<-sigChan
		bt.Resume = false
		f()
	}()
}

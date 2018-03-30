package teststub

import "fmt"

type AAA struct {
}

type AAAer interface {
	sendMsg(a uint32)
	dumpMsg()
}

func (self *AAA) sendMsg(a uint32) {
	fmt.Printf("AAA sendMsg %d\n", a)
}

func (self *AAA) dumpMsg() {
	fmt.Printf("AAA dumpMsg\n")
}

package teststub

import (
	"fmt"
	"testing"
)

type AAAStub struct {
	AAA
}

func (self *AAAStub) sendMsg(a uint32) {
	fmt.Printf("AAAStub sendMsg %d\n", a)
	self.AAA.sendMsg(a)
}

func TestAAA(t *testing.T) {
	var a AAAer

	a = &AAAStub{}

	a.sendMsg(10)
}

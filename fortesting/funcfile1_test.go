package main

import "testing"

func TestFuncAAAA1(t *testing.T) {
	t.Logf("TestFuncAAAA1 enter")
}

func TestFuncAAAA2(t *testing.T) {
	t.Logf("TestFuncAAAA2 enter")
	t.Skip("TestFuncAAAA2 Skip")
	t.Logf(" is skiped")
}

func TestFuncAAAA3(t *testing.T) {
	t.Logf("TestFuncAAAA3 enter")
}

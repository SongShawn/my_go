package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain enter..")
	m.Run()
}

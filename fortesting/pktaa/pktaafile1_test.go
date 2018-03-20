package pktaa

import (
	"fmt"
	"os"
	"testing"
)

func TestAA1(t *testing.T) {
	t.Logf("TestAA1 enter")
}

func TestMain(m *testing.M) {
	fmt.Println("pktaa TestMain enter")
	os.Exit(m.Run())
}

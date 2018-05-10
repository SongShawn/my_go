package pkg

import "fmt"

func UsedForTestMoneyPatch() {
	fmt.Println("this is old.")
}

func PatchFunc() {
	fmt.Println("this is patch")
}

type AAAA struct {
}

func (a *AAAA) UseForPatch() {
	fmt.Println("AAAA old")
}

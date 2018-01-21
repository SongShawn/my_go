package main

import (
	"fmt"
	"stringfactory"
)

func main() {
	strFac := stringfactory.New("", 4, stringfactory.Char_a_z)

	for str := strFac.GetFirstString(); str != ""; str = strFac.GetNextString() {
		fmt.Println(str)
	}
}

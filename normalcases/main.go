package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	name := []byte("指挥官")

	fmt.Println(hex.EncodeToString(name))
}

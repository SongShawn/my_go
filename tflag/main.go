package main

import (
	"flag"
	"fmt"
)

var (
	xx string
)

func init() {
	flag.StringVar(&xx, "xx", "", "xx help")
	flag.Parse()
}

func main() {
	fmt.Println("main")
}

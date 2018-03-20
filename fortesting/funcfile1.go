package main

import "fmt"

func FuncAAAA1() {
	fmt.Println("FuncAAAA1 enter")
	FuncAAAA2()
}

func FuncAAAA2() {
	fmt.Println("FuncAAAA2 enter")
}

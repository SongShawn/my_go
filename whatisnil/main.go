package main

import "fmt"

func checkEqualSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func main() {
	var varSlice []string
	var varInter interface{}

	fmt.Printf("%T %T %T\n", nil, varSlice, varInter)

	fmt.Printf("varSlice %v\n", varSlice == nil)
	fmt.Printf("varInter %v\n", varInter == nil)

	var varSlice1 = []string{}
	var varSlice2 = []string(nil)

	fmt.Printf("varSlice1 len:%d\tvarSlice2 len:%d\n", len(varSlice1), len(varSlice2))
	fmt.Printf("varSlice1 isNil:%v\tvarSlice2 isNil:%v\n", varSlice1 == nil, varSlice2 == nil)

	fmt.Printf("varSlice1 and nil: %v\n", checkEqualSlice(varSlice1, nil))
}

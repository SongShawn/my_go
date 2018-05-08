package main

import (
	"fmt"
	"reflect"
)

func resetValue(value reflect.Value) {

	//fmt.Println(value.Kind())
	switch value.Kind() {
	case reflect.Ptr:
		fmt.Println("ptr*")
		switch value.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println("int*")
			value.SetInt(0)
		}

		//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		//	fmt.Println("uint*")
		//	value.SetUint(0)
	}

}

func main() {

	var a int = 100

	fmt.Printf("old a : %d\n", a)
	resetValue(reflect.ValueOf(&a))
	fmt.Printf("new a : %d\n", a)
}

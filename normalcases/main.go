package main

import (
	"encoding/hex"
	"fmt"
	"regexp"
)

func main() {
	//hexEncode()
	//useRegexpPkt()
	//useMap()
	userInherit()
}

// 继承

type AAA struct {
}

type AAAer interface {
	func1111()
}

func (x *AAA) func1111() {
	fmt.Println("AAA")
}

type BBB struct {
	AAA
	a int
}

func (x *BBB) func1111() {
	fmt.Println("BBB")
}

func Inherittest(in AAAer) {
	in2 := in.(*AAA)
	in2.func1111()
}

func userInherit() {
	a := &BBB{}
	Inherittest(a)
}

func useMap() {

	a := make(map[int]*int)
	var b = 10
	a[10] = &b

	c := a[10]
	if c == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("not nil")
	}

	c = a[20]
	if c == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("not nil")
	}

	fmt.Println(len(a))
}

func hexEncode() {
	name := []byte("指挥官")
	fmt.Println(hex.EncodeToString(name))
}

func useRegexpPkt() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`)

	content1 := []byte(`<a href="https://sobooks.cc/go.html?url=https://pan.baidu.com/s/1CnEep8caA7MlklKhTeUnCg" target="_blank" rel="nofollow" >百度网盘</a>`)

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	pattern2 := regexp.MustCompile(
		`<a href="https://sobooks.cc/go.html\?url=(?P<key>\w+)".*百度网盘<\/a>`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := []byte("$key=$value\n")

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		fmt.Println(string(content[submatches[0]:submatches[1]]))
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))

	template1 := []byte("$key ")
	result1 := []byte{}

	result1 = pattern2.Expand(result1, template1, content1, []int{0, len(content1)})
	fmt.Println(result1)
}

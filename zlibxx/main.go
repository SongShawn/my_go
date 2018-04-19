package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
)

var filename = "C:/Users/duoyi/go/src/mygo/zlibxx/英文版圣经.txt"

func main() {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile failed, err=", err)
		return
	}
	fmt.Printf("before compress size:%d\n", len(data))
	compedData := bytes.Buffer{}
	zlibW, _ := zlib.NewWriterLevel(&compedData, zlib.DefaultCompression)
	zlibW.Write(data)
	zlibW.Close()

	fmt.Printf("after compress size:%d\n", compedData.Len())
}

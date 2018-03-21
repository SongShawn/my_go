package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func main() {

	urlValue := url.Values{}
	urlValue.Set("e_secret_key", "201818")

	cli := &http.Client{}
	rsp, err := cli.PostForm("https://sobooks.cc/books/8132.html", urlValue)
	if err != nil {
		os.Exit(4)
	}

	defer rsp.Body.Close()
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		os.Exit(5)
	}

	// <a href="https://sobooks.cc/go.html?url=https://pan.baidu.com/s/1NbTLo9YBuLYrUpKrJUyLxw" target="_blank" rel="nofollow">百度网盘</a>

	baiduUrlRe, err := regexp.Compile(`<a href="https://sobooks.cc/go.html?.*百度网盘<\/a>`)
	if err != nil {
		os.Exit(7)
	}

	for _, v := range baiduUrlRe.FindAllSubmatchIndex(data, -1) {
		fmt.Println(v)
		fmt.Println(string(data[v[0]:v[1]]))
	}
}

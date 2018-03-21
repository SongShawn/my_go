package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "mac return")
	})

	fmt.Println(http.ListenAndServe(":12345", nil))
}

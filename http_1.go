package main

import (
	"fmt"
	"net/http"
)

var i = 1

func hello(w http.ResponseWriter, r *http.Request) {
	i++
	fmt.Println("hello world!", i)
	fmt.Fprintf(w, "HELLO WORLD!")
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	fmt.Fprintf(w, r.Method)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/test", test)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("http listen error!")
	}
}

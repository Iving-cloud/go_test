package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayhi(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./1.tmpl")
	if err != nil {
		fmt.Printf("parse the file error\n")
	}
	mp := map[string]interface{}{
		"name": "张三",
		"sex":  "男",
	}
	err = t.Execute(w, mp)
	if err != nil {
		fmt.Println("error")
	}
}
func main() {
	http.HandleFunc("/", sayhi)
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		fmt.Println("listen error")
	}
}

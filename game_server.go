package main

import (
	"fmt"
	"net/http"
)

func game_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	fmt.Fprintf(w, r.URL.String())
}

type Myhandler struct {
}

func (h *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "demo")
}
func main() {
	http.HandleFunc("/game/{index}", game_handler)
	myhandler := Myhandler{}
	http.Handle("/demo", &myhandler)
	http.ListenAndServe("0.0.0.0:8811", nil)
}

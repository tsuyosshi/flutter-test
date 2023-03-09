package main

import (
	"goapp/data"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", testLogin)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hoge!"))
}

func testLogin(w http.ResponseWriter, r *http.Request) {
	_, user_name := data.TestPost()
	w.Write([]byte(user_name))
}

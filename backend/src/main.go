package main

import (
	"encoding/json"
	"goapp/data"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", testLogin)
	http.HandleFunc("/get", getHandler)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hoge!"))
}

func testLogin(w http.ResponseWriter, r *http.Request) {
	_, user_name := data.TestPost()
	w.Write([]byte(user_name))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	id, user_name := data.TestPost()
	user := User{id, user_name}
	res, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

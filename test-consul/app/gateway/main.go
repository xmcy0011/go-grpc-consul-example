package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/user/create", func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("userName")
		pwd := request.URL.Query().Get("userPwd")
		log.Println(name, pwd)
	})
	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}

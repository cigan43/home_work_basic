package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string
	LastName string
	Age      int64
}

func main() {
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/get", handlerGet)
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		log.Fatalln("error listen server")
	}
}

func handlerPost(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		w.Write([]byte("Hellooooooo"))
	}
}

func handlerGet(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Println("Hellooo")
	}
}

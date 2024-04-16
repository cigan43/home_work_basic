package main

import (
	"encoding/json"
	"fmt"
	"io"
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
		// w.Write([]byte("Hellooooooo"))
		user := &User{}
		err := json.NewDecoder(req.Body).Decode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(user)
	}
}

func handlerGet(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		user := &User{
			Name:     "John",
			LastName: "Pup",
			Age:      30,
		}
		// b := new(bytes.Buffer)
		// err := json.NewEncoder(b).Encode(user)
		// if err != nil{
		// 	return
		// }
		io.WriteString(w, "name: "+user.Name)
		io.WriteString(w, "\n lastname: "+user.LastName)
		io.WriteString(w, "\n age:"+string(user.Age))
	}
}

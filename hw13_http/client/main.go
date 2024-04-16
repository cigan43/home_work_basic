package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server"
)

func clientPost(url string) error {
	user := &server.User{
		Name:     "John",
		LastName: "Pup",
		Age:      30,
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(user)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", b)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	return nil
}

func main() {
	resp, err := http.Get("127.0.0.1:8000/")
	if err != nil {
		fmt.Println("Ошибка Запроса", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ошибка чтения тела ответа", err)
	}
	fmt.Println(string(body))
}

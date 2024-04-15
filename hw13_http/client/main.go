package main

import (
	"fmt"
	"io"
	"net/http"
)

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

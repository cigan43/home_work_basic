package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/pflag"
)

type Config struct {
	address string
	url     string
}

func (cfg *Config) ConfigAddress(value string) {
	// fmt.Println(value)
	switch {
	case value == "":
		cfg.address = "127.0.0.1:8000"
	default:
		cfg.address = value
	}
}

func (cfg *Config) ConfigUrl(value string) {
	cfg.url = value
}

func clientPost(address, url string) error {
	user := []struct {
		Name     string
		LastName string
		Age      int64
	}{
		{
			Name:     "John",
			LastName: "Pup",
			Age:      30,
		},
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(user)
	if err != nil {
		return err
	}
	fmt.Println(address, url, b)
	resp, err := http.Post(fmt.Sprintf("%s/%s", address, url), "application/json", b)
	fmt.Println((resp))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	return nil
}

func clentGet(address, url string) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", address, url))
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

func main() {
	var (
		address  string
		url      string
		showHelp bool
	)

	pflag.StringVarP(&address, "address", "a", "",
		"address ip:port")
	pflag.StringVarP(&url, "url", "u", "",
		"url:  'post' or 'get'")
	pflag.BoolVarP(&showHelp, "help", "h", false,
		"Show help message")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		return
	}

	c := Config{}
	c.ConfigAddress(address)
	c.ConfigUrl(url)

	if c.url == "post" {
		clientPost(c.address, c.url)
	} else if c.url == "get" {
		clentGet(c.address, c.url)
	} else {
		log.Fatal("url empty")
	}
}

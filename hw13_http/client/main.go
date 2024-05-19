package main

import (
	"bytes"
	"context"
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

func (cfg *Config) ConfigURL(value string) {
	cfg.url = value
}

func clientPost(address, url string, data []byte) (string, error) {
	// jsonStr := []byte(`{"Name":"John", "LastName": "Pup", "Age":30}`)
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("http://%s/%s", address, url), bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ошибка чтения тела ответа", err)
	}
	fmt.Println(string(body))
	return resp.Status, nil
}

func clientGet(address, url string) (string, error) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("http://%s/%s", address, url), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ошибка чтения тела ответа", err)
	}
	fmt.Println(string(body), "!!!!!!!")
	return resp.Header[][], nil
	// return string(body), nil
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
	c.ConfigURL(url)

	switch c.url {
	case "post":
		jsonStr := []byte(`{"Name":"John", "LastName": "Pup", "Age":30}`)
		_, err := clientPost(c.address, c.url, jsonStr)
		if err != nil {
			fmt.Println(err)
		}
	case "get":
		_, err := clientGet(c.address, c.url)
		if err != nil {
			fmt.Println(err)
		}
	default:
		log.Fatal("url empty")
	}
}

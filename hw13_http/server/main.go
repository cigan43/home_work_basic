package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/pflag"
)

type User struct {
	Name     string
	LastName string
	Age      int64
}

type Server struct {
	Address string
	Port    int64
}

func (cfg *Server) ConfigAddress(value string) {
	// fmt.Println(value)
	switch {
	case value == "":
		cfg.Address = "0.0.0.0"
	default:
		cfg.Address = value
	}
}

func (cfg *Server) ConfigPort(value int64) {
	// fmt.Println(value)
	switch {
	case value == 0:
		cfg.Port = 8000
	default:
		cfg.Port = value
	}
}

func main() {
	var (
		address  string
		port     int64
		showHelp bool
	)

	pflag.StringVarP(&address, "address", "a", "",
		"address run server")
	pflag.Int64VarP(&port, "port", "p", 0,
		"port run server")
	pflag.BoolVarP(&showHelp, "help", "h", false,
		"Show help message")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		return
	}

	c := Server{}
	c.ConfigAddress(address)
	c.ConfigPort(port)
	mux := http.NewServeMux()
	mux.HandleFunc("/post", handlerPost)
	mux.HandleFunc("/get", handlerGet)
	srv := http.Server{
		Addr:              fmt.Sprintf("%s:%d", c.Address, c.Port),
		WriteTimeout:      1 * time.Second,
		Handler:           mux,
		ReadHeaderTimeout: 1 * time.Second,
	}
	http.HandleFunc("/get", handlerGet)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

func handlerPost(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
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
		fmt.Println(req)
		io.WriteString(w, "name: "+user.Name)
		io.WriteString(w, "\nlastname: "+user.LastName)
		io.WriteString(w, "\nage:"+fmt.Sprintf("%d", user.Age))
	}
}

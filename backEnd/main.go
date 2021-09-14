package main

import (
	"net"
	"net/http"
)

func main() {
	r, err := router.NewRouter()
	if err != nil {
		return
	}

	server := &http.Server{
		Addr:    net.JoinHostPort("127.0.0.1", "8080"),
		Handler: r,
	}
}

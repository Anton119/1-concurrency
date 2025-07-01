package main

import (
	"email/configs"
	"email/internal/verify"
	"fmt"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	fmt.Println(conf)

	router := http.NewServeMux()

	verify.NewHandler(router, conf)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Printf("server is lisening on port: %v", server.Addr)
	server.ListenAndServe()

}

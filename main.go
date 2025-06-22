package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func getRandNum(w http.ResponseWriter, r *http.Request) {
	n := rnd.Intn(6) + 1
	w.Write([]byte(fmt.Sprintf("%v", n)))
	fmt.Println("num is returned")
	return
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/rand", getRandNum)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listening")

	server.ListenAndServe()
}

// for pr
// for pr

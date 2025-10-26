package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Run() error {
	port := 7540
	strPort := os.Getenv("TODO_PORT")
	if strPort == "" {
		port = 7540
	} else {
		persPort, err := strconv.Atoi(strPort)
		if err != nil {
			port = persPort
		}
	}
	http.Handle("/", http.FileServer(http.Dir("web")))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	Run()
}

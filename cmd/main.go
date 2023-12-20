package main

import (
	"golang-project-template/internal/common"
	"log"
	"net/http"
)

func main() {
	common.LoadEnv()
	common.SetupDB()
	err := http.ListenAndServe(":5005", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	}))
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	os.Mkdir("file", 0777)

	http.Handle("/pollux/", http.StripPrefix("/pollux/", http.FileServer(http.Dir("file"))))

	http.HandleFunc("/ping/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

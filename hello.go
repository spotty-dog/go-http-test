package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world.")
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/json", JsonServ)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello http world!\n")
}

type ResponseType struct {
	ID   int
	Name string
}

func JsonServ(w http.ResponseWriter, req *http.Request) {

	data := ResponseType{
		ID:   0,
		Name: "brad",
	}

	j, err := json.Marshal(data)
	i, err := w.Write(j)
	log.Printf("bytes written: %d", i)
	if err != nil {
		log.Println("json write error: ", err)
	}
}

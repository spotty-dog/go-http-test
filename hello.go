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
		Name: "Napolean Dynamite",
	}

	pretty := req.URL.Query().Get("pretty")

	j, err := marshalJson(data, pretty)

	if err != nil {
		log.Fatal("marshal error: ", err)
	} else {
		_, err := w.Write(j)
		if err != nil {
			log.Println("json write error: ", err)
		}

	}
}

func marshalJson(data ResponseType, pretty string) ([]byte, error) {
	if pretty != "" {
		return json.MarshalIndent(data, "", "    ")
	} else {
		return json.Marshal(data)
	}
}

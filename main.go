package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8000", "http service address")

func createServer(writer http.ResponseWriter, reader *http.Request) {
	log.Println(reader.URL)
	if reader.URL.Path != "/" {
		http.Error(writer, "Invalid URL", http.StatusNotFound)
		return
	}
	if reader.Method != "GET" {
		http.Error(writer, "Invalid method", http.StatusMethodNotAllowed)
	}
	http.ServeFile(writer, reader, "index.html")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", createServer)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

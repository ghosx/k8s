package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.UserAgent())
	_, _ = fmt.Fprintf(w, "<h1>Hello Index</h1><p>ip:"+addr+"</p>")
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.UserAgent())
	_, _ = fmt.Fprintf(w, "<h1>Health check</h1><p>ip:"+addr+"</p>")
}

func init() {
	addr = os.Getenv("POD_IP")
	serverPort = os.Getenv("SERVER_PORT")
}

var addr string
var serverPort string

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health", health)
	fmt.Println("Server starting...")
	_ = http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}

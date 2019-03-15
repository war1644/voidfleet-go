package main

import (
	"fmt"
	"github.com/zserge/webview"
	"net/http"
	"os"
	"path/filepath"
)

var frame string
var rootDir string
var events chan string // keyboard events

const WWW = "/static/"
const PORT = ":1212"

var W, H = 800, 600

func init() {
	events = make(chan string, 1000)
	var err error
	rootDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	prefixChannel := make(chan string)
	go app(prefixChannel)
	prefix := <-prefixChannel
	err := webview.Open("Void Fleet",
		prefix+WWW+"index.html", W, H, true)
	if err != nil {
		fmt.Println("webview.Open error : ", err)
	}
}

func app(prefixChannel chan string) {
	mux := http.NewServeMux()
	mux.Handle(WWW, http.StripPrefix(WWW, http.FileServer(http.Dir(rootDir+WWW))))
	go GameRun()
	mux.HandleFunc("/frame", loopFrame)
	mux.HandleFunc("/key", captureKeys)
	prefixChannel <- "http://127.0.0.1" + PORT
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Println("http.ListenAndServe error : ", err)
	}
}

package main

import (
	"fmt"
	"github.com/zserge/webview"
	"net/http"
	"os"
	"path/filepath"
)

var rootDir string
var events chan string // js events

const WWW = "/static/"
const PORT = ":1212"

var W, H = 400, 600

func init() {
	events = make(chan string, 1000)
	var err error
	rootDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
}

func handleRPC(w webview.WebView, data string) {
	fmt.Println("js console: ", data)
}

func main() {
	prefixChannel := make(chan string)
	go startServer(prefixChannel)
	url := <-prefixChannel
	w := webview.New(
		webview.Settings{
			Width:                  W,
			Height:                 H,
			Title:                  "Void Fleet",
			URL:                    url + WWW + "index.html",
			ExternalInvokeCallback: handleRPC,
			Debug:                  true,
		})
	defer w.Exit()
	w.Run()
}

func startServer(prefixChannel chan string) {
	mux := http.NewServeMux()
	//static file process
	mux.Handle("/static/", http.StripPrefix(WWW, http.FileServer(http.Dir(rootDir+WWW))))
	mux.Handle("/css/", http.StripPrefix(WWW, http.FileServer(http.Dir(rootDir+WWW))))
	mux.Handle("/html/", http.StripPrefix(WWW, http.FileServer(http.Dir(rootDir+WWW))))
	mux.Handle("/js/", http.StripPrefix(WWW, http.FileServer(http.Dir(rootDir+WWW))))
	go GameRun()
	mux.HandleFunc("/frame", loopFrame)
	mux.HandleFunc("/key", captureKeys)
	prefixChannel <- "http://127.0.0.1" + PORT
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Println("http.ListenAndServe error : ", err)
	}
}

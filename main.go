package main

import (
	"fmt"
	//"github.com/zserge/webview"
	"net/http"
	"os"
	"path/filepath"
)

var rootDir string
var events chan string // js events

const STATIC = "/static/"
const PORT = ":1212"

//var W, H = 400, 600

func init() {
	events = make(chan string, 100)
	var err error
	rootDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
}

//考虑以后可移植性，这个只拿来debug
//数据通信直接用js fetch
//func handleRPC(w webview.WebView, data string) {
//	fmt.Println("js console: ", data)
//}

func main() {
	prefixChannel := make(chan string)
	go startServer(prefixChannel)
	url := <-prefixChannel
	fmt.Println(url, "main开始监听")
	//w := webview.New(
	//	webview.Settings{
	//		Width:                  W,
	//		Height:                 H,
	//		Title:                  "Void Fleet",
	//		URL:                    url + STATIC + "index.html",
	//		ExternalInvokeCallback: handleRPC,
	//		Debug:                  true,
	//	})
	//defer w.Exit()
	//w.Run()
}

func startServer(prefixChannel chan string) {
	mux := http.NewServeMux()
	//static file process
	mux.Handle("/static/", http.StripPrefix(STATIC, http.FileServer(http.Dir(rootDir+STATIC))))
	mux.Handle("/css/", http.StripPrefix(STATIC, http.FileServer(http.Dir(rootDir+STATIC))))
	mux.Handle("/html/", http.StripPrefix(STATIC, http.FileServer(http.Dir(rootDir+STATIC))))
	mux.Handle("/js/", http.StripPrefix(STATIC, http.FileServer(http.Dir(rootDir+STATIC))))
	go GameRun()
	mux.HandleFunc("/frame", loopFrame)
	mux.HandleFunc("/event", captureEvent)
	mux.HandleFunc("/key", captureKey)
	fmt.Println("http://127.0.0.1"+PORT, "开始监听")
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Println("http.ListenAndServe error : ", err)
	}
	//如果接收者没有处理，会阻塞
	prefixChannel <- "http://127.0.0.1" + PORT
}

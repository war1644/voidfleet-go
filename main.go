package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/zserge/webview"
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
)

var frame string
var rootDir string
var events chan string // keyboard events
var gameOver = false
var W, H = 800, 600

//var fps = 60    // fps
var gameDelay = 10 // game speed

const WWW = "/static/"
const PORT = ":1212"

func init() {
	events = make(chan string, 1000)
	var err error
	rootDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	load(asset)
	sprites = assetImages["sprites"]
	background = assetImages["background"]
}

func main() {
	prefixChannel := make(chan string)
	go app(prefixChannel)
	prefix := <-prefixChannel
	err := webview.Open("Void fleet",
		prefix+WWW+"index.html", W, H, true)
	if err != nil {
		fmt.Println("webview.Open error : ", err)
	}
}

// create a frame from the image
func createFrame(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	frame = base64.StdEncoding.EncodeToString(buf.Bytes())
}

// print a line of text to the image
func printLine(img *image.RGBA, x, y int, label string, col color.RGBA) {

	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: inconsolata.Bold8x16,
		Dot:  point,
	}
	d.DrawString(label)
}

func app(prefixChannel chan string) {
	mux := http.NewServeMux()
	mux.Handle(WWW, http.StripPrefix(WWW, http.FileServer(http.Dir(rootDir+WWW))))
	go generateFrames()
	mux.HandleFunc("/frame", loopFrame)
	mux.HandleFunc("/key", captureKeys)
	prefixChannel <- "http://127.0.0.1" + PORT
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Println("http.ListenAndServe error : ", err)
	}
}

func captureKeys(w http.ResponseWriter, r *http.Request) {
	ev := r.FormValue("event")
	if ev == "" {
		ev = "Space"
	}
	events <- ev
	w.Header().Set("Cache-Control", "no-cache")
}

func loopFrame(w http.ResponseWriter, r *http.Request) {
	str := "<img class='fight-screen' style='display: block; width: 100%; height: 100%' src=data:image/png;base64," + frame + ">"
	w.Header().Set("Cache-Control", "no-cache")
	len, err := w.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len)
	}
}

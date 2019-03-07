package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/zserge/webview"
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
	"html/template"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var frame string       // game frames
var rootDir string     // current directory
var events chan string // keyboard events
var gameOver = false   // end of game
var W, H = 800, 600    // width and height of the window
var frameRate int      // how many frames to show per second (fps)
var gameDelay int      // delay time added to each game loop
var assetImages map[string]image.Image
var html string

const WWW = "/static/"
const PORT = ":1212"

var asset = [][2]string{
	{"background", "asset/img/test.png"},
}

func init() {
	// events is a channel of string events that come from the front end
	events = make(chan string, 1000)
	// getting the current directory to access resources
	var err error
	rootDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	frameRate = 60 // fps
	gameDelay = 10 // 20 ms delay  maybe is cpu
	//sprites = getImage(dir + "/public/images/sprites.png") // spritesheet
	//background = getImage(dir + "/public/images/bg.png")   // background image
}

func getImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return img
}

func load(asset [][2]string) map[string]image.Image {
	if len(asset) > 0 {
		for _, v := range asset {
			assetImages[v[0]] = getImage(rootDir + v[1])
		}
	}
	return assetImages
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
	mux.HandleFunc("/start", start)
	mux.HandleFunc("/frame", loopFrame)
	mux.HandleFunc("/key", captureKeys)
	prefixChannel <- "http://127.0.0.1" + PORT
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Println("http.ListenAndServe error : ", err)
	}
}

// start the game
func start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(rootDir + "/static/html/home.html")
	go generateFrames()
	err := t.Execute(w, 1000/frameRate)
	if err != nil {
		fmt.Println("t.Execute error : ", err)
	}
}

// capture keyboard events
func captureKeys(w http.ResponseWriter, r *http.Request) {

}

// get the game frames
func loopFrame(w http.ResponseWriter, r *http.Request) {
	str := "data:image/png;base64," + frame
	w.Header().Set("Cache-Control", "no-cache")
	w.Write([]byte(str))
}

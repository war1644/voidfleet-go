package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"time"
	"void_fleet/game"
)

var asset = [][2]string{
	{"background", "asset/img/bg.png"},
	{"sprites", "asset/img/sprites.png"},
}

type ResponseJson struct {
	PlayerStatus *game.Player `json:"playerStatus"`
}

var frameData *ResponseJson

//var fps = 60

func GameRun() {
	g := game.NewGame()
	generateFrames(g)
}

func generateFrames(g *game.Game) {
	loop := 0
	for !g.Stop {
		time.Sleep(time.Millisecond * time.Duration(g.Delay))
		select {
		case ev := <-events:
			if ev == "Space" { // space bar
				fmt.Println("game loop", ev)
				//if beam.Status == false {
				//	beamShot = true
				//}
			}
		default:
		}
		NewData(g)
		loop++
		//dst := image.NewRGBA(image.Rect(0, 0, W, H))
		//gift.New().Draw(dst, assetImages["background"])
		//createFrame(dst)
		//if gameOver {
		//	//playSound("explosion")
		//	time.Sleep(time.Second)
		//}
	}
}

func NewData(g *game.Game) {
	//player
	frameData = &ResponseJson{PlayerStatus: g.Player}
}

func createFrame(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	base64.StdEncoding.EncodeToString(buf.Bytes())
}

func loopFrame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	err := json.NewEncoder(w).Encode(frameData)
	if err != nil {
		fmt.Println(err)
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

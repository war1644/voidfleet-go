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

type PlayerJson struct {
}

type PlanetJson struct {
}

type StarJson struct {
}

type GalaxyJson struct {
}

type ResponseJson struct {
	Player *PlayerJson `json:"player"`
	Planet *PlanetJson `json:"planet"`
	Galaxy *StarJson   `json:"galaxy"`
	Star   *GalaxyJson `json:"star"`
}

var frameData []byte

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
	var err error
	frameData, err = json.Marshal(struct {
		Player *game.Player
		Planet *game.Planet
		Galaxy *game.Galaxy
	}{
		Player: g.Player,
		Planet: g.CurrentPlanet,
		Galaxy: g.Galaxy,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func InitData(g *game.Game) {
	var err error
	frameData, err = json.Marshal(struct {
		Player *game.Player
		Planet *game.Planet
		Galaxy *game.Galaxy
	}{
		Player: g.Player,
		Planet: g.CurrentPlanet,
		Galaxy: g.Galaxy,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func createFrame(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	base64.StdEncoding.EncodeToString(buf.Bytes())
}

func loopFrame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	len, err := w.Write(frameData)
	//err := json.NewEncoder(w).Encode(frameData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len)
}

func eventProcess(enevt string) {

}

func captureEvent(w http.ResponseWriter, r *http.Request) {
	event := r.FormValue("event")
	switch event {
	case "key":
		break
	case "event":
		eventValue := r.FormValue("value")
		eventProcess(eventValue)
		break
	default:
		break
	}
	w.Header().Set("Cache-Control", "no-cache")
}

func captureKey(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if key == "" {
		key = "Space"
	}
	events <- key
	w.Header().Set("Cache-Control", "no-cache")
}

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/disintegration/gift"
	"image"
	"image/png"
	"math/rand"
	"net/http"
	"time"
	"void_fleet/ecs"
	"void_fleet/game/systems"
)

var gameOver = false
var asset = [][2]string{
	{"background", "asset/img/bg.png"},
	{"sprites", "asset/img/sprites.png"},
}

//var fps = 60
var gameDelay = 10

func NewGame() {
	rand.Seed(time.Now().UTC().UnixNano())
	world := ecs.NewWorld()
	NewSystem(world)
	world.Start()
	for !world.Stop {
		time.Sleep(time.Millisecond * time.Duration(gameDelay))
		world.Update()
	}
	world.Remove()
}

func NewSystem(world *ecs.World) {
	world.AddSystem(
		systems.NewLoad(asset, rootDir),
	)

}

func generateFrames() {
	rand.Seed(time.Now().UTC().UnixNano())
	loop := 0 // game loop
	// main game loop
	for !gameOver {
		// game speed
		time.Sleep(time.Millisecond * time.Duration(gameDelay))
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
		dst := image.NewRGBA(image.Rect(0, 0, W, H))
		gift.New().Draw(dst, assetImages["background"])
		createFrame(dst)
		//if gameOver {
		//	//playSound("explosion")
		//	time.Sleep(time.Second)
		//}
		loop++
	}
}

// create a frame from the image
func createFrame(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	frame = base64.StdEncoding.EncodeToString(buf.Bytes())
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

func captureKeys(w http.ResponseWriter, r *http.Request) {
	ev := r.FormValue("event")
	if ev == "" {
		ev = "Space"
	}
	events <- ev
	w.Header().Set("Cache-Control", "no-cache")
}

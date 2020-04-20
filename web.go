package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	//jsoniter "github.com/json-iterator/go"
	//"github.com/json-iterator/go"
	"image"
	"image/png"
	"net/http"
	"time"
	"void_fleet/game"
)

//var asset = [][2]string{
//	{"background", "asset/img/bg.png"},
//	{"sprites", "asset/img/sprites.png"},
//}

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

//var startData []byte
//var msgData []byte
var frameData []byte
var g *game.Game

//var fps = 60

func GameRun() {
	g = game.NewGame()
	generateFrames()
}

func generateFrames() {
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
		//UpdateData()
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

func StartData() []byte {
	//var err error
	data, err := json.Marshal(struct {
		Player *game.Player
		Planet *game.Planet
		Galaxy *game.Galaxy
		Msg    []game.Msg
	}{
		Player: g.Player,
		Planet: g.CurrentPlanet,
		Galaxy: g.Galaxy,
		Msg:    g.Event.Get(6),
	})

	if err != nil {
		fmt.Println(err)
	}
	return data
}

func UpdateData() {
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

func MsgData() []byte {
	data, err := json.Marshal(struct {
		Msg []game.Msg
	}{
		Msg: g.Event.Get(3),
	})

	if err != nil {
		fmt.Println(err)
	}
	return data
}

func createFrame(img image.Image) {
	var buf bytes.Buffer
	var err = png.Encode(&buf, img)
	if err != nil {
		fmt.Println(err)
	}

	base64.StdEncoding.EncodeToString(buf.Bytes())
}

//func startFrame(w http.ResponseWriter, r *http.Request) {
//
//	len, err := w.Write(startData)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(len)
//}

func loopFrame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	i, err := w.Write(frameData)
	//err := json.NewEncoder(w).Encode(frameData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

func eventProcess(event string) []byte {
	switch event {
	case "msg":
		//获取msg
		return MsgData()
	case "start":
		//刷新数据
		return StartData()
	//case "start_data":
	//	//初始数据
	//	return startData
	default:
	}
	return []byte{}
}

func captureEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")
	event := r.FormValue("event")
	w.Header().Set("Cache-Control", "no-cache")
	data := eventProcess(event)
	i, err := w.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

func captureKey(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if key == "" {
		key = "Space"
	}
	w.Header().Set("Cache-Control", "no-cache")
	events <- key
}

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
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

		NewDomData(g)

		//dst := image.NewRGBA(image.Rect(0, 0, W, H))
		//gift.New().Draw(dst, assetImages["background"])
		//createFrame(dst)
		//if gameOver {
		//	//playSound("explosion")
		//	time.Sleep(time.Second)
		//}
		loop++
	}
}

func NewDomData(g *game.Game) {
	template.Must(template.New("issuelist").Parse(`
<span class="xx-small badge badge-info move-dom id-gate" id="">跳跃门</span>
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

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

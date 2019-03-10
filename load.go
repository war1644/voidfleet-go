package main

import (
	"fmt"
	"image"
	"os"
)

var assetImages = make(map[string]image.Image)
var asset = [][2]string{
	{"background", "asset/img/bg.png"},
	{"sprites", "asset/img/sprites.png"},
}

func load(assets [][2]string) {
	if len(assets) > 0 {
		for _, v := range assets {
			assetImages[v[0]] = getImage(rootDir + "/" + v[1])
		}
	}
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

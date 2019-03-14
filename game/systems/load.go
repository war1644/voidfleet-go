package systems

import (
	"fmt"
	"image"
	"os"
	"void_fleet/ecs"
)

type Load struct {
	RootDir     string
	assets      [][2]string
	AssetImages map[string]image.Image
}

func NewLoad(assets [][2]string, dir string) ecs.System {
	return &Load{
		RootDir:     dir,
		assets:      assets,
		AssetImages: make(map[string]image.Image),
	}
}

func (s *Load) Start(world *ecs.World) {
	if len(s.assets) > 0 {
		for _, v := range s.assets {
			s.AssetImages[v[0]] = s.getImage(s.RootDir + "/" + v[1])
		}
	}
}

func (s *Load) Update(w *ecs.World) {

}

func (s *Load) Remove() {

}

func (s *Load) getImage(filePath string) image.Image {
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

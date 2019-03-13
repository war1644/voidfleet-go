package systems

//import (
//	"golang.org/x/image/font"
//	"golang.org/x/image/font/inconsolata"
//	"golang.org/x/image/math/fixed"
//	"image"
//	"image/color"
//)
//
//// print a line of text to the image
//func printLine(img *image.RGBA, x, y int, label string, col color.RGBA) {
//
//	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}
//	d := &font.Drawer{
//		Dst:  img,
//		Src:  image.NewUniform(col),
//		Face: inconsolata.Bold8x16,
//		Dot:  point,
//	}
//	d.DrawString(label)
//}

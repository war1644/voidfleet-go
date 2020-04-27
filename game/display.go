package game

import "fmt"

type Display struct{}

func NewDisplay() *Display {
	return &Display{}
}

func (s *Display) Show(msg string) {
	fmt.Println(msg)
}

func (s *Display) AutoShow(msg string) {

	fmt.Println(msg)
}

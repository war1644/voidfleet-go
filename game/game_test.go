package game

import (
	"fmt"
	"testing"
)

func TestGame_Item_Build(t *testing.T) {
	game := NewGame()
	fmt.Println(game.Galaxy.List["PLA"][1].Fleet)
	fmt.Println(game.Ships)
}

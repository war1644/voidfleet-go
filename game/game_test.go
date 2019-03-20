package game

import (
	"fmt"
	"testing"
)

func TestGame_Item_Build(t *testing.T) {
	game := NewGame()
	for _, v := range game.Galaxy.Current {
		fmt.Println(v)
	}
	fmt.Println(game.Items.Goods)
	//for i:=0; i<10;i++  {
	//	fmt.Println(RandNum(0,80,i))
	//}
}

func BenchmarkJSON(b *testing.B) {

	for i := 0; i < b.N; i++ {

	}
}

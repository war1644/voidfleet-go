package game

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestGame_Item_Build(t *testing.T) {
	game := NewGame()
	for _, v := range game.Galaxy.Current {
		fmt.Println(v)
	}
	fmt.Println(game.Items.Goods)
	for i := 0; i < 10; i++ {
		fmt.Println(RandNum(0, 80, i))
	}
}

func TestEvent_Get(t *testing.T) {
	e := NewEvent()
	for i := 0; i < 6; i++ {
		e.NewMsg("info", strconv.Itoa(i))
	}
	fmt.Println(e.Get(2))
	fmt.Println(e.MsgList)
}

func TestGame_Reflect(t *testing.T) {
	game := NewGame()
	msgType := "info"
	mtm := reflect.ValueOf(game.Event.MsgTypeMap)
	fmt.Println(mtm.FieldByName(msgType))
	//fmt.Println(game.Event.MsgTypeMap.info)
}

func BenchmarkJSON(b *testing.B) {

	for i := 0; i < b.N; i++ {

	}
}

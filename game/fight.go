package game

import (
	"time"
)

type Fight struct {
	AiHp     int
	PlayerHp int
	game     *Game
}

func Start(item *Items, player *Player, aiShip *Ship) bool {
	display := NewDisplay()
	display.Show("遭遇舰队")
	playerHp := player.Ship.HP
	aiHp := aiShip.HP

	isFirst := RandNum(1, 2, 0)
	haveGoods := RandNum(0, 9, 1)
	weapon := [...]string{"导弹", "EMP", "鱼雷", "激光炮", "轨道炮"}
	//weaponHit := [...]int{150,200,300,500,1000}
	//aiWeaponIndex := 0
	for aiHp > 0 && playerHp > 0 {
		weaponIndex := RandNum(0, len(weapon), playerHp)
		aiWeaponIndex := RandNum(0, len(weapon), aiHp)
		if isFirst == 1 {
			display.AutoShow(weapon[weaponIndex] + " 攻击！")
			display.AutoShow("敌方受到" + string(player.Ship.Attack))
			aiHp -= player.Ship.Attack
			isFirst = 0
		} else {
			display.AutoShow(weapon[aiWeaponIndex] + " 攻击！")
			display.AutoShow("我方受到{aiHp}" + string(aiShip.Attack))
			playerHp -= aiShip.Attack
			isFirst = 1
		}
		time.Sleep(time.Millisecond * time.Duration(1000))
	}
	if playerHp > 0 {
		display.AutoShow("我方大破敌舰")
	} else {
		display.AutoShow("我方被炸成了太空垃圾")
		return false
	}

	player.AddKill(1)
	//装备掉落
	if haveGoods >= 8 {
		index := RandNum(0, len(aiShip.Equips), 1)
		goods := item.Goods[aiShip.Equips[index]]
		display.AutoShow("获得" + goods.Name)
		goods.Quantity = 1
		player.AddCargoGoods(goods)
	}
	return true
}

func Fleet() {

}

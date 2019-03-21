package game

type Fight struct {
	AiHp     int
	PlayerHp int
	game     *Game
}

func Battle(item *Items, player *Player, AiShip *Ship) bool {
	display := NewDisplay()
	display.Show("遭遇舰队")
	playerHp := player.Ship.HP
	aiHp := AiShip.HP

	isFirst := RandNum(1, 2, 0)
	haveGoods := RandNum(0, 9, 1)
	//weapon := [...]string{"导弹","EMP","鱼雷","激光炮","轨道炮"}
	//weaponHit := [...]int{150,200,300,500,1000}
	//aiWeaponIndex := 0
	for aiHp > 0 && playerHp > 0 {
		//weaponIndex = rand.Next(0,weapon.Length)
		//aiWeaponIndex = rand.Next(0,weapon.Length)
		if isFirst == 1 {
			display.AutoShow("我方沉着应战")
			display.AutoShow("{weapon[weaponIndex]} 齐射！")
			display.AutoShow("敌方躲闪不急")
			aiHp -= player.Ship.Attack
			isFirst = 0
		} else {
			display.AutoShow("敌人气势汹汹")
			display.AutoShow("weapon[weaponIndex]密集射来！")
			display.AutoShow("我方机体受损，装甲减少{aiHp}")
			playerHp -= AiShip.Attack
			isFirst = 1
		}
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
		index := RandNum(0, len(AiShip.Equips), 1)
		goods := item.Goods[AiShip.Equips[index]]
		display.AutoShow("获得{tmpName}")
		goods.Quantity = 1
		player.AddCargoGoods(goods)
	}
	return true
}

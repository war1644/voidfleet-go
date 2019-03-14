package tmp

type Event struct {
	eventType   int
	title       string
	desc        string
	effectId    int
	effectWord  string
	effectValue int
}

type GameData struct {
	goods []Goods
	ships []Ship
}

var Game map[string]interface{}

//战斗 --》海盗来袭，支援a
//获得货物，金钱
var randEvents = []Event{
	{
		eventType:   1,
		title:       "海盗",
		desc:        "海盗来袭",
		effectId:    1,
		effectWord:  "僚机损失",
		effectValue: 200,
	},
	{
		eventType:   2,
		title:       "货箱",
		desc:        "遇到太空漂浮货箱",
		effectId:    11,
		effectWord:  "获得物品",
		effectValue: 200,
	},
	{
		eventType:   2,
		title:       "货箱",
		desc:        "遇到太空漂浮货箱",
		effectId:    12,
		effectWord:  "获得装备",
		effectValue: 200,
	},
	{
		eventType:   2,
		title:       "货箱",
		desc:        "遇到太空漂浮货箱",
		effectId:    13,
		effectWord:  "获得金币",
		effectValue: 200,
	},
}

type GameObject interface {
}

type Goods struct {
}
type Ship struct {
}

func createGoods() {
	//	[
	//	new Goods("营养液", 200),
	//	new Goods("金", 12000),
	//	new Goods("钻石", 1400),
	//	new Goods("矿石", 30),
	//	new Goods("武器", 60),
	//	new Goods("木材", 10),
	//	new Goods("铜", 50),
	//	new Goods("暗物质", 450000),
	//	new Goods("生活包", 500),
	//];
}

func init() {
	//gameDtat := GameData()
}

func createShips() {
	//	Game.ships = [
	//		new Ship("帝国采矿船",6000, 150, 3, 100000, 25),
	//	new
	//	Ship("帝国货船", 5000, 50, 3, 25000, 20),
	//		new
	//	Ship("帝国战斗机", 10000, 300, 4, 75000, 20),
	//		new
	//	Ship("帝国护卫舰", 300000, 750, 7, 125000, 30),
	//		new
	//	Ship("帝国驱逐舰", 1200000, 50, 9, 400000, 40),
	//		new
	//	Ship("帝国航母", 50100100, 1000, 7, 500000, 50),
	//		new
	//	Ship("帝国歼星舰", 150100100, 2000, 5, 1000000, 100),
	//];
}

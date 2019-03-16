package game

import (
	"fmt"
)

type Player struct {
	Name       string
	Money      int
	Kill       int
	Reputation int
	Cargo      map[string]Goods
	Fleet      []Ship
	ShipsCount int
	Year       int
	Day        int
	Planet     *Planet
	GoodsCount int
	Ship       *Ship
}

func NewPlayer(money int, newShip Ship, planet *Planet, game *Game) *Player {
	player := &Player{
		Cargo:      make(map[string]Goods, 128),
		Fleet:      []Ship{},
		Name:       "player",
		Money:      money,
		Kill:       0,
		Reputation: 0,
		Year:       0,
		Day:        0,
		GoodsCount: 0,
		ShipsCount: 0,
		Ship:       &newShip,
	}
	player.SetPlant(planet, game)
	player.AddShip(newShip, true)
	return player
}

func (s *Player) CalculateYears() {
	if s.Day > 365 {
		s.Day -= 365
		s.Year += 1
	}
}

func (s *Player) SetGoToPlant(toPlanet *Planet, game *Game) bool {
	distance := s.CalculatePlanetsDistance(toPlanet)
	if s.Ship.CalculateFuel(distance) {
		if (distance / s.Ship.Speed) <= 0 {
			s.Day += 1
		} else {
			s.Day += distance / s.Ship.Speed
		}
		s.CalculateYears()
		s.SetPlant(toPlanet, game)
		return true
	}
	return false
}

func (s *Player) SetPlant(toPlanet *Planet, game *Game) {
	s.Planet = toPlanet
	if s.Planet.Name == "星系跳跃门" {
		game.IsJump = true
	} else {
		game.IsJump = false
	}
}

func (s *Player) CalculatePlanetsDistance(toPlanet *Planet) int {
	return abs(toPlanet.Distance - s.Planet.Distance)
}

func (s *Player) AddCredits(money int) {
	s.Money += money
}

func (s *Player) BuyShip(newShip Ship, isPlayer bool) {
	if s.Money >= newShip.Price {
		//买小船,抛弃货物
		if s.GoodsCount < newShip.Cargo {
			for k, v := range s.Cargo {
				s.GoodsCount -= v.Quantity
				delete(s.Cargo, k)
				if s.GoodsCount <= newShip.Cargo {
					break
				}
			}
		}
		s.AddCredits(-newShip.Price)
		s.AddShip(newShip, isPlayer)
	}
}

func (s *Player) AddShip(newShip Ship, isPlayer bool) {
	if len(s.Fleet) > 0 {
		isAdd := false
		for k := range s.Fleet {
			if s.Fleet[k].Name == newShip.Name {
				s.Fleet[k].Count += 1
				isAdd = true
			}
		}
		if !isAdd {
			s.Fleet = append(s.Fleet, newShip)
			if isPlayer {
				s.Ship = &s.Fleet[len(s.Fleet)]
			}
		}
	} else {
		s.Fleet = make([]Ship, 1)
		s.Fleet[0] = newShip
		if isPlayer {
			s.Ship = &s.Fleet[0]
		}
	}
}

func (s *Player) SellShip(newShip Ship) {
	if s.Money >= newShip.Price {
		//买小船,抛弃货物
		if s.GoodsCount < newShip.Cargo {
			for k, v := range s.Cargo {
				s.GoodsCount -= v.Quantity
				delete(s.Cargo, k)
				if s.GoodsCount <= newShip.Cargo {
					break
				}
			}
		}
		s.AddCredits(newShip.Price)
		//s.Ship = newShip
	}
}

func (s *Player) SellGood(name string, number, price int) {
	goods, ok := s.Cargo[name]
	if !ok {
		return
	}
	if number == goods.Quantity {
		delete(s.Cargo, name)
	}
	totalPrice := price * number
	s.GoodsCount -= number
	s.AddCredits(totalPrice)
}

func (s *Player) BuyGood(goods Goods) {
	totalPrice := goods.Price * goods.Quantity
	surplus := s.Ship.Cargo - s.GoodsCount
	if s.Money < totalPrice {
		return
	}
	if goods.Quantity <= surplus {
		s.GoodsCount += goods.Quantity
		s.AddCargoGoods(goods)
		s.AddCredits(-totalPrice)
	}
}

func (s *Player) AddCargoGoods(goods Goods) {
	// v是copy
	v, ok := s.Cargo[goods.Name]
	if ok {
		v.Quantity += goods.Quantity
		s.Cargo[goods.Name] = v
	} else {
		s.Cargo[goods.Name] = goods
	}
}

func (s *Player) GetState() {
	fmt.Println("--状态--",
		"Money：{credits}，旅行时间： {year} 年 {day} 天",
		"位置：{planet.galaxy} -- {planet.name}，旗舰：{ship.name}",
		"燃料： {ship.fuel}，飞行状态：{(!Game.docked).ToString()}",
		"装甲：{ship.hp}，杀敌：{kill}",
		"空余货仓：{(ship.cargo-goodsCount)}")
}

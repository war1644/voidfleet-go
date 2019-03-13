package components

import "fmt"

type Ship struct {
	name     string
	price    int
	describe string
	hp       int
	ep       int
	maxHp    int
	maxEp    int
	cargo    int
	speed    int
	fuel     int
	maxFuel  int
}

func (s *Ship) Name() string {
	return "ship"
}

func (s *Ship) CalculateFuel(fuelValue int) bool {
	tmpFuel := s.fuel + fuelValue
	if tmpFuel < 0 {
		return false
	}
	s.fuel = tmpFuel
	return true
}

func (s *Ship) CalculateHp(hpValue int) bool {
	tmpHp := s.hp + hpValue
	if tmpHp < 0 {
		return false
	}
	s.hp = tmpHp
	return true
}

func (s *Ship) Refuel(player *Player) {
	refuelPrice := s.maxFuel - s.fuel
	if player.credits-refuelPrice < 0 {
		fmt.Println("没钱加燃料")
	} else {
		fmt.Println("燃料已加满，花费{}", refuelPrice)
		player.AddCredits(-refuelPrice)
		s.fuel = s.maxFuel
	}
}

func (s *Ship) Repair(player *Player) {
	repairPrice := s.maxHp - s.hp
	if (player.credits - repairPrice) < 0 {
		fmt.Println("没钱修理")
	} else {
		fmt.Println("修理完成，花费{}", repairPrice)
		player.AddCredits(-repairPrice)
		s.hp = s.maxHp
	}
}

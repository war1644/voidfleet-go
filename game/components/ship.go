package components

import "fmt"

type Ship struct {
	ShipName      string
	Price         int
	StandardPrice int
	Describe      string
	HP            int
	EP            int
	MaxHp         int
	MaxEp         int
	Cargo         int
	Speed         int
	Fuel          int
	MaxFuel       int
}

func (s *Ship) Name() string {
	return "ship"
}

func (s *Ship) CalculateFuel(fuelValue int) bool {
	tmpFuel := s.Fuel + fuelValue
	if tmpFuel < 0 {
		return false
	}
	s.Fuel = tmpFuel
	return true
}

func (s *Ship) CalculateHp(hpValue int) bool {
	tmpHp := s.HP + hpValue
	if tmpHp < 0 {
		return false
	}
	s.HP = tmpHp
	return true
}

func (s *Ship) Refuel(player *Player) {
	refuelPrice := s.MaxFuel - s.Fuel
	if player.Credits-refuelPrice < 0 {
		fmt.Println("没钱加燃料")
	} else {
		fmt.Println("燃料已加满，花费{}", refuelPrice)
		player.AddCredits(-refuelPrice)
		s.Fuel = s.MaxFuel
	}
}

func (s *Ship) Repair(player *Player) {
	repairPrice := s.MaxHp - s.HP
	if (player.Credits - repairPrice) < 0 {
		fmt.Println("没钱修理")
	} else {
		fmt.Println("修理完成，花费{}", repairPrice)
		player.AddCredits(-repairPrice)
		s.HP = s.MaxHp
	}
}

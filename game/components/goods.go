package components

type Goods struct {
	name          string
	price         int
	standardPrice int
	IsEquip       bool
	quantity      int
	describe      string
}

func (s *Goods) Name() string {
	return "goods"
}

func (s *Goods) Sell() {
}

func (s *Goods) Buy() {
}

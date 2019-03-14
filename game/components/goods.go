package components

type Goods struct {
	GoodsName     string
	Price         int
	StandardPrice int
	IsEquip       bool
	quantity      int
	Describe      string
}

func (s *Goods) Name() string {
	return "goods"
}

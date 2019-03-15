package game

type Goods struct {
	Name          string
	Price         int
	StandardPrice int
	IsEquip       bool
	Quantity      int
	Describe      string
}

func NewGoods(goodName, goodDescribe string, goodPrice int, isEquip bool) Goods {
	return Goods{
		Name:          goodName,
		Price:         goodPrice,
		StandardPrice: goodPrice,
		Describe:      goodDescribe,
		Quantity:      0,
		IsEquip:       isEquip,
	}
}

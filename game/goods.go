package game

type Goods struct {
	Name          string
	Price         int
	StandardPrice int
	IsEquip       bool
	Quantity      int
	Describe      string
}

func NewGoods(goodsName, goodsDescribe string, goodsPrice int, isEquip bool) Goods {
	return Goods{
		Name:          goodsName,
		Price:         goodsPrice,
		StandardPrice: goodsPrice,
		Describe:      goodsDescribe,
		Quantity:      0,
		IsEquip:       isEquip,
	}
}

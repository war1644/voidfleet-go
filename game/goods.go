package game

type Goods struct {
	Name          string
	Price         int
	StandardPrice int
	IsEquip       bool
	Quantity      int
	EffectValue   int
	EffectId      int
	Describe      string
}

func NewGoods(goodsName, goodsDescribe string, goodsPrice int, isEquip bool) Goods {
	return Goods{
		Name:          goodsName,
		Price:         goodsPrice,
		StandardPrice: goodsPrice,
		Describe:      goodsDescribe,
		Quantity:      1,
		IsEquip:       isEquip,
	}
}

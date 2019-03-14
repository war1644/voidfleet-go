package game

type Goods struct {
	Name          string
	Price         int
	StandardPrice int
	IsEquip       bool
	Quantity      int
	Describe      string
}

func NewGoods(goodName, goodDescribe string, goodPrice int) Goods {
	return Goods{
		Name:     goodName,
		Price:    goodPrice,
		Describe: goodDescribe,
	}
}

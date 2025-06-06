package domain

type (
	ItemParams struct {
		ProductId string
		Count     int
		Price     float64
	}

	Item struct {
		productId string
		count     int
		price     float64
	}
)

func CreateItem(input ItemParams) Item {
	item := Item{
		productId: input.ProductId,
		count:     input.Count,
		price:     input.Price,
	}

	return item
}

func (i *Item) Map() *ItemParams {
	ip := ItemParams{}
	ip.ProductId = i.productId
	ip.Count = i.count
	ip.Price = i.price

	return &ip
}

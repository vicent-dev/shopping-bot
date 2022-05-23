package buy

import (
	"fmt"
	"shopping-bot/pkg/storage"
)

var (
	carts map[int64]*Cart
)

type Cart struct {
	Channel     int64    `json:"channel"`
	Products    []string `json:"products"`
	SuperMarket string   `json:"super_market"`
}

func GetCart(c int64) *Cart {
	if carts == nil {
		carts = make(map[int64]*Cart)
	}

	storage.Load[map[int64]*Cart](carts)

	if cart, ok := carts[c]; ok {
		return cart
	}

	carts[c] = &Cart{
		Channel: c,
	}

	return carts[c]
}

func (c *Cart) AddProduct(p string) {
	defer c.storeCarts()

	for _, cp := range c.Products {
		if cp == p {
			return
		}
	}

	c.Products = append(c.Products, p)
}

func (c *Cart) RemoveProduct(p string) {
	defer c.storeCarts()

	for i, cp := range c.Products {
		if cp == p {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			return
		}
	}
}
func (c *Cart) Reset() {
	defer c.storeCarts()

	c.Products = []string{}
}

func (c *Cart) storeCarts() {
	carts[c.Channel] = c
	fmt.Println(carts)
	storage.Store[map[int64]*Cart](carts)
}

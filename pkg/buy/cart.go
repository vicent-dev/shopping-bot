package buy

var (
	carts map[int64]*Cart
)

type Cart struct {
	Channel     int64
	Products    []string
	SuperMarket string
}

func GetCart(c int64) *Cart {
	if carts == nil {
		carts = make(map[int64]*Cart)
	}

	if cart, ok := carts[c]; ok {
		return cart
	}

	carts[c] = &Cart{
		Channel: c,
	}

	return carts[c]
}

func (c *Cart) AddProduct(p string) {
	for _, cp := range c.Products {
		if cp == p {
			return
		}
	}

	c.Products = append(c.Products, p)
}

func (c *Cart) RemoveProduct(p string) {
	for i, cp := range c.Products {
		if cp == p {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			return
		}
	}
}
func (c *Cart) Reset() {
	c.Products = []string{}
}

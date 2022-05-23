package buy

type Cart struct {
	Channel     int64    `json:"channel"`
	Products    []string `json:"products"`
	SuperMarket string   `json:"super_market"`
}

func (c *Cart) AddProduct(p string) {
	defer GetRepository().Store(c)

	for _, cp := range c.Products {
		if cp == p {
			return
		}
	}

	c.Products = append(c.Products, p)
}

func (c *Cart) RemoveProduct(p string) {
	defer GetRepository().Store(c)

	for i, cp := range c.Products {
		if cp == p {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			return
		}
	}
}
func (c *Cart) Reset() {
	defer GetRepository().Store(c)
	c.Products = []string{}
}

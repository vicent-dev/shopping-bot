package telegram

import (
	"shopping-bot/pkg/buy"
	"strings"
)

type messageProcessor func(int64, string) string

var (
	processors = map[string]messageProcessor{
		AddProductCommand:    addProduct,
		RemoveProductCommand: removeProduct,
		ListProductsCommand:  listProducts,
		ResetCartCommand:     resetCart,
	}
)

func addProduct(chatId int64, message string) string {
	c := buy.GetCart(chatId)

	c.AddProduct(message)

	return ""
}

func removeProduct(chatId int64, message string) string {
	c := buy.GetCart(chatId)
	splited := strings.Split(message, " ")
	ps := splited[1:]

	for _, p := range ps {
		c.RemoveProduct(p)
	}

	return "Removed products: " + strings.Join(ps, ",")
}

func listProducts(chatId int64, _ string) string {
	c := buy.GetCart(chatId)
	return strings.Join(c.Products, "\n")
}

func resetCart(chatId int64, _ string) string {
	c := buy.GetCart(chatId)
	c.Reset()

	return "Cart reset"
}

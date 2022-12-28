package repositories

import "github.com/IGabor98/api-cart/models"

type CartRepository interface {
	AddItemToCart(cartToken string, Item *models.Item) (*models.Cart, error)
	GetCart(cartToken string) (*models.Cart, error)
}

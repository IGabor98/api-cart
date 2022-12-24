package requests

import "github.com/IGabor98/api-cart/models"

type AddItemToCartRequest struct {
	CartToken string      `json:"cart_token"`
	Item      models.Item `json:"item"`
}

package controllers

import (
	"encoding/json"
	"github.com/IGabor98/api-cart/repositories"
	"github.com/IGabor98/api-cart/requests"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type CartController struct {
	CartRepository repositories.CartRepository
}

func (c *CartController) AddItem(w http.ResponseWriter, r *http.Request) {
	request := &requests.AddItemToCartRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart, err := c.CartRepository.AddItemToCart(request.CartToken, request.Item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *CartController) GetCart(w http.ResponseWriter, r *http.Request) {
	cartToken := chi.URLParam(r, "cartToken")
	cart, err := c.CartRepository.GetCart(cartToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

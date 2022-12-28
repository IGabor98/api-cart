package models

import "github.com/IGabor98/api-cart/models/internal"

type Item struct {
	CartToken              string                   `json:"cart_token" dynamodbav:"cart_token"`
	SK                     string                   `json:"sk" dynamodbav:"sk"`
	SearchProductResultID  uint64                   `json:"search_product_result_id" dynamodbav:"search_product_result_id"`
	SearchCriteria         internal.SearchCriteria  `json:"search_criteria" dynamodbav:"search_criteria"`
	ProductCode            string                   `json:"product_code" dynamodbav:"product_code"`
	InventoryProvider      string                   `json:"inventory_provider" dynamodbav:"inventory_provider"`
	InventoryID            string                   `json:"inventory_id" dynamodbav:"inventory_id"`
	InventoryItem          internal.InventoryItem   `json:"inventory_item" dynamodbav:"inventory_item"`
	InventoryOption        internal.InventoryOption `json:"inventory_option" dynamodbav:"inventory_option"`
	IsAvailable            bool                     `json:"is_available" dynamodbav:"is_available"`
	IsProtected            bool                     `json:"is_protected" dynamodbav:"is_protected"`
	Taxes                  []internal.Tax           `json:"taxes" dynamodbav:"taxes"`
	Fees                   []internal.Fee           `json:"fees" dynamodbav:"fees"`
	Totals                 Totals                   `json:"totals" dynamodbav:"totals"`
	CancellationProtection CancellationProtection   `json:"cancellation_protection" dynamodbav:"cancellation_protection"`
	Discount               Discount                 `json:"discount" dynamodbav:"discount"`
	UpdatedAt              string                   `json:"updated_at" dynamodbav:"updated_at"`
	CreatedAt              string                   `json:"created_at" dynamodbav:"created_at"`
}

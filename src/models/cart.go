package models

type Cart struct {
	Token         string  `json:"token" dynamodbav:"cart_token"`
	SK            string  `json:"sk" dynamodbav:"sk"`
	ChannelID     uint64  `json:"channel_id" dynamodbav:"channel_id"`
	Status        string  `json:"status" dynamodbav:"status"`
	Items         []*Item `json:"items" dynamodbav:"-"`
	RevalidatedAt string  `json:"revalidated_at" dynamodbav:"revalidated_at"`
	UpdatedAt     string  `json:"updated_at" dynamodbav:"updated_at"`
	CreatedAt     string  `json:"created_at" dynamodbav:"created_at"`
}

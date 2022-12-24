package models

type Cart struct {
	Token         string `dynamodbav:"token"`
	ChannelID     uint64 `dynamodbav:"channel_id"`
	Status        string `dynamodbav:"status"`
	Items         []Item `dynamodbav:"items"`
	RevalidatedAt string `dynamodbav:"revalidated_at"`
	UpdatedAt     string `dynamodbav:"updated_at"`
	CreatedAt     string `dynamodbav:"created_at"`
}

package repositories

import (
	"context"
	"github.com/IGabor98/api-cart/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"time"
)

type DynamoDBCartRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewCartRepository(client *dynamodb.Client) *DynamoDBCartRepository {
	return &DynamoDBCartRepository{client: client, tableName: "carts"}
}

func (r *DynamoDBCartRepository) AddItemToCart(cartToken string, item *models.Item) (*models.Cart, error) {

	if cartToken == "" {

		return r.createCart(context.TODO(), []*models.Item{item})
	} else {
		cart, err := r.GetCart(cartToken)

		if err != nil {
			return cart, err
		}

		item, err = r.addItem(context.TODO(), cart, item)

		if err != nil {
			return cart, err
		}
		cart.Items = append(cart.Items, item)

		return cart, nil
	}

}

func (r *DynamoDBCartRepository) GetCart(cartToken string) (*models.Cart, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		KeyConditionExpression: aws.String("cart_token = :token"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":token": &types.AttributeValueMemberS{Value: cartToken},
		},
	}

	result, err := r.client.Query(context.TODO(), input)

	if err != nil {
		return &models.Cart{}, err
	}

	cart := &models.Cart{}

	for _, item := range result.Items {
		if item["sk"].(*types.AttributeValueMemberS).Value == "cart" {
			err = attributevalue.UnmarshalMap(item, cart)
			if err != nil {
				return cart, err
			}
		} else {
			cartItem := &models.Item{}
			err = attributevalue.UnmarshalMap(item, cartItem)
			if err != nil {
				return cart, err
			}
			cart.Items = append(cart.Items, cartItem)
		}
	}

	return cart, nil
}

func (r *DynamoDBCartRepository) createCart(ctx context.Context, items []*models.Item) (*models.Cart, error) {
	var writeReqs []types.WriteRequest

	cart := &models.Cart{
		Token:     uuid.New().String(),
		SK:        "cart",
		Items:     items,
		CreatedAt: time.Now().String(),
	}

	putItem, err := attributevalue.MarshalMap(cart)
	if err != nil {
		return &models.Cart{}, err
	}

	writeReqs = append(writeReqs, types.WriteRequest{
		PutRequest: &types.PutRequest{
			Item: putItem,
		}})

	for _, item := range cart.Items {
		item.SK = "item:" + uuid.New().String()
		item.CartToken = cart.Token
		item.CreatedAt = time.Now().String()

		putItem, err := attributevalue.MarshalMap(item)
		if err != nil {
			return &models.Cart{}, err
		}

		writeReqs = append(writeReqs, types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: putItem,
			}})
	}

	_, err = r.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{r.tableName: writeReqs}})

	if err != nil {
		return &models.Cart{}, err
	}

	return cart, nil
}

func (r *DynamoDBCartRepository) addItem(ctx context.Context, cart *models.Cart, item *models.Item) (*models.Item, error) {

	item.CartToken = cart.Token
	item.SK = "item:" + uuid.New().String()

	putItem, err := attributevalue.MarshalMap(item)
	if err != nil {
		return &models.Item{}, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      putItem,
	}

	_, err = r.client.PutItem(ctx, input)
	if err != nil {
		return &models.Item{}, err
	}

	return item, nil
}

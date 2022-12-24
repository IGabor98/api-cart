package repositories

import (
	"context"
	"github.com/IGabor98/api-cart/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"time"
)

type DynamoDBCartRepository struct {
	client *dynamodb.Client
}

func NewCartRepository(client *dynamodb.Client) *DynamoDBCartRepository {
	return &DynamoDBCartRepository{client: client}
}

func (r *DynamoDBCartRepository) AddItemToCart(cartToken string, item models.Item) (models.Cart, error) {

	if cartToken == "" {
		cart := models.Cart{
			Token: uuid.New().String(),
			Items: []models.Item{item},
		}

		return r.createCart(context.TODO(), cart)
	} else {
		cart, err := r.GetCart(cartToken)

		if err != nil {
			return cart, err
		}

		cart.Items = append(cart.Items, item)
		return r.updateCart(context.TODO(), cart)
	}

}

func (r *DynamoDBCartRepository) GetCart(cartToken string) (models.Cart, error) {
	primaryKey := map[string]string{
		"token": cartToken}

	pk, err := attributevalue.MarshalMap(primaryKey)
	if err != nil {
		return models.Cart{}, err
	}
	input := &dynamodb.GetItemInput{
		Key:       pk,
		TableName: aws.String("Carts"),
	}

	output, err := r.client.GetItem(context.TODO(), input)
	if err != nil {
		panic(err)
	}

	cart := models.Cart{}
	err = attributevalue.UnmarshalMap(output.Item, &cart)
	if err != nil {
		panic(err)
	}

	return cart, nil
}

func (r *DynamoDBCartRepository) createCart(ctx context.Context, cart models.Cart) (models.Cart, error) {
	putItem, err := attributevalue.MarshalMap(cart)
	if err != nil {
		return models.Cart{}, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Carts"),
		Item:      putItem,
	}

	_, err = r.client.PutItem(ctx, input)
	if err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (r *DynamoDBCartRepository) updateCart(ctx context.Context, cart models.Cart) (models.Cart, error) {

	primaryKey := map[string]string{
		"token": cart.Token}

	pk, err := attributevalue.MarshalMap(primaryKey)
	if err != nil {
		return models.Cart{}, err
	}

	upd := expression.
		Set(expression.Name("updated_at"), expression.Value(time.Now())).
		Set(expression.Name("items"), expression.Value(cart.Items))

	expr, err := expression.NewBuilder().WithUpdate(upd).Build()

	if err != nil {
		return cart, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String("Carts"),
		Key:                       pk,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = r.client.UpdateItem(ctx, input)

	if err != nil {
		return cart, err
	}

	return cart, nil

}

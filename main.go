package main

import (
	"context"
	"github.com/IGabor98/api-cart/controllers"
	"github.com/IGabor98/api-cart/repositories"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {

	cartRepository := repositories.NewCartRepository(CreateLocalClient())
	cartController := &controllers.CartController{
		CartRepository: cartRepository,
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/v1/carts/add-item", cartController.AddItem)
	r.Get("/api/v1/carts/{cartToken}", cartController.GetCart)
	r.Delete("/api/v1/carts/{cartToken}", cartController.DeleteCart)
	r.Delete("/api/v1/carts/{cartToken}/{itemId}", cartController.DeleteItemFromCart)

	http.ListenAndServe(":3000", r)
}

func CreateLocalClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "bmm11p", SecretAccessKey: "pmbrgr", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	products := GetProducts()

	headers := map[string]string{
		"content-type":                "application/json; charset=utf-8",
		"access-control-allow-origin": "*"}

	search := req.QueryStringParameters["search"]
	log.Printf(search)
	if len(search) > 0 {
		products = SearchProduct(search, products)
	}

	res, err := json.Marshal(products)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(res),
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(handler)
}

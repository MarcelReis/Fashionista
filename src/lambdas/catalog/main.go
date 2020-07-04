package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	products := GetProducts()

	headers := map[string]string{
		"content-type":                "application/json; charset=utf-8",
		"cache-control":               "public, max-age=3600",
		"access-control-allow-origin": "*"}

	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	search := request.QueryStringParameters["search"]
	log.Printf("TESTING: %s", search)

	if len(request.QueryStringParameters["search"]) > 0 {
		products = SearchProduct(request.QueryStringParameters["search"], products)
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

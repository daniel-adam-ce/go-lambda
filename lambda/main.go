package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type AuthRequest struct {
	Token string `json:"token"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func someIDPRequest(token string) bool {
	return token == "test"
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req AuthRequest
	if err := json.Unmarshal([]byte(event.Body), &req); err != nil {
		log.Printf("Invalid request body: %v", err)
		return errorResponse(http.StatusBadRequest, "Invalid request body"), nil
	}

	log.Printf("Received token: %s", req.Token)

	if !someIDPRequest(req.Token) {
		return successResponse(http.StatusForbidden, ResponseBody{Message: "invalid token"})
	}

	return successResponse(http.StatusOK, ResponseBody{Message: "success!"})
}

func successResponse(status int, body ResponseBody) (events.APIGatewayProxyResponse, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return errorResponse(http.StatusInternalServerError, "Internal server error"), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(payload),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func errorResponse(status int, message string) events.APIGatewayProxyResponse {
	body := ResponseBody{Message: message}
	payload, _ := json.Marshal(body)

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(payload),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func main() {
	lambda.Start(handler)
}

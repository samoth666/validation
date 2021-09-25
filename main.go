package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/personal/validation/models"
	"github.com/personal/validation/mutant"
	"github.com/personal/validation/storage"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
)

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	dna := new(models.Dna)

	err := json.Unmarshal([]byte(req.Body), dna)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if dna.DnaSecuences == nil {
		return clientError(http.StatusBadRequest)
	}

	dnaInput := dna.DnaSecuences
	mu := mutant.IsMutant(dnaInput)
	dna.Id = fmt.Sprint(dnaInput)
	dna.IsMutant = mu

	err = storage.PutItem(dna)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/mutant?id=%s", dna.Id)},
	}, nil
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func main() {
	lambda.Start(handler)
}

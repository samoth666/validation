package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
)

func TestAPIGateway_success(t *testing.T) {
	c := require.New(t)

	inputJSON, err := ioutil.ReadFile("samples/apigateway-success.json")
	c.Nil(err)

	var req events.APIGatewayProxyRequest
	err = json.Unmarshal(inputJSON, &req)
	c.Nil(err)

	response, err := handler(req)
	c.Nil(err)
	c.NotEmpty(response.Body)
}

func TestAPIGateway_empty(t *testing.T) {
	c := require.New(t)

	inputJSON, err := ioutil.ReadFile("samples/apigateway-empty.json")
	c.Nil(err)

	var req events.APIGatewayProxyRequest
	err = json.Unmarshal(inputJSON, &req)
	c.Nil(err)

	response, err := handler(req)
	c.Nil(err)
	c.Equal(response.Body, "Method Not Allowed")
}

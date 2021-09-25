package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/personal/validation/models"
)

var (
	db = dynamodb.New(session.New(),
		aws.NewConfig().WithRegion("us-east-1"),
	)
)

// PutItem add a dna secuence to DynamoDB.
func PutItem(dna *models.Dna) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Dna"),
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(dna.Id),
			},
			"DnaSecuences": {
				SS: aws.StringSlice(dna.DnaSecuences),
			},
			"IsMutant": {
				BOOL: aws.Bool(dna.IsMutant),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}

package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bitsmag/accSlLog/src/accSlLog/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ReadBalance returns the current ballance of the account
func ReadBalance() (float64, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Acc_balances"),
		Key: map[string]*dynamodb.AttributeValue{
			"AccId": {S: aws.String("default")},
		},
	})

	if err != nil {
		return 0, err
	}

	balanceObj := balanceObj{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &balanceObj)
	balance, err := strconv.ParseFloat(balanceObj.Balance, 64)

	if err != nil {
		return 0, fmt.Errorf("Failed to unmarshal and parse balance: %v", err)
	}

	return balance, nil
}

// ReadLogs returns all logEntries
func ReadLogs() ([]types.LogEntry, error) {
	tablenameLog := os.Getenv("TABLENAME_LOG")
	if len(tablenameLog) == 0 {
		tablenameLog = "Acc_logs"
	}

	var entries []types.LogEntry

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		TableName: aws.String(tablenameLog),
	}
	result, err := svc.Scan(input)

	if err != nil {
		return entries, err
	}

	logObjs := []logObj{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &logObjs)

	if err != nil {
		return entries, fmt.Errorf("Failed to unmarshal and parse logs: %v", err)
	}

	for _, logObj := range logObjs {
		entries = append(entries, logObj.LogEntry)
	}

	return entries, nil
}

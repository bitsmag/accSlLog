package main

import (
	"context"
	"fmt"

	"acc_lambda_log/cmd"
	"acc_lambda_log/types"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Order    string `json:"order"`
	Date     string `json:"date"`
	Category string `json:"category"`
}

func HandleRequest(ctx context.Context, event Event) (string, error) {
	var order types.Order
	if event.Order != "" {
		order.Set(event.Order)
	}
	var date types.Date
	if event.Date != "" {
		date.Set(event.Date)
	}
	var category types.Category
	if event.Category != "" {
		date.Set(event.Category)
	}

	resp, err := cmd.LogCmdHandler(order, date, category)

	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("Some error occured on the server"), err
	} else {
		return fmt.Sprintf(resp), nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}

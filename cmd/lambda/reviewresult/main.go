package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shibataka000/heimdall/internal/aws/bedrock"
)

func handler(ctx context.Context, event *bedrock.ActionGroupRequest) (*bedrock.ActionGroupResponse, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return bedrock.NewActionGroupResponse(event, 500, nil), err
	}
	log.Println(string(b))
	return bedrock.NewActionGroupResponse(event, 201, nil), nil
}

func main() {
	lambda.Start(handler)
}

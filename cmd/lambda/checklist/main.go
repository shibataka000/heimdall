package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shibataka000/heimdall/internal/aws/bedrock"

	wa "github.com/shibataka000/heimdall/internal/checklist/awswellarchitectedframework"
)

func handler(ctx context.Context, event *bedrock.ActionGroupRequest) (*bedrock.ActionGroupResponse, error) {
	checklistID, err := event.GetParameter("checklistId")
	if err != nil {
		return bedrock.NewActionGroupResponse(event, 400, nil), err
	}
	requirementID, err := event.GetParameter("requirementId")
	if err != nil {
		return bedrock.NewActionGroupResponse(event, 400, nil), err
	}
	if checklistID != "aws-well-architected-framework" {
		return bedrock.NewActionGroupResponse(event, 400, nil), fmt.Errorf("checklist '%s' is not supported", checklistID)
	}
	requirement, err := wa.GetRequirement(requirementID)
	if err != nil {
		return bedrock.NewActionGroupResponse(event, 404, nil), err
	}
	b, err := json.Marshal(requirement)
	if err != nil {
		return bedrock.NewActionGroupResponse(event, 500, nil), err
	}
	responseBody := map[string]bedrock.ActionGroupResponseResponseResponseBody{
		"application/json": {
			Body: string(b),
		},
	}
	return bedrock.NewActionGroupResponse(event, 200, responseBody), nil
}

func main() {
	lambda.Start(handler)
}

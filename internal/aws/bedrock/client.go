package bedrock

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent"
	agenttypes "github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	agentruntimetypes "github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"github.com/google/uuid"
)

const (
	agentAliasID = "TSTALIASID"
)

type Client struct {
	agent        *bedrockagent.Client
	agentruntime *bedrockagentruntime.Client
}

func NewClient(ctx context.Context) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		agent:        bedrockagent.NewFromConfig(cfg),
		agentruntime: bedrockagentruntime.NewFromConfig(cfg),
	}, nil
}

func (c *Client) Ingest(ctx context.Context, knowledgeBaseID string, dataSourceID string) error {
	// Start the ingestion job.
	startIngestionJobOutput, err := c.agent.StartIngestionJob(ctx, &bedrockagent.StartIngestionJobInput{
		DataSourceId:    aws.String(dataSourceID),
		KnowledgeBaseId: aws.String(knowledgeBaseID),
	})
	if err != nil {
		return err
	}

	// Wait for the ingestion job to complete.
	for {
		getIngestionJobOutput, err := c.agent.GetIngestionJob(ctx, &bedrockagent.GetIngestionJobInput{
			DataSourceId:    startIngestionJobOutput.IngestionJob.DataSourceId,
			IngestionJobId:  startIngestionJobOutput.IngestionJob.IngestionJobId,
			KnowledgeBaseId: startIngestionJobOutput.IngestionJob.KnowledgeBaseId,
		})
		if err != nil {
			return err
		}
		switch getIngestionJobOutput.IngestionJob.Status {
		case agenttypes.IngestionJobStatusComplete:
			return nil
		case agenttypes.IngestionJobStatusFailed:
			return errors.New("ingestion job failed")
		case agenttypes.IngestionJobStatusStopped:
			return errors.New("ingestion job stopped")
		default:
			time.Sleep(10 * time.Second)
		}
	}
}

func (c *Client) InvokeAgent(ctx context.Context, agentID string, inputText string) ([]byte, error) {
	sessionID, err := newSessionID()
	if err != nil {
		return nil, err
	}
	output, err := c.agentruntime.InvokeAgent(ctx, &bedrockagentruntime.InvokeAgentInput{
		AgentAliasId: aws.String(agentAliasID),
		AgentId:      aws.String(agentID),
		SessionId:    aws.String(sessionID),
		EnableTrace:  aws.Bool(true),
		InputText:    aws.String(inputText),
	})
	if err != nil {
		return nil, err
	}
	response := []byte{}
	stream := output.GetStream()
	for event := range stream.Events() {
		switch resp := event.(type) {
		case *agentruntimetypes.ResponseStreamMemberChunk:
			response = append(response, resp.Value.Bytes...)
		case *agentruntimetypes.ResponseStreamMemberTrace:
			b, err := json.Marshal(resp)
			if err != nil {
				return nil, err
			}
			log.Println(string(b))
		}
	}
	return response, stream.Close()
}

func newSessionID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

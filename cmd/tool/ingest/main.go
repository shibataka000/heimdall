// Ingest the documents in the data source into the knowledge base.
package main

import (
	"context"
	"os"

	"github.com/shibataka000/heimdall/internal/aws/bedrock"
	"github.com/spf13/cobra"
)

// ingest the documents in the data source into the knowledge base.
func ingest(ctx context.Context, knowledgeBaseID, dataSourceID string) error {
	client, err := bedrock.NewClient(ctx)
	if err != nil {
		return err
	}
	return client.Ingest(ctx, knowledgeBaseID, dataSourceID)
}

func main() {
	var (
		knowledgeBaseID string
		dataSourceID    string
	)

	command := &cobra.Command{
		Use:   "ingest",
		Short: "Ingest the documents in the data source into the knowledge base.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return ingest(cmd.Context(), knowledgeBaseID, dataSourceID)
		},
		SilenceUsage: true,
	}

	command.Flags().StringVar(&knowledgeBaseID, "knowledge-base-id", "", "The unique identifier of the knowledge base for the data ingestion job.")
	command.Flags().StringVar(&dataSourceID, "data-source-id", "", "The unique identifier of the data source you want to ingest into your knowledge base.")

	for _, flag := range []string{"knowledge-base-id", "data-source-id"} {
		if err := command.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	if command.ExecuteContext(context.Background()) != nil {
		os.Exit(1)
	}
}

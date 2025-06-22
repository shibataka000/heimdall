package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shibataka000/heimdall/internal/aws/bedrock"
	"github.com/spf13/cobra"
)

func main() {
	var (
		agentID       string
		checklistID   string
		requirementID string
	)

	command := &cobra.Command{
		Use:   "review",
		Short: "Review the design documents stored in the knowledge base according to the requirements in the checklist.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := bedrock.NewClient(cmd.Context())
			if err != nil {
				return err
			}
			prompt := fmt.Sprintf("ナレッジベースに格納された設計書がチェックリスト '%s' の要件 '%s' を満たしているか判定してください。", checklistID, requirementID)
			resp, err := client.InvokeAgent(cmd.Context(), agentID, prompt)
			if err != nil {
				return err
			}
			println(string(resp))
			return nil
		},
		SilenceUsage: true,
	}

	command.Flags().StringVar(&agentID, "agent-id", "", "The unique identifier of the agent to invoke.")
	command.Flags().StringVar(&checklistID, "checklist-id", "", " The unique identifier of the checklist.")
	command.Flags().StringVar(&requirementID, "requirement-id", "", "The unique identifier of the requirement.")

	for _, flag := range []string{"agent-id", "checklist-id", "requirement-id"} {
		if err := command.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	if command.ExecuteContext(context.Background()) != nil {
		os.Exit(1)
	}
}

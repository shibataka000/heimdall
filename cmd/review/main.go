// Review the design documents stored in the knowledge base according to the requirements in the checklist.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shibataka000/heimdall/internal/aws/bedrock"
	wa "github.com/shibataka000/heimdall/internal/checklist/awswellarchitectedframework"
	"github.com/spf13/cobra"
)

func main() {
	var (
		agentID string
	)

	command := &cobra.Command{
		Use:   "review <url>",
		Short: "Review the design documents stored in the knowledge base according to the requirements in the checklist.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := bedrock.NewClient(cmd.Context())
			if err != nil {
				return err
			}
			requirements, err := wa.Requirements()
			if err != nil {
				return err
			}
			for _, requirement := range requirements {
				prompt, err := wa.Prompt(requirement)
				if err != nil {
					return err
				}
				response, err := client.InvokeAgent(cmd.Context(), agentID, prompt)
				if err != nil {
					return err
				}
				fmt.Println(string(response))
			}
			return nil
		},
		SilenceUsage: true,
	}

	command.Flags().StringVar(&agentID, "agent-id", "", "The unique identifier of the agent to invoke.")

	for _, flag := range []string{"agent-id"} {
		if err := command.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	if command.ExecuteContext(context.Background()) != nil {
		os.Exit(1)
	}
}

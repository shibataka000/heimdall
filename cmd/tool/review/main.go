// Review the design documents stored in the knowledge base according to the requirements in the checklist.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"time"

	"github.com/shibataka000/heimdall/internal/aws/bedrock"
	wa "github.com/shibataka000/heimdall/internal/checklist/awswellarchitectedframework"
	"github.com/spf13/cobra"
)

const (
	interval = 30 * time.Second
)

// review the design documents stored in the knowledge base according to the requirements in the checklist.
func review(ctx context.Context, agentID string, filter *regexp.Regexp) error {
	client, err := bedrock.NewClient(ctx)
	if err != nil {
		return err
	}
	requirements := slices.DeleteFunc(slices.Clone(wa.Requirements), func(req wa.Requirement) bool {
		return !filter.MatchString(string(req.Title))
	})
	for _, requirement := range requirements {
		prompt, err := wa.NewPrompt(requirement).Render()
		if err != nil {
			return err
		}
		log.Println(prompt)
		resp, err := client.InvokeAgent(ctx, agentID, prompt)
		if err != nil {
			return err
		}
		result, err := wa.NewReviewResult(requirement, resp)
		if err != nil {
			return err
		}
		fmt.Println(result.String())
		time.Sleep(interval)
	}
	return nil
}

func main() {
	var (
		agentID      string
		filterRegexp string
	)

	command := &cobra.Command{
		Use:   "review",
		Short: "Review the design documents stored in the knowledge base according to the requirements in the checklist.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			filter, err := regexp.Compile(filterRegexp)
			if err != nil {
				return err
			}
			return review(cmd.Context(), agentID, filter)
		},
		SilenceUsage: true,
	}

	command.Flags().StringVar(&agentID, "agent-id", "", "The unique identifier of the agent to invoke.")
	command.Flags().StringVar(&filterRegexp, "filter", ".*", " A regular expression to filter requirements by title.")

	for _, flag := range []string{"agent-id"} {
		if err := command.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	if command.ExecuteContext(context.Background()) != nil {
		os.Exit(1)
	}
}

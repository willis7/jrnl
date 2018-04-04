package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:        "remove",
	Short:      "Removes a journal entry.",
	SuggestFor: []string{"delete", "rm"},
	Args:       cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, date := range args {
			err := client.DeleteEntry(date)
			if err != nil {
				return fmt.Errorf("failed to delete \"%s\". error: %s", date, err)
			}
			fmt.Printf("deleted \"%s\"\n", date)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

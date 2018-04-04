package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your entries.",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := client.AllEntries()
		if err != nil {
			return fmt.Errorf("failed to retrieve entries: %s", err)
		}
		if len(entries) == 0 {
			fmt.Println("you have no entries")
			return nil
		}
		fmt.Println("Here's your jrnl entries:")
		for i, entry := range entries {
			fmt.Printf("%d. %s %s\n", i+1, entry.Date, entry.Text)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

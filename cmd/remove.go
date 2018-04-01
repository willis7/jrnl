package cmd

import (
	"fmt"

	"github.com/willis7/jrnl/db"

	"github.com/spf13/cobra"
)

// CreateRemoveCmd closes over a client and adds a remove command
func CreateRemoveCmd(client db.IBoltClient) {
	// removeCmd represents the remove command
	var removeCmd = &cobra.Command{
		Use:        "remove",
		Short:      "Removes a journal entry.",
		SuggestFor: []string{"delete", "rm"},
		Args:       cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, date := range args {
				err := client.DeleteEntry(date)
				if err != nil {
					fmt.Printf("failed to delete \"%s\". error: %s", date, err)
				} else {
					fmt.Printf("deleted \"%s\"\n", date)
				}
			}
		},
	}
	RootCmd.AddCommand(removeCmd)
}

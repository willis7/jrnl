package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/willis7/jrnl/db"
)

func CreateListCmd(client db.IBoltClient) {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all of your entries.",
		Run: func(cmd *cobra.Command, args []string) {
			entries, err := client.AllEntries()
			if err != nil {
				fmt.Println("failed to retrieve entries:", err)
				os.Exit(1)
			}
			if len(entries) == 0 {
				fmt.Println("you have no entries")
				return
			}
			fmt.Println("Here's your jrnl entries:")
			for i, entry := range entries {
				fmt.Printf("%d. %s\n", i+1, entry.Value)
			}
		},
	}

	RootCmd.AddCommand(listCmd)
}

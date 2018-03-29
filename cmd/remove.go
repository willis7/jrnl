package cmd

import (
	"fmt"
	"strconv"

	"github.com/willis7/jrnl/db"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:        "remove",
	Short:      "Removes a journal entry.",
	SuggestFor: []string{"delete", "rm"},
	Args:       cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		entries, err := db.AllEntries()
		if err != nil {
			fmt.Println("failed to retrieve entries:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(entries) {
				fmt.Println("invalid entry number:", id)
				continue
			}
			entry := entries[id-1]
			err := db.DeleteEntry(entry.Key)
			if err != nil {
				fmt.Printf("failed to delete \"%d\". error: %s", id, err)
			} else {
				fmt.Printf("deleted \"%d\"", id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

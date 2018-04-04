package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	t "github.com/willis7/jrnl/time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an entry in your journal,",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		date := keywordToDate(args[0])

		// Take a journal entry from the user input
		// and store it in the db
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		id, err := client.CreateEntry(date, text)
		if err != nil {
			fmt.Println("failed to create entry:", err)
			os.Exit(1)
		}
		fmt.Printf("created; %d. %s %s", id, date, text)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

func keywordToDate(word string) string {
	var date string
	lWord := strings.ToLower(word)
	if lWord == "today" {
		date = t.Today()
	} else if lWord == "yesterday" {
		date = t.Yesterday()
	} else {
		date = lWord
	}
	return date
}

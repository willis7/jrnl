package cmd

import (
	"bufio"
	"fmt"
	"os"

	t "github.com/willis7/jrnl/time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an entry in your journal,",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		date := t.KeywordToDate(args[0])
		text, err := getUserInput()
		_, err = client.CreateEntry(date, text)
		if err != nil {
			return fmt.Errorf("failed to create entry: %s", err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

func getUserInput() (string, error) {
	// Take a journal entry from the user input
	// and store it in the db
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to get user input: %s", err)
	}
	return text, nil
}

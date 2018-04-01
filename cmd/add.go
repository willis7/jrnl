package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/willis7/jrnl/db"
	t "github.com/willis7/jrnl/time"

	"github.com/spf13/cobra"
)

// CreateAddCmd closes over a client and adds a add command
func CreateAddCmd(client db.IBoltClient) {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds an entry in your journal,",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Check for date keywords and autofill date
			var date string
			input := strings.ToLower(args[0])
			if input == "today" {
				date = t.Today()
			} else if input == "yesterday" {
				date = t.Yesterday()
			} else {
				date = args[0]
			}

			// Take a journal entry from the user
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

	RootCmd.AddCommand(addCmd)
}

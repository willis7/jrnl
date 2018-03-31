package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/willis7/jrnl/db"

	"github.com/spf13/cobra"
)

func CreateAddCmd(client db.IBoltClient) {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds an entry in your journal,",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			id, err := client.CreateEntry(text)
			if err != nil {
				fmt.Println("failed to create entry:", err)
				os.Exit(1)
			}
			fmt.Printf("created; %d. %s", id, text)
		},
	}

	RootCmd.AddCommand(addCmd)
}

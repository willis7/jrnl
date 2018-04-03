package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/willis7/jrnl/db"
	"github.com/willis7/jrnl/exporter"
)

// CreateExportCmd closes over a client and adds a add command
func CreateExportCmd(client db.IClient) {
	var exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Export a formatted journal",
		Run: func(cmd *cobra.Command, args []string) {
			e, err := client.AllEntries()
			if err != nil {
				fmt.Println("failed to retrieve entries: ", err)
				return
			}
			err = exporter.WriteJSON(e)
			if err != nil {
				fmt.Println("failed to export JSON: ", err)
				return
			}
		},
	}
	RootCmd.AddCommand(exportCmd)
}

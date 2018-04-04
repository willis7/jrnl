package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/willis7/jrnl/exporter"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a formatted journal",
	RunE: func(cmd *cobra.Command, args []string) error {
		e, err := client.AllEntries()
		if err != nil {
			return fmt.Errorf("failed to retrieve entries: %s", err)
		}
		err = exporter.WriteJSON(e)
		if err != nil {
			return fmt.Errorf("failed to export JSON: %s", err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
}

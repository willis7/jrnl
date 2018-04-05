package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/willis7/jrnl/exporter"
)

var setJSON bool
var setXML bool

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a formatted journal",
	RunE: func(cmd *cobra.Command, args []string) error {
		// check the number of set command line flags
		// and error if 0. This will avoid a call to the DB
		// if the user hasnt specified the format they want
		if cmd.Flags().NFlag() == 0 {
			return errors.New("no export format specified. See help for more info")
		}
		e, err := client.AllEntries()
		if err != nil {
			return fmt.Errorf("failed to retrieve entries: %s", err)
		}
		if setJSON {
			err = exporter.WriteJSON(e)
			if err != nil {
				return fmt.Errorf("failed to export JSON: %s", err)
			}
		}

		if setXML {
			err = exporter.WriteXML(e)
			if err != nil {
				return fmt.Errorf("failed to export XML: %s", err)
			}
			return nil
		}
		return errors.New("something is wrong because you should never reach this point")
	},
}

func init() {
	exportCmd.Flags().BoolVarP(&setJSON, "json", "j", false, "write json")
	exportCmd.Flags().BoolVarP(&setXML, "xml", "x", false, "write xml")
	RootCmd.AddCommand(exportCmd)
}

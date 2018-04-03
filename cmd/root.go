package cmd

import (
	"github.com/spf13/cobra"
	"github.com/willis7/jrnl/db"
)

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "jrnl",
	Short: "jrnl is a CLI journal manager",
}

// InitCmd adds the subcommands to the Root command
func InitCmd(client db.IClient) {
	CreateAddCmd(client)
	CreateRemoveCmd(client)
	CreateListCmd(client)
	CreateExportCmd(client)
}

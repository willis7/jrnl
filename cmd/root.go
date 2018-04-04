package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/willis7/jrnl/db"
)

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "jrnl",
	Short: "jrnl is a CLI journal manager",
}

var client db.IClient

func init() {
	home, err := homedir.Dir()
	if err != nil {
		os.Exit(1)
	}
	dbPath := filepath.Join(home, "jrnl.db")
	client, err = db.NewClient("jrnl", dbPath)
	if err != nil {
		if err != nil {
			fmt.Println("failed to create client:", err)
			os.Exit(1)
		}
	}
}

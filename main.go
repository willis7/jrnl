package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/willis7/jrnl/cmd"
	"github.com/willis7/jrnl/db"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		os.Exit(1)
	}
	dbPath := filepath.Join(home, "jrnl.db")
	bolt, err := db.NewBoltClient("jrnl", dbPath)
	if err != nil {
		if err != nil {
			fmt.Println("failed to create client:", err)
			os.Exit(1)
		}
	}
	defer bolt.Close()
	cmd.InitCmd(bolt)
	must(cmd.RootCmd.Execute())

}

// must throws an os.Exit(1) if an error is found
func must(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

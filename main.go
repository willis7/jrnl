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
	must(db.Init(dbPath))
	defer db.Close()
	must(cmd.RootCmd.Execute())

}

// must throws an os.Exit(1) if an error is found
func must(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

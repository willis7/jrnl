package main

import (
	"fmt"
	"os"

	"github.com/willis7/jrnl/cmd"
)

func main() {
	must(cmd.RootCmd.Execute())

}

// must throws an os.Exit(1) if an error is found
func must(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

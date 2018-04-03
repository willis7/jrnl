package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/willis7/jrnl/db"
)

// WriteJSON creates a json file in the same directory of the
// app
func WriteJSON(entries []db.Entry) {
	entriesJSON, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("failed to marshall data:", err)
		return
	}
	err = ioutil.WriteFile("jrnl.json", entriesJSON, 0644)
}

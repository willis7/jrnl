package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/willis7/jrnl/db"
)

// WriteJSON creates a json file in the same directory of the
// app
func WriteJSON(entries []db.Entry) error {
	entriesJSON, err := json.Marshal(entries)
	if err != nil {
		return fmt.Errorf("failed to marshall data: %s", err)
	}
	return ioutil.WriteFile("jrnl.json", entriesJSON, 0644)
}

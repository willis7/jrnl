package exporter

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/willis7/jrnl/db"
)

// WriteXML creates a json file in the same directory of the
// app
func WriteXML(entries []db.Entry) error {
	entriesXML, err := xml.Marshal(entries)
	if err != nil {
		return fmt.Errorf("failed to marshall data: %s", err)
	}
	return ioutil.WriteFile("jrnl.xml", entriesXML, 0644)
}

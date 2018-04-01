package db

import "testing"

type TBoltClient struct{}

func (c *TBoltClient) Init() error                     { return nil }
func (c *TBoltClient) Close()                          {}
func (c *TBoltClient) CreateEntry(string) (int, error) { return 1, nil }
func (c *TBoltClient) DeleteEntry(int) error           { return nil }
func (c *TBoltClient) AllEntries() ([]Entry, error) {
	out := []Entry{Entry{ID: 1, Text: "hello"}}
	return out, nil
}

func TestCreateEntry(t *testing.T) {
	// TODO: write more meaningful tests
	c := TBoltClient{}
	i, err := c.CreateEntry("hello")

	if i != 1 || err != nil {
		t.Fail()
	}
}

func TestDeleteEntry(t *testing.T) {
	// TODO: write more meaningful tests
	c := TBoltClient{}
	err := c.DeleteEntry(1)
	if err != nil {
		t.Fail()
	}
}

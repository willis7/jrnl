package db

import (
	"fmt"

	"github.com/asdine/storm"
)

// Entry represents a Journal entry
type Entry struct {
	Key  int `storm:"id,increment"`
	Date string
	Text string
}

// IBoltClient is an interface which represents behaviour of a BoltClient
type IBoltClient interface {
	Close()
	CreateEntry(string, string) (int, error)
	DeleteEntry(int) error
	AllEntries() ([]Entry, error)
}

// BoltClient implements the IBoltClient interface
type BoltClient struct {
	name   []byte
	db     *storm.DB
	dbPath string
}

// NewBoltClient initialises a database and constructs a BoltClient
func NewBoltClient(name string, dbPath string) (*BoltClient, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return &BoltClient{}, err
	}

	bc := &BoltClient{
		name:   []byte(name),
		dbPath: dbPath,
		db:     db,
	}

	return bc, nil
}

// Close the database once finished
func (c BoltClient) Close() {
	c.db.Close()
}

// CreateEntry adds a new entry to the jrnl database
func (c BoltClient) CreateEntry(date string, text string) (int, error) {
	e := Entry{
		Date: date,
		Text: text,
	}
	err := c.db.Save(&e)
	if err != nil {
		return -1, err
	}
	return e.Key, nil
}

// DeleteEntry removes an entry from the jrnl database
func (c BoltClient) DeleteEntry(key int) error {
	var e Entry
	err := c.db.Find("Key", key, &e)
	if err != nil {
		fmt.Printf("failed to retrieve %d.", key)
	}
	return c.db.DeleteStruct(&e)
}

// AllEntries returns all entries from the database
func (c BoltClient) AllEntries() ([]Entry, error) {
	var entries []Entry
	err := c.db.All(&entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

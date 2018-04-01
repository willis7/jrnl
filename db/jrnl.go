package db

import (
	"github.com/asdine/storm"
)

// Entry represents a Journal entry
type Entry struct {
	ID   int    `storm:"increment"`
	Date string `storm:"index"`
	Text string
}

// IBoltClient is an interface which represents behaviour of a BoltClient
type IBoltClient interface {
	Close()
	CreateEntry(string, string) (int, error)
	DeleteEntry(string) error
	AllEntries() ([]Entry, error)
	FindEntry(string) (Entry, error)
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
	return e.ID, nil
}

// DeleteEntry removes an entry from the jrnl database
// sourced by date
func (c BoltClient) DeleteEntry(date string) error {
	e, err := c.FindEntry(date)
	if err != nil {
		return err
	}
	return c.db.DeleteStruct(&e)
}

// FindEntry retrieves an Entry from the jrnl sourced
// by date
func (c BoltClient) FindEntry(date string) (Entry, error) {
	var e Entry
	err := c.db.One("Date", date, &e)
	if err != nil {
		return e, storm.ErrNotFound
	}
	return e, nil
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

package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// var jrnlBucket = []byte("jrnl")
// var db *bolt.DB

type Entry struct {
	Key   int
	Value string
}

// IBoltClient is an interface which represents behaviour of a BoltClient
type IBoltClient interface {
	Init() error
	Close()
	CreateEntry(string) (int, error)
	DeleteEntry(int) error
	AllEntries() ([]Entry, error)
}

// BoltClient implements the IBoltClient interface
type BoltClient struct {
	name   []byte
	db     *bolt.DB
	dbPath string
}

func NewBoltClient(name string, dbPath string) (*BoltClient, error) {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return &BoltClient{}, err
	}

	return &BoltClient{
		name:   []byte(name),
		dbPath: dbPath,
		db:     db,
	}, nil
}

// Init creates or opens a db defined by dbPath
func (c BoltClient) Init() error {
	return c.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(c.name)
		return err
	})
}

// Close the database once finished
func (c BoltClient) Close() {
	c.db.Close()
}

// CreateEntry adds a new entry to the jrnl database
func (c BoltClient) CreateEntry(text string) (int, error) {
	var id int
	err := c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.name)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(text))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// DeleteEntry removes an entry from the jrnl database
func (c BoltClient) DeleteEntry(key int) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.name)
		return b.Delete(itob(key))
	})
}

// AllEntries returns all entries from the database
func (c BoltClient) AllEntries() ([]Entry, error) {
	var entries []Entry
	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(c.name)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			entries = append(entries, Entry{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi returns an int big endian representation of b.
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

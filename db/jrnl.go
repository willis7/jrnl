package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var jrnlBucket = []byte("jrnl")
var db *bolt.DB

type Entry struct {
	Key   int
	Value string
}

// Init creates or opens a db defined by dbPath
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(jrnlBucket)
		return err
	})
}

// Close the database once finished
func Close() {
	db.Close()
}

// CreateEntry adds a new entry to the jrnl database
func CreateEntry(text string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(jrnlBucket)
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
func DeleteEntry(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(jrnlBucket)
		return b.Delete(itob(key))
	})
}

// AllEntries returns all entries from the database
func AllEntries() ([]Entry, error) {
	var entries []Entry
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(jrnlBucket)
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

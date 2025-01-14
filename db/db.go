package db

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type Database struct {
	db *bolt.DB
}

var defaultBucket = []byte("default")

// newdatabase returns an instance of a db that we can work with
func NewDatabase(dbPath string) (db *Database, closeFunc func() error, err error) {
	boltDb, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, nil, err
	}

	db = &Database{
		db: boltDb,
	}
	closeFunc = boltDb.Close

	if err := db.createDefaultBucket(); err != nil {
		closeFunc()
		return nil, nil, fmt.Errorf("error creating bucket: %w", err)
	}

	return db, closeFunc, nil
}

// SetKey sets the key to the requested value, or returns error
func (d *Database) SetKey(key string, value []byte) error {
	err := d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		err := b.Put([]byte(key), value)
		return err
	})

	return err
}

func (d *Database) createDefaultBucket() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(defaultBucket)
		return err
	})
}

func (d *Database) GetKey(key string) ([]byte, error) {
	var result []byte
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		result = b.Get([]byte(key))
		// check again!!!!
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

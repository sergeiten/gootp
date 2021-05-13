package database

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger/v2"
)

var db *badger.DB

func init() {
	var err error

	db, err = badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		fmt.Printf("failed to open database: %v\n", err)
		os.Exit(1)
	}
}

func Insert(key, value string) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte(key), []byte(value))
	if err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func All() [][]byte {
	txn := db.NewTransaction(true)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	opts.PrefetchSize = 10
	it := txn.NewIterator(opts)
	defer it.Close()

	var keys [][]byte
	for it.Rewind(); it.Valid(); it.Next() {
		keys = append(keys, it.Item().Key())
	}

	return keys
}

func Get(key string) ([]byte, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(key))
	if err != nil {
		return nil, err
	}

	value, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func Close() error {
	return db.Close()
}

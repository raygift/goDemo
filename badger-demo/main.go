package main

import (
	"log"

	"github.com/Connor1996/badger"
)

func main() {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	err = txn.Set([]byte("answer"), []byte("42"))
	if err != nil {
		print(err)
		return
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		print(err)
		return
	}
	value, err := txn.Get([]byte("answer"))
	if err != nil {
		print(err)
		return
	}
	print(value)
}

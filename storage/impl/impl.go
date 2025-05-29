package impl

import (
	"app/storage/contract"
	"app/storage/impl/postgres"
)

// Current storage implementation
var Impl contract.Storage

// Initialize current storage implementation
func init() {
	var err error
	Impl, err = postgres.New()
	if err != nil {
		panic("Can't create connection to database")
	}
}

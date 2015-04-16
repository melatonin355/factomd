// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package database_test

import (
	"fmt"

	"github.com/FactomProject/btcd/chaincfg"
	"github.com/FactomProject/btcd/database"
	_ "github.com/FactomProject/btcd/database/memdb"
	"github.com/FactomProject/btcutil"
)

// This example demonstrates creating a new database and inserting the genesis
// block into it.
func ExampleCreateDB() {
	// Notice in these example imports that the memdb driver is loaded.
	// Ordinarily this would be whatever driver(s) your application
	// requires.
	// import (
	//	"github.com/btcsuite/btcd/database"
	// 	_ "github.com/btcsuite/btcd/database/memdb"
	// )

	// Create a database and schedule it to be closed on exit.  This example
	// uses a memory-only database to avoid needing to write anything to
	// the disk.  Typically, you would specify a persistent database driver
	// such as "leveldb" and give it a database name as the second
	// parameter.
	db, err := database.CreateDB("memdb")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Insert the main network genesis block.
	genesis := btcutil.NewBlock(chaincfg.MainNetParams.GenesisBlock)
	newHeight, err := db.InsertBlock(genesis)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("New height:", newHeight)

	// Output:
	// New height: 0
}

// exampleLoadDB is used in the example to elide the setup code.
func exampleLoadDB() (database.Db, error) {
	db, err := database.CreateDB("memdb")
	if err != nil {
		return nil, err
	}

	// Insert the main network genesis block.
	genesis := btcutil.NewBlock(chaincfg.MainNetParams.GenesisBlock)
	_, err = db.InsertBlock(genesis)
	if err != nil {
		return nil, err
	}

	return db, err
}

// This example demonstrates querying the database for the most recent best
// block height and hash.
func ExampleDb_newestSha() {
	// Load a database for the purposes of this example and schedule it to
	// be closed on exit.  See the CreateDB example for more details on what
	// this step is doing.
	db, err := exampleLoadDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	latestHash, latestHeight, err := db.NewestSha()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Latest hash:", latestHash)
	fmt.Println("Latest height:", latestHeight)

	// Output:
	// Latest hash: 320442d7fb5873368119cac11f0d0d98be58ada384194b69c26dad8a84d4fe2f
	// Latest height: 0
}

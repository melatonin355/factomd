// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package test

import (
    "encoding/hex"
    "encoding/binary"
    "fmt"
    "time"
    cv "strconv"
    fct "github.com/FactomProject/factoid"
    "github.com/FactomProject/factoid/state/stateinit"
    "github.com/FactomProject/factoid/state"
    "github.com/FactomProject/factoid/database"
    "github.com/FactomProject/factoid/wallet"
    "github.com/agl/ed25519"
    "math/rand"
    "testing"
)

var _ = time.Second
var _ = state.FactoidState{}
var _ = hex.EncodeToString
var _ = fmt.Printf
var _ = ed25519.Sign
var _ = rand.New
var _ = binary.Write
var _ = fct.Prtln 
var _ = stateinit.GetDatabase
var _ = database.MapDB{}

var fs *Test_state
// sets up teststate.go                                         
func Test_setup_FactoidState (test *testing.T) {
    // Create a Test State
    fs = new(Test_state)
    
    fs.inputAddresses = make([]fct.IAddress,0,10)
    fs.outputAddresses = make([]fct.IAddress,0,10)
    fs.twallet = new(wallet.SCWallet)              // Wallet for our tests
    fs.twallet.Init()
    
    for i:=0; i<10; i++ {
        addr, err := fs.twallet.GenerateAddress([]byte("tes,mbbm,btin_"+cv.Itoa(i)),1,1)
        if err != nil { fct.Prtln(err); test.Fail() }
        fs.inputAddresses = append(fs.inputAddresses,addr)
        fs.outputAddresses = append(fs.outputAddresses,addr)
    }
    for i:=0; i<10; i++ {
        addr, err := fs.twallet.GenerateAddress([]byte("testout_"+cv.Itoa(i)),1,1)
        if err != nil { fct.Prtln(err); test.Fail() }
        fs.outputAddresses = append(fs.outputAddresses,addr)
    }
}


func Test_create_genesis_FactoidState (test *testing.T) {
        
    // Use Bolt DB
    if true {
        fs.SetDB(new(database.MapDB))
        fs.GetDB().Init()
        db := stateinit.GetDatabase("/tmp/fct_test.db")
        fs.GetDB().SetPersist(db)
        fs.GetDB().SetBacker(db)
        fs.GetDB().DoNotPersist(fct.DB_F_BALANCES)
        fs.GetDB().DoNotPersist(fct.DB_EC_BALANCES)
         
    }else{
        fs.SetDB(stateinit.GetDatabase("/tmp/fct_test.db"))
    }
    // Set the price for Factoids
    fs.SetFactoshisPerEC(10000)
   
    err := fs.LoadState()
    if err != nil {
        fct.Prtln(err)
        test.Fail()
        return
    }
    // Create a number of blocks (i)
    for i:=0; i<10; i++ {
        fct.Prt(" ",fs.GetDBHeight())
        // Create a new block
        for j:=0; j<100; j++ {
            t := fs.newTransaction()
            added := fs.AddTransaction(t)
            if !added { 
                fct.Prt("F:",i,"-",j," ",t) 
            }
            time.Sleep(time.Second/100000)
        }
        fs.ProcessEndOfBlock()
        time.Sleep(time.Second)
    }
    fmt.Println("\nDone")
}

func Test_build_blocks_FactoidState (test *testing.T) {
    
    
}



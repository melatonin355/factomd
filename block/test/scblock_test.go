// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package block

import (
	"encoding/binary"
	"fmt"
    cv "strconv"
	"github.com/agl/ed25519"
    sc "github.com/FactomProject/factoid"
    "github.com/FactomProject/factoid/wallet"
    "github.com/FactomProject/factoid/block"
    "math/rand"
	"testing"
    
)

var _ = sc.Prt
var _ = fmt.Printf
var _ = ed25519.Sign
var _ = rand.New
var _ = binary.Write

var fakeAddr sc.IHash = new(sc.Hash)

func newFakeAddr() sc.IAddress {
    fakeAddr = sc.Sha(fakeAddr.Bytes())
    return fakeAddr
}

func Test_create_block(test *testing.T) {
    w := new(wallet.SCWallet)          // make me a wallet
    
    scb := block.NewSCBlock(1000, 0)
	
    for i:=0;i<1;i++ {
        h0,err := w.GenerateAddress([]byte("test "+cv.Itoa(i)+"-0"),1,1)
        if err != nil { test.Fail() }
        h1,err := w.GenerateAddress([]byte("test "+cv.Itoa(i)+"-1"),1,1)
        if err != nil { test.Fail() }
        h2,err := w.GenerateAddress([]byte("test "+cv.Itoa(i)+"-2"),1,1)
        if err != nil { test.Fail() }
        h3,err := w.GenerateAddress([]byte("test "+cv.Itoa(i)+"-3"),1,1)
        if err != nil { test.Fail() }
        h4,err := w.GenerateAddress([]byte("test "+cv.Itoa(i)+"-4"),1,1)
        if err != nil { test.Fail() }
        h5,err := w.GenerateAddress([]byte("test "+cv.Itoa(i)+"-5"),1,1)
        if err != nil { test.Fail() }
        
        t := w.CreateTransaction()
        
        w.AddInput(t,h1,1000000)
        w.AddInput(t,h2,1000000)
        w.AddOutput(t,h3,1000000)
        w.AddOutput(t,h4,500000)
        w.AddECOutput(t,h5,500000)
        w.AddInput(t,h0,0)
        fee, err := t.CalculateFee(1000)
        w.UpdateInput(t,2,h0,fee)
        
        t.SetLockTime(1257894000000000000+uint64(13424323*i))
        signed,err := w.SignInputs(t)
        if err != nil { sc.Prtln("Error found: ",err); test.Fail(); return; }
        if !signed { sc.Prtln("Not valid"); test.Fail(); return; } 
        
        valid, err := scb.AddTransaction(t)
        if err != nil { sc.Prtln("Error found: ",err); test.Fail(); return; }
        if !valid { sc.Prtln("Not valid"); test.Fail(); return; } 
    }
    data,err := scb.MarshalBinary()
    if err != nil {
		sc.PrtStk()
		test.Fail()
	}
	scb2 := block.NewSCBlock(1000, 0)
	_, err = scb2.UnmarshalBinaryData(data)

	if !scb.IsEqual(scb2) {
		sc.PrtStk()
		test.Fail()
	}

}

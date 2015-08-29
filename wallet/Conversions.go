package wallet

import (
	"fmt"
	"github.com/btcsuitereleases/btcutil/base58"
	"github.com/FactomProject/go-bip32"
	"github.com/FactomProject/go-bip39"
	"strings"
)

func MnemonicStringToPrivateKey(mnemonic string) ([]byte, error) {
	mnemonic = strings.ToLower(strings.TrimSpace(mnemonic))
	seed := bip39.NewSeed(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	child, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 7)
	if err != nil {
		return nil, err
	}

	return child.Key, nil
}

func HumanReadiblyPrivateKeyToPrivateKey(human string) ([]byte, error) {
	human = strings.TrimSpace(human)
	base, version, err := base58.CheckDecode(human)
	if err != nil {
		return nil, err
	}

	if version != 0x64 || base[0] != 0x78 {
		return nil, fmt.Errorf("Invalid prefix")
	}

	return base[1:], nil
}

package test

import (
	"encoding/hex"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/crypto"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"testing"
)

func TestEcdsaCreateKey(t *testing.T) {
	priv, pub, err := crypto.GenerateSubstrateKey(crypto.EcdsaType)
	if err != nil {
		t.Log(err)
		return
	}
	address, err := crypto.CreateSubstrateAddress(pub, ss58.ChainXPrefix)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("PrivateKey: ", hex.EncodeToString(priv))
	fmt.Println("Address1: ", address)
	fmt.Println("===============================================")
	pub2, err := crypto.GenerateSubstrateKeyBySeed(priv, crypto.EcdsaType)
	if err != nil {
		t.Log(err)
		return
	}
	address2, err := crypto.CreateSubstrateAddress(pub2, ss58.ChainXPrefix)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("Address2: ", address2)
}

func TestSr25519CreateKey(t *testing.T) {
	priv, pub, err := crypto.GenerateSubstrateKey(crypto.Sr25519Type)
	if err != nil {
		t.Log(err)
		return
	}
	address, err := crypto.CreateSubstrateAddress(pub, ss58.PolkadotPrefix)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("PrivateKey: ", hex.EncodeToString(priv))
	fmt.Println("Address1: ", address)
	fmt.Println("===============================================")
	pub2, err := crypto.GenerateSubstrateKeyBySeed(priv, crypto.Sr25519Type)
	if err != nil {
		t.Log(err)
		return
	}
	address2, err := crypto.CreateSubstrateAddress(pub2, ss58.PolkadotPrefix)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("Address2: ", address2)
}

func TestSign(t *testing.T) {
	priv, pub, err := crypto.GenerateSubstrateKey(crypto.EcdsaType)
	if err != nil {
		t.Log(err)
		return
	}
	sig, err := crypto.Sign(priv, pub, crypto.EcdsaType)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("EcdsaSig: ", hex.EncodeToString(sig))
	fmt.Println(len(sig))
	fmt.Println("=========================================")

	priv2, pub2, err := crypto.GenerateSubstrateKey(crypto.Sr25519Type)
	if err != nil {
		t.Log(err)
		return
	}
	sig2, err := crypto.Sign(priv2, pub2, crypto.Sr25519Type)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("Sr25519Sig: ", hex.EncodeToString(sig2))
	fmt.Println(len(sig2))
}
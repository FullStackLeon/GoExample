package main

import (
	"fmt"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return
	}

	// 根据助记词生成种子
	seed := bip39.NewSeed(mnemonic, "")
	// 生成主私钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return
	}
	fmt.Println("masterKey", masterKey)

	// 生成第一个子私钥和公钥
	child0Key, err := masterKey.NewChildKey(0)
	if err != nil {
		return
	}
	fmt.Println("child0 private key", child0Key)
	fmt.Println("child0 public key", child0Key.PublicKey())

	// 生成第2个子私钥和公钥
	child1Key, err := child0Key.NewChildKey(1)
	if err != nil {
		return
	}
	fmt.Println("child1 private key", child1Key)
	fmt.Println("child1 public key", child1Key.PublicKey())

	// 将*Key转为string
	fmt.Println("child0 private key to string", child0Key.String())

	// 将私钥序列化为字节切片，并反序列化为*Key
	byteSerialize, err := child0Key.Serialize()
	if err != nil {
		return
	}
	fmt.Println("byteSerialize", byteSerialize)
	k1, err := bip32.Deserialize(byteSerialize)
	if err != nil {
		return
	}
	fmt.Println("k1", k1)

	// 将私钥序列化为Base58格式，并反序列化为*Key
	base58Serialize := child0Key.B58Serialize()
	k2, err := bip32.B58Deserialize(base58Serialize)
	fmt.Println("k2", k2)
}

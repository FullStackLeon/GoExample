package main

import (
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

func main() {
	// 生成长度为128位的秘钥熵
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		fmt.Println(" Generated entropy err:", err)
		return
	}
	fmt.Println("Entropy:", entropy, "length:", len(entropy))
	// 生成助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		fmt.Println("Error generating mnemonic:", err)
		return
	}
	fmt.Println("Generated mnemonic:", mnemonic, "length:", len(mnemonic))

	// 验证助记词有效性
	valid := bip39.IsMnemonicValid(mnemonic)
	fmt.Println("Is mnemonic valid?", valid)

	// 提取熵
	entropy, err = bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		fmt.Println("Extracting entropy err:", err)
		return
	}
	fmt.Println("Entropy:", entropy)

	// 生成种子
	seed := bip39.NewSeed(mnemonic, "123")
	fmt.Println("Seed:", seed)

	// 检查助记词的合法性，然后生成种子
	checkMnemonicAndSeed, err := bip39.NewSeedWithErrorChecking(mnemonic, "123")
	if err != nil {
		fmt.Println("An error is returned if the mnemonic is not convertible to a byte array", err)
		return
	}
	fmt.Println("checkMnemonicAndSeed:", checkMnemonicAndSeed)

	// 获取助记词词典
	mnemonicList := bip39.GetWordList()
	fmt.Println("mnemonic directory ", mnemonicList, " length ", len(mnemonicList))
}

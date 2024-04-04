package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	mnemonic := "spot sad wheel shove vintage age express empower poet describe mandate eternal"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	basePath := "m/44'/60'/0'/0"
	addressList := make([]accounts.Account, 0)
	for i := 0; i < 1; i++ {
		path := fmt.Sprintf("%s/%d", basePath, i)
		derivationPath, err := hdwallet.ParseDerivationPath(path)
		if err != nil {
			log.Fatal(err)
		}
		account, err := wallet.Derive(derivationPath, true)
		if err != nil {
			return
		}
		fmt.Println("address hex:", account.Address.Hex())
		addressList = append(addressList, account)
	}
	// 检测wallet是否存在某个账户，go-ethereum-hdwallet是通过检测账户的派送路径是否存在于钱包派生路径中，因此在wallet.Derive(derivationPath, true)时需要将第二个参数设置为true，将派生路径加入到钱包路径中
	fmt.Println("Wallet contains account?", wallet.Contains(addressList[0]))

	// 获取指定账号的派生路径
	path, err := wallet.Path(addressList[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("account derivation path", path)
	// 获取指定账户的私钥
	fmt.Println("Wallet account path", path)
	privateKey, err := wallet.PrivateKey(addressList[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privateKey", privateKey)

	// 获取指定账户私钥的字节表示
	privateKeyByte, err := wallet.PrivateKeyBytes(addressList[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privateKeyByte", privateKeyByte)

	// 获取指定账户私钥的十六进制表示
	privateKeyHex, err := wallet.PrivateKeyHex(addressList[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privateKeyHex", privateKeyHex)

	// 获取指定账户的公钥
	publicKey, err := wallet.PublicKey(addressList[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("publicKey", publicKey)

	// 获取指定账号公钥的字节表示
	publicKeyByte, err := wallet.PublicKeyBytes(addressList[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("publicKeyByte", publicKeyByte)

	// 获取指定账户公钥的十六进制表示
	publicKeyHex, err := wallet.PublicKeyHex(addressList[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("publicKeyHex", publicKeyHex)

	// 签名文本格式的数据
	signData, err := wallet.SignData(addressList[0], "text/plain", []byte("Hello HD Wallet"))
	if err != nil {
		log.Fatal("signData ", addressList[0], err)
	}
	fmt.Println("signData", signData)

	// 签名Hash格式的数据
	signHash, err := wallet.SignHash(addressList[0], []byte("2A5742705C663CA1FE3E43DF5A0C7C42"))
	if err != nil {
		log.Fatal("signHash err:", addressList[0], err)
	}
	fmt.Println("signHash", signHash)

	// 获取钱包状态
	walletStatus, err := wallet.Status()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wallet Status", walletStatus)

	// 将账户解除pin之后，账户将不包含在钱包中
	fmt.Println("Wallet Unpin", wallet.Unpin(addressList[0]))
	fmt.Println("account contains?", wallet.Contains(addressList[0]))
}

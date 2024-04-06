package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	url := "HTTP://127.0.0.1:7545"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal("Ganache client dial failed", err)
	}
	// test private key
	privateKey := "ea1f4a1007ec5468753aebf79da8b3aa342968e0279860d02eb1382b734279e4"
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("private key string convert *ecdsa.PrivateKey failed", err)
	}
	publicKey, ok := priKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("public key type err", publicKey)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	toAddress := common.HexToAddress("0x0B551B9b7d2a182acA3B6944DF5F916C871f7ec1")

	ctx := context.Background()
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal("send nonce got failed", err)
	}

	amount := big.NewInt(1e17)
	gasLimit := uint64(31000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal("gas price got failed", err)
	}
	// 方法1：
	//tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)
	//stx, err := types.SignTx(tx, types.HomesteadSigner{}, priKey)
	//if err != nil {
	//	log.Fatal("Transaction sign failed", err)
	//}

	// 方法2：NewTransaction已弃用，改用SignTx，发现SignTx的gas费更高
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    amount,
		Data:     []byte("transaction test"),
	}
	stx, err := types.SignTx(types.NewTx(txData), types.HomesteadSigner{}, priKey)
	if err != nil {
		log.Fatal("Transaction sign failed", err)
	}
	if err := client.SendTransaction(ctx, stx); err != nil {
		log.Fatal("Send transaction failed", err)
	}

	fmt.Println("Transaction successfully")
}

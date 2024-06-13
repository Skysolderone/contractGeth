package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"v1/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://data-seed-prebsc-1-s1.bnbchain.org:8545")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	defer conn.Close()
	privateKey, err := crypto.HexToECDSA("meta private")
	if err != nil {
		log.Fatal(err)
	}
	chainid, err := conn.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// bind.NewKeyedTransactor(privateKey) 该方式已被废除
	singer, err := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	if err != nil {
		log.Fatal(err)
	}
	token, err := gen.NewGen(common.HexToAddress("0x9f0961F536c623F1432A242d13F49ca587D5104D"), conn)
	symbol, err := token.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}
	value, err := token.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is total %d\n", symbol, value)
	// transfer
	a := common.HexToAddress("0x35A42428a5446E35158b90D6339F8eAaEf95c272")
	b := common.HexToAddress("0x5F39632728F4C43fA51AA2A991F359D666A9c19E")
	aTotal, err := token.BalanceOf(nil, a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("a token:", aTotal)

	// 发送交易需要签名
	tx, err := token.Transfer(singer, b, big.NewInt(1000))
	if err != nil {
		log.Fatal(err)
	}
	receiptet, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receiptet.Status == types.ReceiptStatusSuccessful {
		bTotal, err := token.BalanceOf(nil, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("b total:", bTotal)
	} else {
		fmt.Println("transfer failed")
	}
}

// WS is total 1000000000
// a token: 1000000000
// 1000

// WS is total 1000000000
// a token: 999999000
// b total: 2000

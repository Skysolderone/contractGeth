package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"v1/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://data-seed-prebsc-1-s1.bnbchain.org:8545") // bsctestnet
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	defer conn.Close()
	chainId, err := conn.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("metamakse priveate ") // metamask privatekey
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	dec := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	n := big.NewInt(100000)
	num := new(big.Int).Mul(dec, n)
	addres, tx, instance, err := gen.DeployGen(auth, conn, num)
	if err != nil {
		log.Fatal(err)
	}
	result, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance
	fmt.Println(addres.Hex())
	fmt.Println(result.Hex())
}

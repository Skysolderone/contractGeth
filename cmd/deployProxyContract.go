package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"v1/login"
	"v1/proxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	privateKeyHash := os.Getenv("privateKey")
	conn, err := ethclient.Dial(os.Getenv("bsctestnet"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	chainId, _ := conn.ChainID(ctx)
	privateKey, err := crypto.HexToECDSA(privateKeyHash)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	address, tx, instance, err := login.DeployLogin(auth, conn)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance
	resutl, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatal(err)
	}
	if address != resutl {
		log.Fatal("err not")
	}
	addres2, tx2, instance2, err := proxy.DeployProxy(auth, conn, resutl)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance2
	resutl2, _ := bind.WaitDeployed(ctx, conn, tx2)
	if addres2 != resutl2 {
		log.Fatal("err not")
	}
	fmt.Println(resutl)
	fmt.Println(resutl2)
}

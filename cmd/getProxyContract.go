package main

import (
	"fmt"
	"log"
	"os"

	"v1/proxy"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// ctx := context.Background()
	conn, err := ethclient.Dial(os.Getenv("bsctestnet"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// chainid, _ := conn.ChainID(ctx)
	// privateKey, _ := crypto.HexToECDSA(os.Getenv("privateKey"))
	// ABI := "[{\"inputs\":[],\"name\":\"GetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_data\",\"type\":\"uint256\"}],\"name\":\"SetData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	// praseABI, err := abi.JSON(strings.NewReader(ABI))

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	address := common.HexToAddress("0xf638C3039D7b5BB55fee87F0D05C7d43fef3e9c4")

	// instance := bind.NewBoundContract(address, praseABI, conn, conn, conn)
	// result := []any{}
	// instance.Call(&bind.CallOpts{Context: ctx}, &result, "data")
	// fmt.Println(result...)

	proxy, _ := proxy.NewProxy(address, conn)
	value, _ := proxy.LogicContract(nil)
	fmt.Println(value)
}

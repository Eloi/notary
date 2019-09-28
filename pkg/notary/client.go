package notary

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client ethclient.Client
var account = common.HexToAddress("0xFe5a62707F8934b8C8245dfdc8E6Ce5f2a8A9B63")
var address = common.HexToAddress("0x0d4f9651b432F709CBB5076e35BAC24B7068B2a5")

func InitClient() {
	rpcServer := "HTTP://127.0.0.1:7545"
	if c, err := ethclient.Dial(rpcServer); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		client = *c
	}
}

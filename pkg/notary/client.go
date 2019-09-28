package notary

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/eloi/notary/pkg/notarydapp"
)

var client *ethclient.Client
var dapp *notarydapp.Notarydapp

func InitClient(rpcServer string, address string) {
	if c, err := ethclient.Dial(rpcServer); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		client = c
	}

	if d, err := notarydapp.NewNotarydapp(common.HexToAddress(address), client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		dapp = d
	}
}

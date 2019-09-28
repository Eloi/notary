package notary

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type asset struct {
	hash      [32]byte
	owner     string
	timestamp big.Int
}

func (a *asset) hashString() string {
	return hex.EncodeToString(a.hash[:])
}

func (a *asset) Print() {
	fmt.Printf("hash: %s, timestamp: %v, owner: %s\n", a.hashString(), a.timestamp.String(), a.owner)
}

func getAsset(hash [32]byte) *asset {
	a := asset{}
	timestamp, owner, err := dapp.GetDocumentByHash(&bind.CallOpts{Pending: true}, hash)

	if err != nil {
		fmt.Println("Error retrieving asset:", err)
	} else {
		a.hash = hash
		a.owner = owner
		a.timestamp = *timestamp
	}
	return &a
}

func printAssets(hashes [][32]byte) {
	for _, hash := range hashes {
		a := getAsset(hash)
		a.Print()
	}
}

// Note: To show how we will implement this with goroutines and channels, here is an example.
// It does not work, as it seems that the timestamp (*big.Int) is reused internally at some point in the notarydapp ABI, so
// all the assets have the same value when using coroutines.

func printAssetsCoroutines(hashes [][32]byte) {
	c := make(chan *asset)

	for _, hash := range hashes {
		go func() {
			a := getAsset(hash)
			c <- a
		}()
	}

	for range hashes {
		a := <-c
		a.Print()
	}
}

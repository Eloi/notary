package notary

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func ListByOwner(owner string) {
	fmt.Println("Listing assets for", owner)

	hashes, err := dapp.GetDocumentHashesByOwner(&bind.CallOpts{Pending: true}, owner)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Found %v assets\n\n", len(hashes))

	printAssets(hashes)

}

func ListAll() {
	fmt.Println("Listing all assets")

	it, err := dapp.NotarydappFilterer.FilterDocumentRegistration(&bind.FilterOpts{}, nil, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for it.Next() {
		// Note: Retrieving the event to get back the owner, as it keccak256 hashed
		// The rest of the fields (timestamp, hash) are already in the event
		e := it.Event
		a := getAsset(e.Hash)
		a.Print()
	}
}

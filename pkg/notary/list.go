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
	fmt.Println("WIP, notary dapp must implement GetDocumentHashes")

	// Uncomment below when notary dapp must implement GetDocumentHashes

	// fmt.Println("Listing assets")
	// hashes, err := dapp.GetDocumentHashes(&bind.CallOpts{Pending: true})
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Found %v assets\n\n", len(hashes))
	// printAssets(hashes)
}

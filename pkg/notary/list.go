package notary

import (
	"context"
	"fmt"
	"log"
)

func ListByOwner(owner string) {
	InitClient()

	fmt.Println("ListByOwner", owner)

}
func ListAll() {
	InitClient()
	fmt.Println("ListAll")

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 25893180161173005034
}

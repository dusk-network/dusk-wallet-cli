package main

import (
	"fmt"

	"github.com/dusk-network/dusk-wallet-cli/menu"
)

func main() {
	if err := menu.Start(); err != nil {
		fmt.Println(err)
	}
}

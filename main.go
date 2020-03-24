package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dusk-network/dusk-protobuf/autogen/go/node"
	"github.com/dusk-network/dusk-wallet-cli/prompt"
)

func main() {

	conf := initConfig()

	// Establish a gRPC connection with the node.
	client := newNodeClient()

	if err := client.Connect(conf.RPC); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	// TODO: deferred functions are not run when os.Exit is called
	defer client.Close()

	// Inquire node about its wallet state, so we know which menu to open.
	resp, err := client.c.GetWalletStatus(context.Background(), &node.EmptyRequest{})
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	// If we have no wallet loaded, we open the menu to load or
	// create one.
	if !resp.Loaded {
		if err := prompt.LoadMenu(client.c); err != nil {
			// If we get an error from `LoadMenu`, it means we lost
			// our connection to the node.
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
	}

	// Once loaded, we open the menu for wallet operations.
	if err := prompt.WalletMenu(client.c); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
}

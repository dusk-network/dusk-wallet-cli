package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dusk-network/dusk-protobuf/autogen/go/node"
	"github.com/dusk-network/dusk-wallet-cli/prompt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	initConfig()

	// Establish a gRPC connection with the node.
	client, err := connectToNode()
	if err != nil {
		// TODO: implement checking on intervals up to a limit
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	// Inquire node about its wallet state, so we know which menu to open.
	resp, err := client.GetWalletStatus(context.Background(), &node.EmptyRequest{})
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	// If we have no wallet loaded, we open the menu to load or
	// create one.
	if !resp.Loaded {
		if err := prompt.LoadMenu(client); err != nil {
			// If we get an error from `LoadMenu`, it means we lost
			// our connection to the node.
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
	}

	// Once loaded, we open the menu for wallet operations.
	if err := prompt.WalletMenu(client); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName("dusk")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.dusk/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stdout, "Config file not found. Please place dusk-wallet-cli in the same directory as your dusk.toml file.")
		os.Exit(0)
	}
}

func connectToNode() (node.NodeClient, error) {
	conn, err := grpc.Dial(viper.Get("rpc.address").(string))
	if err != nil {
		return nil, err
	}

	return node.NewNodeClient(conn), nil
}

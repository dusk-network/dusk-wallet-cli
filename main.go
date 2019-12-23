package main

import (
	"fmt"
	"os"

	"github.com/dusk-network/dusk-wallet-cli/prompt"
	"github.com/dusk-network/dusk-wallet-cli/rpc"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	// First, establish that the node is running by asking it about
	// it's wallet state.
	loaded, err := rpc.IsWalletLoaded()
	if err != nil {
		// TODO: implement checking on intervals up to a limit
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	// If we have no wallet loaded, we open the menu to load or
	// create one.
	if loaded == "false" {
		if err := prompt.LoadMenu(); err != nil {
			// If we get an error from `LoadMenu`, it means we lost
			// our connection to the node.
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
	}

	// Once loaded, we open the menu for wallet operations.
	if err := prompt.WalletMenu(); err != nil {
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

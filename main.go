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

	for {
		// First, establish that the node is running by asking it about
		// it's wallet state.
		loaded, err := rpc.IsWalletLoaded()
		if err != nil {
			// TODO: implement checking on intervals up to a limit
			continue
		}

		// If we have no wallet loaded, we open the menu to load or
		// create one.
		if loaded == "0" {
			if err := prompt.LoadMenu(); err != nil {
				// If we get an error from `LoadMenu`, it means we lost
				// our connection to the node. We will restart the loop
				// to attempt to regain our connection.
				fmt.Fprintln(os.Stdout, err.Error())
				continue
			}
		}

		// Once loaded, we open the menu for wallet operations.
		if err := prompt.WalletMenu(); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
		}
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

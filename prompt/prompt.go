package prompt

import (
	"fmt"
	"os"

	"github.com/dusk-network/dusk-wallet-cli/rpc"
	"github.com/manifoldco/promptui"
)

// LoadMenu opens the prompt for loading a wallet. Returns an error
// in case of a connectivity problem with the node.
func LoadMenu() error {
	for {
		prompt := promptui.Select{
			Label: "Select action",
			Items: []string{"Load Wallet", "Create Wallet", "Load Wallet From Seed", "Exit"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			// Prompts should always succeed.
			panic(err)
		}

		var resp string
		switch result {
		case "Load Wallet":
			resp, err = loadWallet()
		case "Create Wallet":
			resp, err = createWallet()
		case "Load Wallet From Seed":
			resp, err = loadFromSeed()
		case "Exit":
			os.Exit(0)
		}

		// On network errors, we return them to the main function.
		// This will re-establish the connection if possible.
		if nerr, ok := err.(*rpc.NetworkError); ok {
			return nerr
		}

		fmt.Fprintln(os.Stdout, resp)
		// If the method call was successful, we can break out and
		// go to the wallet menu.
		if err == nil {
			return nil
		}
	}
}

// WalletMenu opens the prompt for doing wallet operations. Returns an
// rpc.NetworkError in case of request failure.
func WalletMenu() error {
	for {
		prompt := promptui.Select{
			Label: "Select action",
			Items: []string{"Transfer DUSK", "Stake DUSK", "Bid DUSK", "Show Balance", "Show Address", "Exit"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		var resp string
		switch result {
		case "Transfer DUSK":
			resp, err = transferDusk()
		case "Stake DUSK":
			resp, err = stakeDusk()
		case "Bid DUSK":
			resp, err = bidDusk()
		case "Show Balance":
			resp, err = rpc.GetBalance()
		case "Show Address":
			resp, err = rpc.GetAddress()
		case "Exit":
			os.Exit(0)
		}

		if nerr, ok := err.(*rpc.NetworkError); ok {
			return nerr
		}

		fmt.Fprintln(os.Stdout, resp)
		// We don't check for any other error type. Whether or not the
		// method call failed, we still want to return to this same
		// menu in the end.
	}
}

package prompt

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// LoadMenu opens the prompt for loading a wallet. Returns an error
// in case of a connectivity problem with the node.
func LoadMenu() error {
	// Looping until user chooses "Exit", successfully loads a wallet,
	// or a connectivity error occurs.
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

		switch result {
		case "Load Wallet":
			return loadWallet()
		case "Create Wallet":
			return createWallet()
		case "Load Wallet From Seed":
			return loadFromSeed()
		case "Exit":
			// Simply exiting here saves us from having to write a
			// lot of complicated handling code in the main function.
			os.Exit(0)
		}
	}
}

func WalletMenu() error {
	// Looping until user chooses "Exit", or a connectivity error
	// occurs.
	for {
		prompt := promptui.Select{
			Label: "Select action",
			Items: []string{"Transfer DUSK", "Stake DUSK", "Bid DUSK", "Show Balance", "Show Address", "Exit"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		switch result {
		case "Transfer DUSK":
			return transferDusk()
		case "Stake DUSK":
			return stakeDusk()
		case "Bid DUSK":
			return bidDusk()
		case "Show Balance":
			balance, err := client.GetBalance()
			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, balance)
		case "Show Address":
			address, err := client.GetAddress()
			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, address)
		case "Exit":
			os.Exit(0)
		}
	}
}

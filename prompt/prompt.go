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

func WalletMenu() {
	// Looping until user chooses "Exit", or a connectivity error
	// occurs.
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
		fmt.Println("Hello, world!")
	case "Stake DUSK":
		fmt.Println("Hello, world!")
	case "Bid DUSK":
		fmt.Println("Hello, world!")
	case "Show Balance":
		fmt.Println("Hello, world!")
	case "Show Address":
		fmt.Println("Hello, world!")
	case "Exit":
		os.Exit(0)
	}
}

package prompt

import (
	"errors"

	"github.com/dusk-network/dusk-wallet-cli/rpc"
	"github.com/manifoldco/promptui"
)

func loadWallet() (string, error) {
	pw := getPassword()
	return rpc.LoadWallet(pw)
}

func createWallet() (string, error) {
	pw := getPassword()
	return rpc.CreateWallet(pw)
}

func loadFromSeed() (string, error) {
	validate := func(input string) error {
		if len(input) < 64 {
			return errors.New("Seed must be 64 characters or more")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Seed",
		Validate: validate,
	}

	seed, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	pw := getPassword()
	return rpc.LoadWalletFromSeed(seed, pw)
}

func getPassword() string {
	prompt := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}

	pw, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return pw
}

package prompt

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func loadWallet() error {
	pw := getPassword()
	return client.LoadWallet(pw)
}

func createWallet() error {
	pw := getPassword()
	return client.LoadWallet(pw)
}

func loadFromSeed() error {
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
	client.LoadWalletFromSeed(seed, pw)
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

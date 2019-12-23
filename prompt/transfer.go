package prompt

import (
	"strconv"

	"github.com/dusk-network/dusk-wallet-cli/rpc"
	"github.com/dusk-network/dusk-wallet/key"
	"github.com/manifoldco/promptui"
)

func transferDusk() (string, error) {
	amount := getAmount()

	validateAddress := func(input string) error {
		address := key.PublicAddress(input)
		// TODO: use netprefix inferred from config
		if _, err := address.ToKey(2); err != nil {
			return err
		}

		return nil
	}

	addressPrompt := promptui.Prompt{
		Label:    "Address",
		Validate: validateAddress,
	}

	address, err := addressPrompt.Run()
	if err != nil {
		return "", err
	}

	return rpc.TransferDUSK(amount, address)
}

func bidDusk() (string, error) {
	amount := getAmount()
	lockTime := getLockTime()
	return rpc.BidDUSK(amount, lockTime)
}

func stakeDusk() (string, error) {
	amount := getAmount()
	lockTime := getLockTime()
	return rpc.StakeDUSK(amount, lockTime)
}

func getAmount() string {
	validate := func(input string) error {
		if _, err := strconv.ParseFloat(input, 64); err != nil {
			return err
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Amount",
		Validate: validate,
	}

	amount, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return amount
}

func getLockTime() string {
	validate := func(input string) error {
		if _, err := strconv.Atoi(input); err != nil {
			return err
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Lock Time",
		Validate: validate,
	}

	lockTime, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return lockTime
}

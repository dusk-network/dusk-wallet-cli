package prompt

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dusk-network/dusk-protobuf/autogen/go/node"
	"github.com/dusk-network/dusk-wallet/v2/wallet"
	"github.com/manifoldco/promptui"
)

// LoadMenu opens the prompt for loading a wallet.
func LoadMenu(client node.NodeClient) error {

	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{"Load Wallet", "Create Wallet", "Load Wallet From Seed", "Exit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		// Prompts should always succeed.
		panic(err)
	}

	var resp *node.LoadResponse
	switch result {
	case "Load Wallet":
		resp, err = loadWallet(client)
	case "Create Wallet":
		resp, err = createWallet(client)
	case "Load Wallet From Seed":
		resp, err = loadFromSeed(client)
	case "Exit":
		os.Exit(0)
	}

	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, string(resp.Key.PublicKey))
	return nil
}

// WalletMenu opens the prompt for doing wallet operations.
func WalletMenu(client node.NodeClient) error {
	for {
		// Get sync progress first and print it
		resp, err := client.GetSyncProgress(context.Background(), &node.EmptyRequest{})
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Sync progress: %.2f%", resp.Progress)

		prompt := promptui.Select{
			Label: "Select action",
			Items: []string{"Transfer DUSK", "Stake DUSK", "Bid DUSK", "Show Balance", "Show Address", "Show Transaction History", "Automate Consensus Participation", "Exit"},
			Size:  8,
		}

		_, result, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		var res string
		switch result {
		case "Transfer DUSK":
			resp, err := transferDusk(client)
			if err != nil {
				return err
			}

			res = "Tx hash: " + hex.EncodeToString(resp.Hash)
		case "Stake DUSK":
			resp, err := stakeDusk(client)
			if err != nil {
				return err
			}

			res = "Tx hash: " + hex.EncodeToString(resp.Hash)
		case "Bid DUSK":
			resp, err := bidDusk(client)
			if err != nil {
				return err
			}

			res = "Tx hash: " + hex.EncodeToString(resp.Hash)
		case "Show Balance":
			resp, err := client.GetBalance(context.Background(), &node.EmptyRequest{})
			if err != nil {
				return err
			}

			res = fmt.Sprintf("Unlocked balance: %.8f\nLocked balance: %.8f\n", float64(resp.UnlockedBalance)/float64(wallet.DUSK), float64(resp.LockedBalance)/float64(wallet.DUSK))
		case "Show Address":
			resp, err := client.GetAddress(context.Background(), &node.EmptyRequest{})
			if err != nil {
				return err
			}

			res = "Address: " + string(resp.Key.PublicKey)
		case "Show Transaction History":
			resp, err := client.GetTxHistory(context.Background(), &node.EmptyRequest{})
			if err != nil {
				return err
			}

			s := formatRecords(resp)
			res = s.String()
		case "Automate Consensus Participation":
			resp, err := client.AutomateConsensusTxs(context.Background(), &node.EmptyRequest{})
			if err != nil {
				return err
			}

			res = resp.Response
		case "Exit":
			os.Exit(0)
		}

		fmt.Fprintln(os.Stdout, res)
	}
}

func formatRecords(resp *node.TxHistoryResponse) strings.Builder {
	s := strings.Builder{}
	for _, record := range resp.Records {
		if record.Direction == node.Direction_IN {
			s.WriteString("IN / ")
		} else {
			s.WriteString("OUT / ")
		}
		// Height
		s.WriteString(strconv.FormatUint(record.Height, 10) + " / ")
		// Time
		s.WriteString(time.Unix(record.Timestamp, 0).Format(time.UnixDate) + " / ")
		// Amount
		s.WriteString(fmt.Sprintf("%.8f DUSK", float64(record.Amount)/float64(wallet.DUSK)) + " / ")
		// Unlock height
		s.WriteString("Unlocks at " + strconv.FormatUint(record.UnlockHeight, 10) + " / ")

		s.WriteString("\n")
	}

	return s
}

package main

import "github.com/dusk-network/dusk-wallet-cli/prompt"

func main() {
	for {
		// First, establish that the node is running by asking it about
		// it's wallet state.
		loaded, err := client.IsWalletLoaded()
		if err != nil {
			// TODO: implement checking on intervals up to a limit
			continue
		}

		// If we have no wallet loaded, we open the menu to load or
		// create one.
		if !loaded {
			if err := prompt.LoadMenu(); err != nil {
				// If we get an error from `LoadMenu`, it means we lost
				// our connection to the node. We will restart the loop
				// to attempt to regain our connection.
				// TODO: log
				continue
			}
		}

		// Once loaded, we open the menu for wallet operations.
		// The only reason this function returns is when we have lost
		// connection to the node. Hence, we have no return value
		// to check.
		prompt.WalletMenu()
	}
}

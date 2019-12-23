# dusk-wallet-cli

`dusk-wallet-cli` is a seperate application which connects to a DUSK node, in order to perform wallet and node management operations.

## Building

To build the wallet application, enter the following commands into the terminal:

```bash
go get github.com/dusk-network/dusk-wallet-cli
cd $GOPATH/src/github.com/dusk-network/dusk-wallet-cli
go build
```

Make sure you move the resulting `dusk-wallet-cli` executable into the same folder as your `dusk.toml` file. This is how the application figures out how to talk to your node.

## Usage

The CLI wallet can be started with the command:

```bash
./dusk-wallet-cli
```

Depending on whether a wallet is loaded or not on the running node, it will present you with a menu of your current options. You can navigate with the arrow keys, and select with enter.

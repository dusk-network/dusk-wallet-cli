package menu

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Start() error {
	// Looping until user chooses "Exit"
	for {
		prompt := promptui.Select{
			Label: "Select action",
			Items: []string{"Hello, world!", "Exit"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			return err
		}

		switch result {
		case "Hello, world!":
			fmt.Println("Hello, world!")
		case "Exit":
			return nil
		}
	}
}

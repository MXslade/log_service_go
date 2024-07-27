package admin_cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
		text = strings.TrimSpace(text)
		switch text {
		case "1":
			createAdmin()
		case "2":
			removeAdmin()
		case "0":
			return
		default:
			fmt.Println("Not valid option")
		}

	}
}

func printMainMenu() {
	fmt.Println("Choose one of the following: ")
	fmt.Println("1. Create Admin")
	fmt.Println("2. Remove Admin")
	fmt.Println("0. Exit")
	fmt.Print("Your choice: ")
}

func createAdmin() {
}

func removeAdmin() {
}

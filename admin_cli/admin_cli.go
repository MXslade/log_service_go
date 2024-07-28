package admin_cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/MXslade/log_service_go/db/repo/admin_repo"
)

type AdminCli struct {
	reader    *bufio.Reader
	adminRepo admin_repo.AdminRepo
}

func New() *AdminCli {
	reader := bufio.NewReader(os.Stdin)
	adminRepo := admin_repo.New()
	return &AdminCli{reader, adminRepo}
}

func (a *AdminCli) Start() {
	for {
		printMainMenu()
		text, err := a.reader.ReadString('\n')
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
		text = strings.TrimSpace(text)
		switch text {
		case "1":
			a.createAdmin()
		case "2":
			a.removeAdmin()
		case "3":
			a.showAllAdmins()
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
	fmt.Println("3. Show all Admins")
	fmt.Println("0. Exit")
	fmt.Print("Your choice: ")
}

func (a *AdminCli) createAdmin() {
	fmt.Println()
	fmt.Print("Username: ")
	username, err := a.reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Print("Password: ")
	password, err := a.reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	fmt.Printf("Username: %v, Password: %v\n", username, password)
}

func (a *AdminCli) removeAdmin() {
}

func (a *AdminCli) showAllAdmins() {
	admins, err := a.adminRepo.GetAll(context.Background())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

    if len(admins) == 0 {
        fmt.Println("You don't have any admins.\n")
    }

	for idx, admin := range admins {
		fmt.Printf("%v. ID: %v, Name: %v\n", idx, admin.ID, admin.Name)
	}
	fmt.Println()
}

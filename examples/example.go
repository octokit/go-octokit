package main

import (
	"fmt"
	"github.com/octokit/go-octokit/octokit"
)

func main() {
	client := octokit.NewClient()
	usersLink := &octokit.AllUsersHyperlink

	fmt.Println("Printing GitHub users for the first 5 pages")
	for i := 0; i < 5; i++ {
		if usersLink == nil {
			return
		}

		usersService, err := client.Users(usersLink, nil)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		users, result := usersService.GetAll()
		if result.HasError() {
			fmt.Println(result)
			return
		}

		for _, user := range users {
			fmt.Printf("%v - %s\n", user.ID, user.Login)
		}

		usersLink = result.NextPage
	}
}

package main

import (
	"fmt"
	"github.com/octokit/go-octokit"
)

func main() {
	client := octokit.NewClient()
	users, err := client.Users(&octokit.UsersHyperlink, octokit.M{"user": "jingweno"})
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	user, result := users.Get()
	if result.HasError() {
		fmt.Println(result)
		return
	}

	fmt.Println(user.Login)
}

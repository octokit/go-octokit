package main

import (
	"fmt"
	"github.com/octokit/go-octokit/octokit"
)

func main() {
	client := octokit.NewClient(nil)
	userURL := &octokit.UserURL

	fmt.Println("Printing GitHub users for the first 5 pages")
	for i := 0; i < 5; i++ {
		if userURL == nil {
			return
		}

		url, err := userURL.Expand(nil)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		users, result := client.Users(url).All()
		if result.HasError() {
			fmt.Println(result)
			return
		}

		for _, user := range users {
			fmt.Printf("%v - %s\n", user.ID, user.Login)
		}

		userURL = result.NextPage
	}
}

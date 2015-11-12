package main

import (
	"fmt"
	"os"

	"../octokit"
)

func all(client *octokit.Client) {
	fmt.Println("Getting all labels for octokit/go-octokit:")

	labels, result := client.Labels().All(nil, octokit.M{"owner": "octokit", "repo": "go-octokit"})

	if result.HasError() {
		fmt.Println("ERROR:", result)
		os.Exit(1)
	}

	for i, label := range labels {
		fmt.Printf("\t%d. %s", i, label.Name)
		fmt.Println()
	}
}

func main() {
	client := octokit.NewClient(nil)

	all(client)
}

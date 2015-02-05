package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	//"github.com/octokit/go-octokit/octokit"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var githubURL = "https://api.github.com/"

var sourceFolder = filepath.FromSlash(os.Getenv("GOPATH") +
	"/src/github.com/octokit/go-octokit/octokit/")

var URLDeclarationMatcher = regexp.MustCompile(
	`[A-Za-z]+URL = Hyperlink\(\"([\/a-z0-9_{}]+)\"\)`)

func listSourceFiles(dirname string) []string {
	var result []string
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") &&
			!strings.HasSuffix(f.Name(), "_test.go") {
			result = append(result, f.Name())
		}
	}
	return result
}

func extractURLFromSourceFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := URLDeclarationMatcher.FindStringSubmatch(scanner.Text())
		if len(res) > 1 {
			return res[1]
		}
	}
	return ""
}

func main() {

	resp, err := http.Get(githubURL)
	if err != nil {
		fmt.Printf("Error opening %s\n", githubURL)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error fetching data from %s\n", githubURL)
		panic(err)
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	existingValues := make(map[string]bool)
	extraValues := make(map[string]bool)

	for _, v := range dat {
		v := v.(string)
		existingValues[v[len(githubURL):]] = false
	}

	for _, f := range listSourceFiles(sourceFolder) {
		url := extractURLFromSourceFile(sourceFolder +
			string(os.PathSeparator) + f)
		if url != "" {
			if _, ok := existingValues[url]; ok {
				existingValues[url] = true
			} else {
				extraValues[url] = true
			}
		}
	}

	fmt.Println("\nExisting APIs:")

	for k, v := range existingValues {
		if v {
			fmt.Println(k)
		}
	}

	fmt.Println("\nUnaccounted APIs:")

	for k, _ := range extraValues {
		fmt.Println(k)
	}

	fmt.Println("\nMissing APIs:")

	for k, v := range existingValues {
		if !v {
			fmt.Println(k)
		}
	}
}

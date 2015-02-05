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
	`[A-Za-z]+URL\s*=\s*Hyperlink\(\"([\/a-z0-9_{}]+)\"\)`)

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

func extractURLsFromSourceFile(filename string) (results []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	varBlockPresent := false
	for scanner.Scan() {
		text := scanner.Text()
		if text == "var (" {
			varBlockPresent = true
		}
		res := URLDeclarationMatcher.FindStringSubmatch(text)
		if len(res) > 1 {

			results = append(results, res[1])
			if !varBlockPresent {
				return
			}
		}
		if varBlockPresent && text == ")" {
			return
		}
	}
	return
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
		existingValues[v[strings.Index(v, ".com/")+5:]] = false
	}

	var apisFoundCount = 0

	for _, f := range listSourceFiles(sourceFolder) {
		urls := extractURLsFromSourceFile(sourceFolder +
			string(os.PathSeparator) + f)
		for _, url := range urls {
			if len(url) > 1 && url[0] == '/' {
				url = url[1:]
			}
			if url != "" {
				apisFoundCount += 1
				if _, ok := existingValues[url]; ok {
					existingValues[url] = true
				} else {
					extraValues[url] = true
				}
			}
		}
	}

	fmt.Printf("%d APIs found on %s.\n", len(existingValues), githubURL)
	fmt.Printf("%d APIs implementations found.\n", apisFoundCount)

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

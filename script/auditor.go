package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// The root of go-octokit by traversing the go directory structure
var octokitRoot = filepath.FromSlash(os.Getenv("GOPATH") +
	"/src/github.com/octokit/go-octokit/")

// The octokit package folder within the octokit source folder
var sourceFolder = filepath.FromSlash(octokitRoot + "octokit/")

// The path to the TODO file that store the list of implemented and
// unimplemented APIs
var todoFile = filepath.FromSlash(octokitRoot + "TODO.md")

// The pattern to look for that matches the documentation URLs
var URLDeclarationMatcher = regexp.MustCompile(
	`//[\t ]+(http[s]?://[\/a-z0-9\-_{}\.:#]+)`)

// List all non-test Go source files
func listSourceFiles(dirname string) []string {
	var result []string
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") &&
			!strings.HasSuffix(f.Name(), "_test.go") {
			result = append(result, path.Join(dirname, f.Name()))
		}
	}
	return result
}

// List all documentation URLs from the specified source file
func extractURLsFromSourceFile(filename string) (results []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		res := URLDeclarationMatcher.FindStringSubmatch(text)
		if len(res) > 1 {
			results = append(results, res[1])
		}
	}
	return
}

func main() {
	file, err := os.Open(todoFile)
	if err != nil {
		fmt.Printf("Error opening TODO file: %s\n", todoFile)
		panic(err)
	}

	var todoUrls []string
	todoMap := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		url := line[6:]
		todoUrls = append(todoUrls, url)
		if line[3] == 'x' {
			todoMap[url] = true
		} else {
			todoMap[url] = false
		}
	}

	file.Close()

	for _, sourceFile := range listSourceFiles(sourceFolder) {
		for _, url := range extractURLsFromSourceFile(sourceFile) {
			todoMap[url] = true
		}
	}

	var outString string

	for _, url := range todoUrls {
		status := " "
		if todoMap[url] {
			status = "x"
		}

		outString += fmt.Sprintf("- [%s] %s\n", status, url)
	}

	err = ioutil.WriteFile(todoFile, []byte(outString), 0644)
	if err != nil {
		fmt.Printf("Error opening TODO file: %s\n", todoFile)
		panic(err)
	}
}

package main

import (
	"bufio"
	//"bytes"
	//"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var octokitRoot = filepath.FromSlash(os.Getenv("GOPATH") +
	"/src/github.com/octokit/go-octokit/")

var sourceFolder = filepath.FromSlash(octokitRoot + "octokit/")

var todoFile = filepath.FromSlash(octokitRoot + "TODO.md")

var URLDeclarationMatcher = regexp.MustCompile(
	`// See ([\/a-z0-9_{}\.:]+)`)

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
	for scanner.Scan() {
		text := scanner.Text()
		res := URLDeclarationMatcher.FindStringSubmatch(text)
		if len(res) > 1 {
			results = append(results, res[1])
		}
	}
	return
}

func findAllExistingApiUrls() (urls []string) {
	for _, f := range listSourceFiles(sourceFolder) {
		sourceUrls := extractURLsFromSourceFile(sourceFolder +
			string(os.PathSeparator) + f)
		urls = append(urls, sourceUrls...)
	}
	return urls
}

func main() {
	/*file, err := os.Open(todoFile)
	if err != nil {
		fmt.Printf("Error opening TODO file: %s\n", todoFile)
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		url := line[4:]

	}

	var dat map[string]interface{}

	existingValues := make(map[string]bool)
	extraValues := make(map[string]bool)

	for _, v := range dat {
		v := v.(string)
		existingValues[v[strings.Index(v, ".com/")+5:]] = false
	}

	var apisFoundCount = 0

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
	*/
}

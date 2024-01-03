package fts

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var ignore_dirs = []string{".git", ".tox", "node_modules", "target", ".jar", ".idea", ".vscode"}
var ignore_extensions = []string{"svg", "png", "jpg", "pdf", "jar", "idea", "xsd", ".gitignore"}

func IsExist(str, filepath string) bool {
	fmt.Println("searching in file: " + filepath)
	//check if file can be read

	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("error reading file: " + filepath)
		return false
	}

	isExist, err := regexp.Match(str, b)
	if err != nil {
		fmt.Println("error matching string in file" + filepath)
		return false
	}
	return isExist
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func containsExt(s []string, str string) bool {
	split := strings.Split(str, ".")
	str = split[len(split)-1]
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func FTS(searchString string, searchDirectory string, ignoreExt []string, ignoreFolders []string, fileName string, extensionType string) []string {
	var files []string
	var result []string
	if len(searchString) <= 0 {
		fmt.Println("Please provide a search string")
		return nil
	}
	if searchDirectory == "." {
		path, _ := os.Getwd()
		searchDirectory = path
	}

	if len(ignoreExt) > 0 {
		ignore_extensions = append(ignore_extensions, ignoreExt...)
	}
	if len(ignoreFolders) > 0 {
		ignore_dirs = append(ignore_dirs, ignoreFolders...)
	}
	status_msg := fmt.Sprintf("Searching for text '%s' in directory '%s'.", searchString, searchDirectory)
	fmt.Println(status_msg)

	err := filepath.Walk(searchDirectory, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() && contains(ignore_dirs, info.Name()) {
			return filepath.SkipDir
		}
		if containsExt(ignore_extensions, info.Name()) {
			fmt.Println("ignoring", info.Name())
			return nil
		}
		if len(fileName) > 0 {
			if !info.IsDir() && info.Name() == fileName {
				files = append(files, path)
			}
		} else if len(extensionType) > 0 {
			if !info.IsDir() && filepath.Ext(path) == extensionType {
				files = append(files, path)

			}
		} else if len(extensionType) == 0 && len(fileName) > 0 {
			if !info.IsDir() && info.Name() == fileName {
				files = append(files, path)

			}
		} else if len(extensionType) == 0 && len(fileName) == 0 {

			if !info.IsDir() {
				files = append(files, path)

			}
		}
		return nil

	})
	fmt.Println(files)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			exists := IsExist(searchString, file)
			if exists {
				result = append(result, file)
			}
		}(file)
	}
	wg.Wait()
	return result
}

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var (
	Version         = "1.0.0"
	Use             = "fts"
	SearchString    string
	FileName        string
	ExtensionType   string
	IgnoreFolders   []string
	IgnoreExt       []string
	SearchDirectory string
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

func fts() {
	var files []string

	if len(SearchString) <= 0 {
		fmt.Println("Please provide a search string")
	}
	if SearchDirectory == "." {
		path, _ := os.Getwd()
		SearchDirectory = path
	}

	if len(IgnoreExt) > 0 {
		ignore_extensions = append(ignore_extensions, IgnoreExt...)
	}
	if len(IgnoreFolders) > 0 {
		ignore_dirs = append(ignore_dirs, IgnoreFolders...)
	}
	status_msg := fmt.Sprintf("Searching for text '%s' in directory '%s'.", SearchString, SearchDirectory)
	fmt.Println(status_msg)

	err := filepath.Walk(SearchDirectory, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() && contains(ignore_dirs, info.Name()) {
			return filepath.SkipDir
		}
		if containsExt(ignore_extensions, info.Name()) {
			fmt.Println("ignoring", info.Name())
			return nil
		}
		if len(FileName) > 0 {
			if !info.IsDir() && info.Name() == FileName {
				files = append(files, path)
			}
		} else if len(ExtensionType) > 0 {
			if !info.IsDir() && filepath.Ext(path) == ExtensionType {
				files = append(files, path)

			}
		} else if len(ExtensionType) == 0 && len(FileName) > 0 {
			if !info.IsDir() && info.Name() == FileName {
				files = append(files, path)

			}
		} else if len(ExtensionType) == 0 && len(FileName) == 0 {

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
			exists := IsExist(SearchString, file)
			if exists {
				fmt.Println(file)
			}
		}(file)
	}
	wg.Wait()
}

var rootCmd = &cobra.Command{
	Use:           Use,
	Version:       Version,
	SilenceErrors: true,
	Short:         "Fast Text Search",
	Long:          "Extremely fast and concurrent text search in Go",
	Run: func(cmd *cobra.Command, args []string) {
		fts()
	},
	SilenceUsage: true,
}

func Execute() {
	rootCmd.Use = os.Args[0]
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&SearchDirectory, "dir", "d", ".", "The directory to search in")
	rootCmd.Flags().StringVarP(&SearchString, "text", "t", "", "The string to search for")
	rootCmd.Flags().StringVarP(&FileName, "file", "f", "", "The file to seach in")
	rootCmd.Flags().StringVarP(&ExtensionType, "ext", "e", "", "The extension of files to search in")
	rootCmd.Flags().StringArrayVarP(&IgnoreFolders, "folderignore", "i", []string{}, "The folders names to ignore during search")
	rootCmd.Flags().StringArrayVarP(&IgnoreExt, "extignore", "x", []string{}, "The extension type to ignore during search")

}

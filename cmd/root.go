package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version         = "1.0.0"
	Use             = "fts"
	SearchString    string
	FileName        string
	ExtensionType   string
	IgnoreFiles     []string
	IgnoreExt       []string
	SearchDirectory string
)

func fts() {
	var files []string

	if len(SearchString) <= 0 {
		fmt.Println("Please provide a search string")
	}
	if SearchDirectory == "." {
		path, _ := os.Getwd()
		SearchDirectory = path
	}
	err := filepath.Walk(SearchDirectory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && contains([]string{".git", ".tox", "node_modules", "target", ".jar", ".idea", ".vscode", "lib64"}, info.Name()) {
			return filepath.SkipDir
		}
	}

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
	rootCmd.Flags().StringArrayVarP(&IgnoreFiles, "fileignore", "i", []string{}, "The file names to ignore during search")
	rootCmd.Flags().StringArrayVarP(&IgnoreExt, "extignore", "x", []string{}, "The extension type to ignore during search")

}

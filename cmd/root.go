package cmd

import (
	"fast-test-search/fts"
	"os"

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

var rootCmd = &cobra.Command{
	Use:           Use,
	Version:       Version,
	SilenceErrors: true,
	Short:         "Fast Text Search",
	Long:          "Extremely fast and concurrent text search in Go",
	Run: func(cmd *cobra.Command, args []string) {
		fts.FTS(SearchString, SearchDirectory, IgnoreExt, IgnoreFolders, FileName, ExtensionType)
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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version       = "1.0.0"
	Use           = "fts"
	SearchString  string
	FileName      string
	ExtensionType string
	IgnoreFiles   []string
	IgnoreExt     []string
)

func fts() (err error) {
	fmt.Println("hello")
	return nil
}

var rootCmd = &cobra.Command{
	Use:           Use,
	Version:       Version,
	SilenceErrors: true,
	Short:         "Fast Text Search",
	Long:          "Extremely fast and concurrent text search in Go",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fts()
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
	rootCmd.Flags().StringVarP(&SearchString, "text", "t", "", "The string to search for")
	rootCmd.Flags().StringVarP(&FileName, "file", "f", "", "The file to seach in")
	rootCmd.Flags().StringVarP(&ExtensionType, "ext", "e", "", "The extension of files to search in")
	rootCmd.Flags().StringArrayVarP(&IgnoreFiles, "fileignore", "fi", []string{}, "The file names to ignore during search")
	rootCmd.Flags().StringArrayVarP(&IgnoreExt, "extignore", "ei", []string{}, "The extension type to ignore during search")

}

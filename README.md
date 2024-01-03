#  ☄️ fast-text-search

A go package to do extremely fast concurrent text search across files and folders.

## Description

Search for text recursively taking advantage of concurrency via GO routines. Import the project has a package and plugin the search functionality or use it as a command line application. Features options to ignore filetypes and directories during search. 


## Package Usage

Use the file search logic and plug it in to existing codebases.

```golang
import "github.com/raiyanyahya/fast-text-search/fts"

var (
	SearchString    string      // String to search for [mandatory]
	FileName        string      // Explicitly mention the filename to look in.
	ExtensionType   string      // Only search files with these extensions.
	IgnoreFolders   []string    // Exclude these directories while searching.
	IgnoreExt       []string    // Do not open files with these extensions while searching.
	SearchDirectory string      // The starting search directory.
)

func main(){
    
    hits := fts.FTS(searchString, SearchDirectory, IgnoreExt, IgnoreFolders, 
            FileName, ExtensionType)

}

```

## Command Line Usage

You could also use this directly from the command line.

```bash

Extremely fast and concurrent text search in Golang

Usage:
  fts [flags]

Flags:
  -d, --dir string                 The directory to search in (default ".")
  -e, --ext string                 The extension of files to search in
  -x, --extignore stringArray      The extension type to ignore during search
  -f, --file string                The file to seach in
  -i, --folderignore stringArray   The folders names to ignore during search
  -h, --help                       help for fts
  -t, --text string                The string to search for
  -v, --version                    version for fts
```

## License
FTS is a free for open-source projects, however, if you are using the library for business and commercial projects you could choose to buy me a coffee.

## Contributing

Contributions are always welcome!
See contributing.md for ways to get started. Please adhere to this project's code of conduct.

## Contact
Contact me through email.
/*
©2023 J.F.Gratton (jean-francois@famillegratton.net)
*/
package executor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
	VERSION HISTORY

version			date			comments
-------			----			--------
1.01.00			2023.04.27		verbosity and cobra framework addition (much overkill there)
1.000			2023.04.26		release
0.100			2023.04.26		initial working version
*/

func Rename(verbose bool, directory string) {
	// Before we proceed, a hard stop if we're working on $HOME !
	currentdir, _ := os.Getwd()
	if os.Getenv("HOME") == currentdir && (directory == "." || directory == os.Getenv("HOME")) {
		fmt.Printf("I'm cowardly that way, %s I'd rename files in %s !\n", Red("no way"), Yellow(currentdir))
		os.Exit(-4)
	}
	//if directory == "." && os
	dirfp, err := os.Open(directory)
	if err != nil {
		fmt.Printf("Error while opening the %s directory\n", directory)
		fmt.Println(err)
		os.Exit(-1)
	}
	files, err := dirfp.Readdir(0)
	if err != nil {
		fmt.Println("Error while reading files in directory: ", err)
		os.Exit(-2)
	}

	for _, file := range files {
		if !file.IsDir() {
			newname := strings.ToLower(file.Name())
			err := os.Rename(filepath.Join(directory, file.Name()), filepath.Join(directory, newname))
			if err != nil {
				fmt.Printf("%s %s:\n", Red("Cannot rename "), file.Name())
				fmt.Println(err)
			} else {
				if verbose {
					fmt.Printf("Rename %s -> %s : %s\n", file.Name(), newname, Green("success"))
				}
			}
		}
	}
}

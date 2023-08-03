/*
©2023 J.F.Gratton (jean-francois@famillegratton.net)
*/
package executor

import (
	"fmt"
	"os"
)

/*
	VERSION HISTORY

version			date			comments
-------			----			--------
0.100			2023.05.11		initial working version
*/

func Rename(verbose bool, normalize bool, uppercase bool, lowercase bool,
	stripPatterns []string, targets []string) {
	if !normalize && !uppercase && !lowercase && len(stripPatterns) == 0 {
		fmt.Println("You need to specify at least one of the following flags: -l, -u, -n or -s")
		os.Exit(0)
	}
	// We fetch the current directory, in case we need to return to it eventually
	originalDir, _ := os.Getwd()

	// Looping through all dirs in argum
	for _, directory := range targets {
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
				currentName := file.Name()
				if lowercase != uppercase {
					_, err := changeCase(lowercase, directory, currentName)
					if err != nil {
						fmt.Printf("%s %s:\n", Red("Cannot rename"), currentName)
						fmt.Println(err)
					}
				}
				if len(stripPatterns) > 0 {
					if err = stripFromName(directory, currentName, stripPatterns); err != nil {
						fmt.Printf("%s %s:\n", Red("Unable to strip"), currentName)
						fmt.Println(err)
					}
				}
				if normalize {
					if err := normalizeFilename(directory, currentName); err != nil {
						fmt.Printf("%s %s:\n", Red("Unable to normalize"), currentName)
						fmt.Println(err)
					}
				}
				if verbose {
					fmt.Printf("%s %s %s", White(currentName), Green(" -> "), White(file.Name()))
				}
			}
		}
	}
	os.Chdir(originalDir)
}

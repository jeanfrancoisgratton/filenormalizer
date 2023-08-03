// filenormalizer
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/executor/helpers.go
// Original time: 2023/05/11 16:44

package executor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Filename is renamed in all-lowercase or all-uppercase
func changeCase(lower bool, directory string, oldname string) (string, error) {
	newname := ""
	var err error

	if lower {
		newname = strings.ToLower(oldname)
		err = os.Rename(filepath.Join(directory, oldname), filepath.Join(directory, newname))
	} else {
		newname = strings.ToUpper(oldname)
		err = os.Rename(filepath.Join(directory, oldname), filepath.Join(directory, newname))
	}

	if err != nil {
		fmt.Printf("%s : %s\n", Red("Error"), err)
	}

	return newname, err
}

// This is where we remove parts of filenames:
// myfile1-SAMPLE.txt, myfile2-SAMPLE.txt become myfile1.txt, myfile2.txt
func stripFromName(dir string, name string, patterns []string) error {
	oldPath := filepath.Join(dir, name)
	for _, pattern := range patterns {
		name = strings.Replace(name, pattern, "", -1)
	}
	if err := os.Rename(oldPath, filepath.Join(dir, name)); err != nil {
		return err
	}
	return nil
}

// This is where we "normalize filenames :
// Patterns such as "dots instead of whitespaces" are whitespaced, etc
func normalizeFilename(dir string, name string) error {
	// Get the filename without the path
	filename := filepath.Base(name)

	// Split the filename into words
	words := strings.Split(filename, "_")

	// If there are no underscores, split the filename into words using dots
	if len(words) == 1 {
		words = strings.Split(filename, ".")
	}

	// Join the words together with spaces
	newname := strings.Join(words, " ")

	// Add the file extension back to the filename
	ext := filepath.Ext(filename)
	newname = strings.TrimSuffix(newname, ext) + ext

	if err := os.Rename(filepath.Join(dir, name), filepath.Join(dir, newname)); err != nil {
		return err
	}
	return nil
}

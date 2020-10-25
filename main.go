package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FileMove .
type FileMove struct {
	OldPath string
	NewPath string
}

func confirm() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hit Enter to Continue (cancel otherwise) [Enter]: ")
	text, err := reader.ReadString('\n')
	check(err)
	text = strings.TrimSpace(text)

	if len(text) == 0 {
		return true
	}

	return false
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// given the current directory,
	// add a prefix to all files in all subdirectoires
	// move the files to the current directory
	// delete the subdirectories
	cwd, err := os.Getwd()
	check(err)

	dirs, err := ioutil.ReadDir(cwd)
	check(err)

	moves := []FileMove{}

	for _, file := range dirs {
		if !file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		subdir := file.Name()
		// fmt.Printf("%v\n", subdir)

		path := filepath.Join(cwd, subdir)
		subfiles, err := ioutil.ReadDir(path)
		check(err)

		for _, subfile := range subfiles {
			if subfile.IsDir() {
				continue
			}

			// fmt.Printf("  |- %v\n", subfile.Name())
			newfile := subdir + "." + subfile.Name()

			oldpath := filepath.Join(path, subfile.Name())
			newpath := filepath.Join(filepath.Dir(path), newfile)

			move := FileMove{oldpath, newpath}
			moves = append(moves, move)

			// fmt.Printf("    Move %v\n", move)
		}

	}

	maxlen := 0
	for _, move := range moves {
		relold, _ := filepath.Rel(cwd, move.OldPath)
		length := len(relold)

		if length > maxlen {
			maxlen = length
		}
	}

	fmtstr := "  %-" + strconv.Itoa(maxlen) + "v => %v\n"

	fmt.Printf("Pending file Moves: \n\n")

	for _, move := range moves {
		// fmt.Printf("%v\n", move)
		relold, _ := filepath.Rel(cwd, move.OldPath)
		relnew, _ := filepath.Rel(cwd, move.NewPath)

		fmt.Printf(fmtstr, relold, relnew)
	}

	fmt.Printf("\n")

	if !confirm() {
		fmt.Printf("Canceled!\n")
		os.Exit(0)
	}

	for _, move := range moves {
		os.Rename(move.OldPath, move.NewPath)

		olddir := filepath.Dir(move.OldPath)
		os.Remove(olddir)
	}

}

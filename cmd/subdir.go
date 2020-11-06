package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidwashere/damv/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(subdirCmd)
}

var subdirCmd = &cobra.Command{
	Use:   "subdir",
	Short: "Moves files in subdirs to current dir while prefixing subdir to filename",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		subdirs()
	},
}

func subdirs() {
	// given the current directory,
	// add a prefix to all files in all subdirectoires
	// move the files to the current directory
	// delete the subdirectories
	cwd, err := os.Getwd()
	util.ErrCheck(err)

	dirs, err := ioutil.ReadDir(cwd)
	util.ErrCheck(err)

	moves := []*util.FileMove{}

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
		util.ErrCheck(err)

		for _, subfile := range subfiles {
			if subfile.IsDir() {
				continue
			}

			// fmt.Printf("  |- %v\n", subfile.Name())
			newfile := subdir + "." + subfile.Name()

			oldpath := filepath.Join(path, subfile.Name())
			newpath := filepath.Join(filepath.Dir(path), newfile)

			move := util.FileMove{oldpath, newpath}
			moves = append(moves, &move)

			// fmt.Printf("    Move %v\n", move)
		}

	}

	util.PrintMoves(cwd, moves)

	if !util.PromptConfirm() {
		fmt.Printf("Canceled!\n")
		os.Exit(0)
	}

	for _, move := range moves {
		os.Rename(move.OldPath, move.NewPath)

		olddir := filepath.Dir(move.OldPath)
		os.Remove(olddir)
	}

}

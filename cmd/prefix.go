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
	rootCmd.AddCommand(prefixCmd)
}

var (
	prefixReplaceStr string
)

var prefixCmd = &cobra.Command{
	Use:   "prefix PREFIX",
	Short: "Adds prefix to files in current directory",
	// Long:  ``,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Prefix: %v, Replacing: [%v]\n", args[0], prefixFlagReplaceStr)
		prefix(args)
	},
}

func init() {
	prefixCmd.Flags().StringVarP(&prefixReplaceStr, "replace", "r", "", "Replaces this existing prefix if found")
}

func prefix(args []string) {
	cwd, err := os.Getwd()
	util.ErrCheck(err)

	dirs, err := ioutil.ReadDir(cwd)
	util.ErrCheck(err)

	moves := []*util.FileMove{}

	newPrefix := args[0]

	for _, file := range dirs {
		if file.IsDir() {
			continue
		}

		newFileName := file.Name()
		if len(prefixReplaceStr) > 0 && strings.HasPrefix(file.Name(), prefixReplaceStr) {
			// Removes the existing prefix that need replacement
			newFileName = strings.Replace(file.Name(), prefixReplaceStr, "", 1)
		}

		newFileName = newPrefix + newFileName

		oldpath := filepath.Join(cwd, file.Name())
		newPath := filepath.Join(cwd, newFileName)

		move := util.FileMove{
			OldPath: oldpath,
			NewPath: newPath,
		}
		moves = append(moves, &move)
	}

	util.PrintMoves(cwd, moves)

	if !util.PromptConfirm() {
		fmt.Printf("Canceled!\n")
		os.Exit(0)
	}

	for _, move := range moves {
		os.Rename(move.OldPath, move.NewPath)
	}

}

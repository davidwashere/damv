package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/davidwashere/damv/util"
	"github.com/spf13/cobra"
)

var (
	startSequenceAt int
)

func init() {
	rootCmd.AddCommand(sequenceCmd)
}

var sequenceCmd = &cobra.Command{
	Use:   "seq NAME",
	Short: "Renames files with a base name + sequence number",
	Long: `Seq will rename files in the current directory to NAME plus an incrementing 
number suffix

Files will first be sorted alphabetically
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sequence(args)
	},
}

func init() {
	sequenceCmd.Flags().IntVarP(&startSequenceAt, "start", "s", 0, "Start incrementing from this number")
}

func sequence(args []string) {
	cwd, err := os.Getwd()
	util.ErrCheck(err)

	dirs, err := ioutil.ReadDir(cwd)
	util.ErrCheck(err)

	moves := []*util.FileMove{}

	newName := args[0]

	for _, file := range dirs {
		if file.IsDir() {
			continue
		}

		oldpath := filepath.Join(cwd, file.Name())

		move := util.FileMove{
			OldPath: oldpath,
		}
		moves = append(moves, &move)
	}

	numMovesStr := strconv.Itoa(len(moves))
	pad := strconv.Itoa(len(numMovesStr))
	newNameTemplate := "%s%0" + pad + "d%s"

	for i, move := range moves {
		oldFileExt := filepath.Ext(move.OldPath)

		num := i + startSequenceAt

		newFileName := fmt.Sprintf(newNameTemplate, newName, num, oldFileExt)
		move.NewPath = filepath.Join(cwd, newFileName)
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

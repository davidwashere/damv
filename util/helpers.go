package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func PromptConfirm() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[Enter] to continue, anything else aborts: ")
	text, err := reader.ReadString('\n')
	ErrCheck(err)
	text = strings.TrimSpace(text)

	if len(text) == 0 {
		return true
	}

	return false
}

func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// FileMove .
type FileMove struct {
	OldPath string
	NewPath string
}

func PrintMoves(base string, moves []FileMove) {
	maxlen := 0
	for _, move := range moves {
		relold, _ := filepath.Rel(base, move.OldPath)
		length := len(relold)

		if length > maxlen {
			maxlen = length
		}
	}

	fmtstr := "  %-" + strconv.Itoa(maxlen) + "v => %v\n"

	fmt.Printf("Pending moves: \n\n")

	for _, move := range moves {
		relold, _ := filepath.Rel(base, move.OldPath)
		relnew, _ := filepath.Rel(base, move.NewPath)

		fmt.Printf(fmtstr, relold, relnew)
	}

	fmt.Printf("\n")
}

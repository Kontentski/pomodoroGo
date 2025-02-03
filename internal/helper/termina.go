package helper

import (
	"os"

	"golang.org/x/term"
)

func HasEnoughSpace() (bool, int, int) {//checks the size of the terminal
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width < MinWidth || height < MinHeight {
		return false, width, height
	}
	return true, width, height
}

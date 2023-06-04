package error

import (
	"fmt"
	"os"
)

const (
	LexError = "LexError"
)

func Error(err_type string, message string, line int, linetext string) {
	fmt.Fprintf(os.Stderr, "\n"+err_type+": "+message+"\n")
	fmt.Fprintf(os.Stderr, fmt.Sprint(line)+" | "+linetext+"\n"+"    ")

	i := 1
	for i <= len(linetext) {
		fmt.Fprintf(os.Stderr, "~")
		i += 1
	}
	fmt.Fprintf(os.Stderr, " <- here\n\n")
}

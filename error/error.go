package error

import (
	"fmt"
	"os"
)

const (
	LexError = "LexError"
)

func Error(err_type string, message string, line int) {
	fmt.Fprintf(os.Stderr, "[at line %d] %s: %s", line, err_type, message+"\n")
}

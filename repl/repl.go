package repl

import (
	"bufio"
	"fir/lexer"
	"fmt"
	"os"
	"strings"
)

func Repl() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		l := lexer.New(text)
		for _, tok := range l.Scan_tokens() {
			fmt.Println(tok.String())
		}
	}
}

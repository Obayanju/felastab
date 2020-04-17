package search

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/obayanju/felastab/lexer"
	"github.com/obayanju/felastab/token"
)

func Start(in string, tokens *[]string) {
	scanner := bufio.NewScanner(strings.NewReader(in))
	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
			*tokens = append(*tokens, "Token: "+string(tok.Type)+" Literal: "+tok.Literal)
		}

	}
}

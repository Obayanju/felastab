package search

import (
	"io"
	"strings"

	"github.com/obayanju/felastab/lexer"
	"github.com/obayanju/felastab/token"
)

func Start(in string, tokens *[]string) error {
	r := strings.NewReader(in)
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err == io.EOF {
			return nil
		}
		l := lexer.New(string(b[0]))

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			*tokens = append(*tokens, "Token: "+string(tok.Type)+" Literal: "+tok.Literal)
		}

	}
}

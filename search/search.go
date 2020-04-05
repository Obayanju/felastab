package search

import (
	"fmt"
	"io"
	"strings"

	"github.com/obayanju/felastab/lexer"
	"github.com/obayanju/felastab/token"
)

func Start(in string) {
	r := strings.NewReader(in)
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		l := lexer.New(string(b[0]))

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}

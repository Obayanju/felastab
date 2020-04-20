package search

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/obayanju/felastab/lexer"
	"github.com/obayanju/felastab/token"
)

func Start(in string, tokens *[]token.Token, filePath string) {
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Split(bufio.ScanLines)
	lineNum := 1
	for scanner.Scan() {
		fmt.Println(lineNum)
		l := lexer.New(scanner.Text())
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if tok.Type == "IDENT" {
				tok.LineNumber = lineNum
				tok.FilePath = filePath
				s := fmt.Sprintf("%s", tok.FilePath+" Token -> "+string(tok.Type)+" Literal -> "+tok.Literal)
				detail := fmt.Sprintf("%d: %s", tok.LineNumber, s)
				*tokens = append(*tokens, tok)
				fmt.Println(detail)
			}
		}
		lineNum++

	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

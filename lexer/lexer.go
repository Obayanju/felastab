package lexer

// The Lexer takes in source code as input and outputs
// the tokens that represents the source code

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

package term

import (
	"errors"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <term>     ::= ATOM | NUM | VAR | <compound>
// <compound> ::= <functor> LPAR <args> RPAR
// <functor>  ::= ATOM
// <args>     ::= <term> | <term> COMMA <args>
//

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

type DAG struct { //implement
	root *Term //change maybe
}

//prev term for comma checks

func parserHelper(inCompound, prevTerm bool, itr int, str string, isValid *bool) {
	lex := newLexer(str)
	token, _ := lex.next()

	if (token.typ == tokenEOF) && (!inCompound) {
		t := true
		isValid = &t
	} else if (token.typ == tokenEOF) && (inCompound) { // compound never closed with parenthR

	}

	if token.typ == tokenNumber || token.typ == tokenVariable {
		parserHelper(inCompound, true, itr+1, str[itr:], isValid)
	} else if token.typ == tokenAtom {

	}

}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	panic("TODO: implement NewParser")
}

package sexpr

import (
	"errors"
	"fmt"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <sexpr>       ::= <atom> | <pars> | QUOTE <sexpr>
// <atom>        ::= NUMBER | SYMBOL
// <pars>        ::= LPAR <dotted_list> RPAR | LPAR <proper_list> RPAR
// <dotted_list> ::= <proper_list> <sexpr> DOT <sexpr>
// <proper_list> ::= <sexpr> <proper_list> | \epsilon
//
type Parser interface {
	Parse(string) (*SExpr, error)
}

func NewParser() Parser {
	return &ParserImpl{
		lex:          nil,
		peekTok:      nil,
		sexpr:        make(map[string]*SExpr),
		sexprID:      make(map[*SExpr]int),
		sexprCounter: 0,
	}
}

type ParserImpl struct {
	lex          *lexer
	peekTok      *token
	sexpr        map[string]*SExpr
	sexprID      map[*SExpr]int
	sexprCounter int
}

func (p *ParserImpl) nextToken() (*token, error) {
	if tok := p.peekTok; tok != nil {
		p.peekTok = nil
		return tok, nil
	}
	return p.lex.next()
}

func (p *ParserImpl) backToken(tok *token) {
	p.peekTok = tok
}

//Parse a sexpr
func (p *ParserImpl) Parse(input string) (*SExpr, error) {
	p.lex = newLexer(input)
	p.peekTok = nil

	//checks empty input
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	if tok.typ == tokenEOF { //empty case
		return nil, ErrParser
	}

	p.backToken(tok)
	//fmt.Printf("inside PARSE tok lit is %d", tok.num)
	sexpr, err := p.parseNextSexpr() //impl

	if err != nil {
		return nil, ErrParser
	}
	fmt.Printf("inside PARSE tok lit is %s", tok.literal)
	if tok, err := p.nextToken(); err != nil || tok.typ != tokenEOF {
		return nil, ErrParser
	}

	return sexpr, nil //giving nil value for sexpr ?
}

func (p *ParserImpl) parseNextSexpr() (*SExpr, error) { //implementing
	tok, err := p.nextToken()
	if err != nil {
		return nil, err
	}

	switch tok.typ {
	case tokenEOF:
		return nil, nil
	case tokenNumber:
		//fmt.Printf("inside PNS tok lit is %s", tok.literal)
		return p.mkSimpleSexpr(tokenNumber, tok), nil
	case tokenSymbol:
		return p.mkSimpleSexpr(tokenSymbol, tok), nil
	case tokenLpar:
		//next has to be a sexpr or epsilon
		nxt, err := p.nextToken()
		if err != nil { // error exists
			return nil, err // has to have something after
		}
		if nxt.typ == tokenEOF { //has to have something after
			return nil, ErrParser
		}

		//epsilon case ?
		if nxt.typ == tokenRpar { //case for its empty  and we use rpar token as our key for that
			return p.mkSimpleSexpr(tokenRpar, tok), nil //arbritrary tok pass
		}

		//not a r par and is a sexpr have to go back and get the sexpr
		//get first argument inside parenth through recursive call cannot be dot
		p.backToken(nxt)               //go back 1 because we are at the first arg ADDED
		arg, err := p.parseNextSexpr() // get first arg
		if err != nil {                // error exists
			return nil, err
		}

		elements := []*SExpr{arg} //use to make cons expr has car as first elem
		nxt, err = p.nextToken()
		if err != nil { // error exists
			return nil, err
		}

		//run args till r par show up
		for ; nxt.typ != tokenRpar && nxt.typ != tokenEOF; nxt, err = p.nextToken() {
			//go back 1 more ?
			if nxt.typ == tokenDot { //JUST ADDED
				peeker, err := p.nextToken() //check if sexpr
				if err != nil {
					return nil, err
				}
				if peeker.typ == tokenRpar || peeker.typ == tokenDot || peeker.typ == tokenEOF {
					return nil, ErrParser //not in grammar
				}
				p.backToken(peeker) //back at dot
				secondArg, err := p.parseNextSexpr()
				if err != nil {
					return nil, err
				}
				if secondArg.isNil() { //cannot cdr to nil
					return nil, ErrParser
				}

				elements = append(elements, secondArg)
			} else {
				p.backToken(nxt)
				tempCdr, err := p.parseNextSexpr()
				if err != nil { // error exists
					return nil, err
				}

				/* arg, err = p.parseNextSexpr() //cdr val?
				if err != nil {               // error exists
					return nil, err
				} */
				addToElem := mkNil()         //placeholder
				peeker, err := p.nextToken() //check if sexpr
				if err != nil {
					return nil, err
				}
				if peeker.typ == tokenDot { // set car and cdr
					peeker, err = p.nextToken() //ADDEDDDDDDDD
					if peeker.typ == tokenRpar || peeker.typ == tokenDot || peeker.typ == tokenEOF {
						return nil, ErrParser //not in grammar
					}
					p.backToken(peeker) //back at dot
					//no idea abt this part
					secondArg, err := p.parseNextSexpr()
					if err != nil {
						return nil, err
					}
					if secondArg.isNil() { //cannot cdr to nil
						return nil, ErrParser
					}
					addToElem = mkConsCell(tempCdr, secondArg)
				} else {
					secondArg := mkNil()
					if peeker.typ != tokenEOF && peeker.typ != tokenRpar {
						p.backToken(peeker)
						secondArg, err = p.parseNextSexpr()
						if err != nil {
							return nil, err
						}
						if secondArg.isNil() { //cannot cdr to nil
							return nil, ErrParser
						}

						if secondArg.isAtom() {
							secondArg = mkConsCell(secondArg, mkNil())
						}

						addToElem = mkConsCell(tempCdr, secondArg)

					} else {
						p.backToken(peeker)
						addToElem = mkConsCell(tempCdr, mkNil())
					}

				}

				elements = append(elements, addToElem) //append all the atoms made arg is first car
			}
		}

		if nxt.typ != tokenRpar {
			return nil, ErrParser
		}

		//return p.mkCaseExpr(elements), nil //make the cons here?
		return mkConsCell(arg, elements[1]), nil

	case tokenQuote:
		nxt, err := p.nextToken()
		if err != nil { // error exists
			return nil, err
		}
		if nxt.typ == tokenEOF { // has to have something after
			return nil, ErrParser
		}

		//has something after
		p.backToken(nxt)
		quoteSexpr, err := p.mkSimpleSexpr(tokenSymbol, tok), nil
		if err != nil { // error exists
			return nil, err
		}
		cdrCar, err := p.parseNextSexpr()
		if err != nil { // error exists
			return nil, err
		}

		cdrForQuote := mkConsCell(cdrCar, mkNil())
		quotedCons := mkConsCell(quoteSexpr, cdrForQuote)
		return quotedCons, nil
	default:
		return nil, ErrParser
	}
}

func (p *ParserImpl) mkSimpleSexpr(typ tokenType, tok *token) *SExpr { //makes atoms
	key := tok.literal
	if typ == tokenNumber {
		key = tok.num.String()
	}

	//fmt.Printf("NEW TEST Key val is %s /n", key)
	sexpr, ok := p.sexpr[key]
	if !ok {
		switch typ {
		case tokenSymbol:
			//fmt.Printf("GOT INSIDE CASE FOR SYMBOL %s", key)
			//sexpr = &SExpr{atom: mkTokenSymbol(key), car: nil, cdr: nil}
			sexpr = mkSymbol(tok.literal)
		case tokenNumber:
			//fmt.Printf("GOT INSIDE CASE FOR VAL %s", key)
			//sexpr = &SExpr{atom: mkTokenNumber(key), car: nil, cdr: nil}
			sexpr = mkNumber(tok.num)
		case tokenRpar: //makes nil atom for empty case
			sexpr = mkNil()
		}

		p.insertSexpr(sexpr, key)
	}

	return sexpr
}

func (p *ParserImpl) mkCaseExprRecur(elem []*SExpr, index int) *SExpr { //recursively makes the list
	if index == len(elem) {
		return elem[index-1]
	} else {
		tempSexpr := elem[index]
		return mkConsCell(tempSexpr, p.mkCaseExprRecur(elem, index+1))
	}
}
func (p *ParserImpl) mkCaseExpr(elem []*SExpr) *SExpr { //first elem is the car
	listVal := p.mkCaseExprRecur(elem, 0)
	return listVal //placeholder
}

func (p *ParserImpl) insertSexpr(sexpr *SExpr, key string) {
	p.sexpr[key] = sexpr
	p.sexprID[sexpr] = p.sexprCounter
	p.sexprCounter++
}


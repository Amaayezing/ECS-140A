package unify

import (
	"errors"
	"fmt"
	"hw4/disjointset"
	"hw4/term"
)

// ErrUnifier is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrUnifier = errors.New("unifier error")

// UnifyResult is the result of unification. For example, for a variable term
// `s`, `UnifyResult[s]` is the term which `s` is unified with.
type UnifyResult map[*term.Term]*term.Term

// Unifier is the interface for the term unifier.
// Do not change the definition of this interface
type Unifier interface {
	Unify(*term.Term, *term.Term) (UnifyResult, error)
}

// NewUnifier creates a struct of a type that satisfies the Unifier interface.
func NewUnifier() Unifier {
	return &UnifierImpl{
		numToTermClass:    disjointset.NewDisjointSet(),
		classPointerInt:   make(map[*term.Term]int),
		intToClassPointer: make(map[int]*term.Term),
		schema:            make(map[*term.Term]*term.Term),
		boolVisited:       make(map[*term.Term]bool),
		boolAcyclic:       make(map[*term.Term]bool),
		vars:              make(map[*term.Term][]*term.Term),
		finalResult:       make(UnifyResult),
		counter:           0,
	}
}

//Solution begins
type UnifierImpl struct { //term DAG for s and t with shared vars
	numToTermClass    disjointset.DisjointSet
	classPointerInt   map[*term.Term]int
	intToClassPointer map[int]*term.Term
	schema            map[*term.Term]*term.Term
	boolVisited       map[*term.Term]bool
	boolAcyclic       map[*term.Term]bool
	vars              map[*term.Term][]*term.Term
	finalResult       UnifyResult
	counter           int
}

func (u *UnifierImpl) classInit(t *term.Term) { //initializing the class and schema
	if t.Typ != term.TermCompound {
		_, okT1 := u.classPointerInt[t]
		if !okT1 {
			u.classPointerInt[t] = u.counter
			u.intToClassPointer[u.counter] = t
			u.numToTermClass.FindSet(u.counter)
			u.schema[t] = t

			/* fmt.Print("at ", u.counter, ": ")
			u.printDebug(t)
			fmt.Print(" schema: ")
			u.printDebug(u.schema[t])
			fmt.Print(" \n") */
			u.counter = u.counter + 1
		}

	} else { // is a compound
		_, okT1 := u.classPointerInt[t]
		if !okT1 {
			u.classPointerInt[t] = u.counter
			u.intToClassPointer[u.counter] = t
			u.numToTermClass.FindSet(u.counter)
			u.schema[t] = t
			u.counter = u.counter + 1
		}
		for _, vals := range t.Args {
			u.classInit(vals)
		}
	}
}

func (u *UnifierImpl) varsInit(t *term.Term) {
	if t.Typ == term.TermVariable {
		var temp []*term.Term
		temp = append(temp, t)
		u.vars[t] = temp
	} else if t.Typ == term.TermCompound {
		var temp []*term.Term //empty set
		u.vars[t] = temp

		if len(t.Args) > 0 { //not empty
			for _, addedVars := range t.Args {
				u.varsInit(addedVars)
			}
		}

	} else { //num or atom
		var temp []*term.Term //empty set
		u.vars[t] = temp
	}
}

func (u *UnifierImpl) Unify(termOne *term.Term, termTwo *term.Term) (UnifyResult, error) {
	/* //init class
	u.classPointerInt[termOne] = u.counter
	u.intToClassPointer[u.counter] = termOne
	u.numToTermClass.FindSet(u.counter) //add to DS
	u.counter = u.counter + 1           //inc
	u.classPointerInt[termTwo] = u.counter
	u.intToClassPointer[u.counter] = termTwo
	u.numToTermClass.FindSet(u.counter) //add to DS
	//init schema
	u.schema[termOne] = termOne
	u.schema[termTwo] = termTwo */
	u.classInit(termOne)
	u.classInit(termTwo)
	//init vars list
	u.varsInit(termOne)
	u.varsInit(termTwo)

	err := u.unifyClosure(termOne, termTwo)
	if err != nil {
		return nil, ErrUnifier
	}

	err = u.findSolution(termOne)
	if err != nil {
		return nil, ErrUnifier
	}
	return u.finalResult, nil

}

func (u *UnifierImpl) unifyClosure(termOne *term.Term, termTwo *term.Term) error {
	s := u.numToTermClass.FindSet(u.classPointerInt[termOne])
	t := u.numToTermClass.FindSet(u.classPointerInt[termTwo])

	if s == t { //do nothing

	} else {
		schemaS := u.schema[u.intToClassPointer[s]]
		schemaT := u.schema[u.intToClassPointer[t]]

		//debug
		/* fmt.Print("unify schemaS: ")
		u.printDebug(schemaS)
		fmt.Print("\n")
		fmt.Print("unify schemaT: ")
		u.printDebug(schemaT)
		fmt.Print("\n") */

		if (schemaS.Typ != term.TermVariable) && (schemaT.Typ != term.TermVariable) {
			var sLiteral string
			var tLiteral string

			if schemaS.Typ == term.TermCompound {
				sLiteral = schemaS.Functor.Literal
			} else {
				sLiteral = schemaS.Literal
			}

			if schemaT.Typ == term.TermCompound {
				tLiteral = schemaT.Functor.Literal
			} else {
				tLiteral = schemaT.Literal
			}

			if sLiteral == tLiteral { //compare atoms
				u.union(s, t)
				//n := len(u.intToClassPointer[s].Args)
				n := len(u.intToClassPointer[s].Args)
				m := len(u.intToClassPointer[t].Args)
				if m != n {
					return ErrUnifier
				}
				for i := 0; i < n; i++ {
					sI := u.intToClassPointer[s].Args[i]
					tI := u.intToClassPointer[t].Args[i]
					//debug
					/* fmt.Print("IN UNIF FOR \n")
					fmt.Print("SI: ")
					u.printDebug(sI)
					fmt.Print("\n")
					fmt.Print("TI: ")
					u.printDebug(tI)
					fmt.Print("\n") */

					err := u.unifyClosure(sI, tI)

					if err != nil {
						return ErrUnifier
					}
				}
			} else {
				return ErrUnifier //symbol clash
			}
		} else {
			//union s,t
			u.union(s, t)
		}
	}

	return nil
}

func (u *UnifierImpl) union(s int, t int) {
	sTerm := u.intToClassPointer[s]
	tTerm := u.intToClassPointer[t]

	//debug
	/* fmt.Print("begS: ")
	u.printDebug(sTerm)
	fmt.Print("\n")
	fmt.Print("begT: ")
	u.printDebug(tTerm)
	fmt.Print("\n") */

	schemaS := u.schema[u.intToClassPointer[s]]
	schemaT := u.schema[u.intToClassPointer[t]]

	//debug
	/* fmt.Print("schemaS: ")
	u.printDebug(schemaS)
	fmt.Print("\n")
	fmt.Print("schemaT: ")
	u.printDebug(schemaT)
	fmt.Print("\n") */

	//union the sets
	u.numToTermClass.UnionSet(s, t)

	//debug
	/* fmt.Print("New Rep: ")
	u.printDebug(u.intToClassPointer[u.numToTermClass.FindSet(s)])
	fmt.Print("\n") */

	if u.numToTermClass.FindSet(s) == s { //s is the rep
		//concatentate
		u.vars[sTerm] = append(u.vars[sTerm], u.vars[tTerm]...)
		if schemaS.Typ == term.TermVariable {
			u.schema[sTerm] = schemaT //changed from schema S to sterm
			//debug
			/* fmt.Print("NEW SCHEMA: ")
			u.printDebug(u.schema[sTerm])
			fmt.Print("\n") */
		}

	} else { // t is the new rep
		u.vars[tTerm] = append(u.vars[tTerm], u.vars[sTerm]...)
		if schemaT.Typ == term.TermVariable {
			u.schema[tTerm] = schemaS
			//debug
			/* fmt.Print("NEW SCHEMA for: ")
			u.printDebug(tTerm)
			fmt.Print(" its ")
			u.printDebug(u.schema[tTerm])
			fmt.Print("\n") */
		}
	}
	//debug
	fmt.Print("\n")
}

func (u *UnifierImpl) printDebug(t *term.Term) {
	if t.Typ == term.TermCompound { //debug
		fmt.Print(" ", t.Functor.Literal, "(")
		for _, i := range t.Args {
			u.printDebug(i)
		}
		fmt.Print(")")
	} else {
		fmt.Print(" ", t.Literal, " ")
	}
}

func (u *UnifierImpl) findSolution(node *term.Term) error {
	sRep := u.numToTermClass.FindSet(u.classPointerInt[node])
	schemaS := u.schema[u.intToClassPointer[sRep]]
	//debug
	/* fmt.Print("Node: ")
	u.printDebug(node)
	fmt.Print("\n")
	fmt.Print("SREP: ")
	u.printDebug(u.intToClassPointer[sRep])
	fmt.Print("Schema: ")
	u.printDebug(schemaS)
	fmt.Print("\n")
	fmt.Print("Acyclic: ", u.boolAcyclic[schemaS])
	fmt.Print("\n")
	fmt.Print("Visited: ", u.boolVisited[schemaS])
	fmt.Print("\n")
	fmt.Print("\n") */

	if u.boolAcyclic[schemaS] {
		return nil
	}
	if u.boolVisited[schemaS] {
		return ErrUnifier //exists a cycle
	}
	if schemaS.Typ == term.TermCompound && len(schemaS.Args) > 0 {
		u.boolVisited[schemaS] = true
		for _, val := range schemaS.Args {
			err := u.findSolution(val)
			if err != nil {
				return ErrUnifier
			}
		}
		u.boolVisited[schemaS] = false
	}
	u.boolAcyclic[schemaS] = true
	sRep = u.numToTermClass.FindSet(u.classPointerInt[schemaS]) //Find(s)
	for _, vals := range u.vars[u.intToClassPointer[sRep]] {
		/* if vals.Typ != schemaS.Typ {
			//add to sol UNFINISHED PART
			u.finalResult[vals] = schemaS //not sure abt this usage
		} else if vals.Literal != schemaS.Literal {
			u.finalResult[vals] = schemaS
		} */
		if vals != schemaS {
			u.finalResult[vals] = schemaS
		}
	}

	return nil //ending
}

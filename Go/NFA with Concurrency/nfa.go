package nfa

import (
	"sync"
)

// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	var wg sync.WaitGroup
	c := make(chan bool)
	quit := make(chan bool)
	var mux sync.Mutex

	wg.Add(1)
	go recurTraverse(transitions, start, final, input, &wg, &mux, c, quit)

	go func() { //executes if all threads done
		wg.Wait()
		close(c)
	}()

	rVal := <-c
	if rVal {
		quit <- true
	}

	return rVal
}

func recurTraverse(
	transitions TransitionFunction,
	start, final state,
	input []rune,
	wg *sync.WaitGroup,
	mux *sync.Mutex,
	c, quit chan bool) {

	defer wg.Done()

	var qVal bool
	go func() {
		qVal = <-quit
	}()
	//if qVal {
	//	return
	//}

	if len(input) == 0 {
		mux.Lock()
		result := (start == final)
		if result {
			c <- true
		}
		mux.Unlock()
		return
	}

	for _, next := range transitions(start, input[0]) {
		wg.Add(1)
		go recurTraverse(transitions, next, final, input[1:], wg, mux, c, quit)
	}

	return
}

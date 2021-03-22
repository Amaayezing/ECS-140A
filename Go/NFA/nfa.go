package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

// You may define helper functions here.

func reachableRecur(transition TransitionFunction, currState state, final state, inputs []rune) bool {
	if len(inputs) == 0 {
		if currState == final {
			return true
		} else {
			return false
		}
	} else {
		nextStates := transition(currState, inputs[0])
		for i := 0; i < len(nextStates); i++ {
			if reachableRecur(transition, nextStates[i], final, inputs[1:]) {
				return true
			}
		}
	}

	return false
}

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO: Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise

	return reachableRecur(transitions, start, final, input)
}

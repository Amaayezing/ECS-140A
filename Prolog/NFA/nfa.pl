reachable(Nfa, StartState, FinalState, Input) :-
	Input = [H|T], transition(Nfa, StartState, H, FinalState) = [A|_], reachable(Nfa, A, FinalState, T).

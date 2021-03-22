/* fooTransitions */
% 0 -a-> 1
% 0 -a-> 2
% 1 -b-> 3
% 2 -c-> 3
transition(fooTransitions, 0, a, [1,2]).
transition(fooTransitions, 0, b, []).
transition(fooTransitions, 0, c, []).
transition(fooTransitions, 1, a, []).
transition(fooTransitions, 1, b, [3]).
transition(fooTransitions, 1, c, []).
transition(fooTransitions, 2, a, []).
transition(fooTransitions, 2, b, []).
transition(fooTransitions, 2, c, [3]).

/* expTransitions */
% 0 -a-> 1
% 0 -a-> 2
% 0 -b-> 2
% 1 -b-> 0
transition(expTransitions, 0, a, [1,2]).
transition(expTransitions, 0, b, [2]).
transition(expTransitions, 1, a, []).
transition(expTransitions, 1, b, [0]).
transition(expTransitions, 2, a, []).
transition(expTransitions, 2, b, []).

/* langTransitions */
% 0 -a-> 0
% 0 -b-> 1
% 1 -a-> 1
% 1 -b-> 0
transition(langTransitions, 0, a, [0]).
transition(langTransitions, 0, b, [1]).
transition(langTransitions, 1, a, [1]).
transition(langTransitions, 1, b, [0]).

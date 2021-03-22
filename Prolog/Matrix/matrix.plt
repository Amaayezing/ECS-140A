:- initialization main.

main :-
    consult(['matrix.pl']),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(mat).

test(are_adjacent01, [fail])   :- are_adjacent([], 1, 1).
test(are_adjacent02, [fail])   :- are_adjacent([1], 1, 1).
test(are_adjacent03, [nondet]) :- are_adjacent([1, 1], 1, 1).
test(are_adjacent04, [fail])   :- are_adjacent([1, 2, 3], 1, 3).
test(are_adjacent05, [fail])   :- are_adjacent([1, 2, 3], 3, 1).
test(are_adjacent06, [nondet]) :- are_adjacent([1, 2, 3], 1, 2).
test(are_adjacent07, [nondet]) :- are_adjacent([1, 2, 3], 2, 3).
test(are_adjacent08, [nondet]) :- are_adjacent([1, 2, 3], 3, 2).
test(are_adjacent09, [nondet]) :- are_adjacent([1, 2, 3], 2, 1).
test(are_adjacent10, [fail])   :- are_adjacent([1, 2, 3], 1, 1).
test(are_adjacent11, [nondet]) :- are_adjacent([1, 2, 3], 1, 2).
test(are_adjacent12, [fail])   :- are_adjacent([1, 2, 1], 1, 4).
test(are_adjacent13, [nondet]) :- are_adjacent([1, 2, 1, 4], 1, 4).

test(matrix_transpose01, [nondet]) :- matrix_transpose([], X), X == [].
test(matrix_transpose02, [nondet]) :- matrix_transpose([[1, 2, 3, 4]], X), X == [[1], [2], [3], [4]].
test(matrix_transpose03, [nondet]) :- matrix_transpose([[1], [2], [3], [4]], X), X == [[1, 2, 3, 4]].
test(matrix_transpose04, [nondet]) :- matrix_transpose([[1, 2], [3, 4]], X), X == [[1, 3], [2, 4]].
test(matrix_transpose05, [nondet]) :- matrix_transpose([[1, 3], [2, 4]], X), X == [[1, 2], [3, 4]].

test(are_neighbors01, [fail])   :- are_neighbors([], 1, 2).
test(are_neighbors02, [nondet]) :- are_neighbors([[1, 2, 3]], 1, 2).
test(are_neighbors03, [fail])   :- are_neighbors([[1, 2, 3]], 1, 3).
test(are_neighbors04, [nondet]) :- are_neighbors([[1], [2], [3]], 1, 2).
test(are_neighbors05, [nondet]) :- are_neighbors([[1, 2, 3], [4, 5, 6]], 1, 2).
test(are_neighbors06, [fail])   :- are_neighbors([[1, 2, 3], [4, 5, 6]], 2, 6).
test(are_neighbors07, [fail])   :- are_neighbors([[1, 2, 3], [4, 5, 6]], 1, 6).

:- end_tests(mat).

% A list is a 1-D array of numbers.
% A matrix is a 2-D array of numbers, stored in row-major order.

% You may define helper functions here.

% are_adjacent(List, A, B) returns true iff A and B are neighbors in List.
are_adjacent(List, A, B) :-
	append(_, [A, B | _ ], List); append(_, [B, A | _ ], List).

% matrix_transpose(Matrix, Answer) returns true iff Answer is the transpose of
% the 2D matrix Matrix.
matrix_transpose([],[]).
matrix_transpose(Matrix, Answer) :-
	Matrix = Answer.

% are_neighbors(Matrix, A, B) returns true iff A and B are neighbors in the 2D
% matrix Matrix.
are_neighbors(Matrix, A, B) :-
	append(_, [A, B | _ ], Matrix); append(_, [B, A | _ ], Matrix).

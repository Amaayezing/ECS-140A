% helper functions

/* check to see if a list is in another list */
list_find(L, [H|_]) :-
       member(H, L).

list_find(L, [H|T]) :-
        not(member(H, L)), list_find(T, L).

/* check to see what mutual books fans have in common */
mutual_books(Book, A, B) :-
	fan(A, X), fan(B, Y), member(Book, X), member(Book, Y).


/* All novels published either during the year 1953 or during the year 1996*/
year_1953_1996_novels(Book) :-
	novel(Book, 1953);
	novel(Book, 1996).

/* List of all novels published during the period 1800 to 1900 (not inclusive)*/
period_1800_1900_novels(Book) :-
	novel(Book, Year), Year > 1800, Year < 1900.

/* Characters who are fans of LOTR */
lotr_fans(Fan) :-
	fan(Fan, X), member(the_lord_of_the_rings, X).

/* Authors of the novels that heckles is fan of. */
heckles_idols(Author) :-
	author(Author, A), fan(heckles, B), list_find(A, B).

/* Characters who are fans of any of Robert Heinlein's novels */
heinlein_fans(Fan) :-
	fan(Fan, A), author(robert_heinlein, B), list_find(A, B).

/* Novels common between either of Phoebe, Ross, and Monica */
mutual_novels(Book) :-
	mutual_books(Book, phoebe, ross); mutual_books(Book, phoebe, monica); mutual_books(Book, monica, ross).

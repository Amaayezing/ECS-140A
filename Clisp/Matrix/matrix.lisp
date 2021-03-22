; A list is a 1-D array of numbers.
; A matrix is a 2-D array of numbers, stored in row-major order.

; If needed, you may define helper functions here.

; AreAdjacent returns true iff a and b are adjacent in lst.
(defun are-adjacent (lst a b)
  (cond ((null lst) nil)
	((null (car lst)) nil)
	((null a) nil)
	((null b) nil)
	((and (equal a (car lst)) (equal b (car (cdr lst)))) t)
	((and (equal b (car lst)) (equal a (car (cdr lst)))) t)
	((are-adjacent (cdr lst) a b)))
)

; Transpose returns the transpose of the 2D matrix mat.
(defun transpose (matrix)
  (if matrix (apply #'mapcar #'list matrix))
)

; AreNeighbors returns true iff a and b are neighbors in the 2D
; matrix mat.
(defun are-neighbors (matrix a b)
  (cond ((null matrix) nil)
	((null (car matrix)) nil)
	((null a) nil)
	((null b) nil)
	((and (equal a (car (car matrix))) (equal b (car (cdr (car matrix))))) t)
	((and (equal a (car (car matrix))) (equal b (car (car (cdr matrix))))) t)
	((are-neighbors (cdr matrix) a b)))
)


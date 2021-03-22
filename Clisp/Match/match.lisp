; You may define helper functions here

; helper function to check if the list contains the item
(defun contains (item str)
  (cond ((null str) ())
	((equal (car str) item) str)
	((equal (car '(?)) item) (contains item (cdr str)))
	((equal (car '(!)) item) (contains item (cdr str)))
	(t (contains item (cdr str)))))

(defun match (pattern assertion)
  (cond ((and (null pattern)(null assertion)) t)
	((contains (car pattern) assertion) (match (cdr pattern) assertion))
	(t()))
)

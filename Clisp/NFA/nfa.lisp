;; You may define helper functions here

(defun reachable (transition start final input)
  (cond ((null start) nil)
	((null final) nil)
	((null input) t) 
	((reachable transition (car (funcall transition start (car input))) final (cdr input))))
)

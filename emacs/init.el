;; ;; some relatively difficult keybindings to keep in mind
;; C-h d for apropros-documentation
;; C-x 4 b for switch-to-buffer-other-windwos

(defun ay/complete-symbol-and-other-window ()
  (interactive)
  (let ((window-count (length (window-list))))
	(complete-symbol nil)
	(if (> (length (window-list))  window-count)
	    (other-window 1)
	  (message "Completion match not found!"))))
(keymap-global-set "s-." #'ay/complete-symbol-and-other-window)
(set-frame-font "JetBrains Mono" nil t)
(set-face-attribute 'default nil
		    :height 140)
;; if things go wrong there's also...
(add-to-list 'default-frame-alist '('font . "JetBrains Mono-14"))

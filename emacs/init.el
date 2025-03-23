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
(menu-bar-showhide-tool-bar-menu-customize-disable)

(defun ay/load-pkg (pkg) (unless (package-installed-p pkg)
  (package-install pkg))
  (require pkg))
(ay/load-pkg 'geiser)
(ay/load-pkg 'geiser-racket)

(menu-bar-showhide-tool-bar-menu-customize-disable)
(menu-bar--visual-line-mode-enable)

(if (eq system-type 'windows-nt)
    (setq w32-use-visible-system-caret nil)) ;; ay caramba https://stackoverflow.com/a/47447326/11849530
(if (eq system-type 'gnu/linux)
    (global-unset-key (kbd "C-z")))
(set-face-font 'default "JetBrains Mono" nil)
(set-face-attribute 'default nil :height 160)

(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/") t)
(package-refresh-contents)
(defun ay/pkg-load (pkg)
  (interactive)
  (unless (package-install pkg)
    (package-refresh-contents)
     (package-installed-p pkg)
     (message ":::%s installed:::" pkg))
   (require pkg))

(ay/pkg-load 'geiser)
(ay/pkg-load 'geiser-racket)
(ay/pkg-load 'csharp-mode)
(ay/pkg-load 'powershell)
(ay/pkg-load 'ob-powershell)
(menu-bar-showhide-tool-bar-menu-customize-disable)
(menu-bar--visual-line-mode-enable)

(if (eq system-type 'windows-nt)
    (setq w32-use-visible-system-caret nil)) ;; ay caramba https://stackoverflow.com/a/47447326/11849530
(if (eq system-type 'gnu/linux)
    (global-unset-key (kbd "C-z")))
(when (eq system-type 'darwin)
  (setq mac-command-modifier 'meta)
  (setq mac-option-modifier 'super))
(set-face-font 'default "JetBrains Mono" nil)
(set-face-attribute 'default nil :height 160)

(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/") t)
(package-refresh-contents)
(defun ay/pkg-load (pkg)
  (interactive)
  (unless (package-install pkg)
    (package-refresh-contents)
     (package-installed-p pkg)
     (message ":::%s installed:::" pkg))
   (require pkg))

(ay/pkg-load 'geiser)
(ay/pkg-load 'geiser-racket)
(ay/pkg-load 'csharp-mode)
(ay/pkg-load 'powershell)
(ay/pkg-load 'ob-powershell)
(setq inferior-lisp-program "sbcl")
(ay/pkg-load 'slime)
(slime-setup '(slime-fancy))
(ay/pkg-load 'beacon)
(beacon-mode 1)
(ay/pkg-load 'lsp-mode)
(ay/pkg-load 'go-mode)
(add-hook 'before-save-hook #'gofmt-before-save)
;;
(ay/pkg-load 'anki-editor)
(ay/pkg-load 'corfu)
(setq corfu-auto t)
;;
(ay/pkg-load 'org-tempo)


(if #t "Hello out there" "Good bye in here")

(let [(x 12)] (print x))
;; #reader #s #x #t #f #\ƛ
(let [(ht1 #hash([Topos  . #e1.2] [Greater . #i1.2]))]
  (hash-set ht1 'Topos "Not the original value")
  (display-ln (hash-ref ht1 'Topos   )))

(let ([k 0])
  (set! k (for/fold ([s 0])
	      ([i '(1 2 3 4 5)])
	    (+ s i)))
  (- k 10))
(p (define k 11)
   (set! k (for/fold ([s 0])
	       ([i '(1 2 3 4 5)])
	     (+ s i)))
   (- k 10))

(define-struct s1 (f1 f2) #:prefab)
(println #s(s1 (#e1.23 "Just a string in f2")))

(define ihash1 #hash(['first . #e11.2]))
(define ihash2 (hash-set ihash1 'second #i3.14))
(printf "ihash1['second]=~a and ihash2['second]=~a"
	(hash-ref ihash1 'second '-)
	(hash-ref ihash2 'second '-))

(define iter-hash1 (lambda (h)
		     (for ([(k v) (in-hash h)])
		       (printf "~a: ~a\n" k v))))

(iter-hash1 ihash2)

(define ihash2 #hash((one . "first") (two . 'secondarily)))
(iter-hash1 ihash2)

(for ([x '(a b c)] [y '(1 2 3)])
  (printf "#~a. ~a\n" y x))

(define k #| block comment is opened |# 'and-closed)
k
#;(printf "This is run\n")
					; s-expression comments?
#;
(printf "luigi\n")
#b101

1+2i
(printf "~a ≅ ~a\n" (exp 3) (expt 2.71828 3))

(quotient 100 17)
(remainder 100 17)
(/ 1 3)
(sqr (+ (/ 1 (sqrt 2)) (* (/ 1 (sqrt 2)) 0+1i))

     (sqr 0+1i)

     (define (fn1 lst) (for ((x lst))
			 (displayln x)))
     (define (fn1 lst) (for-each displayln lst))

     (fn1 '(one two three))
     (define/contract (ay/adder lst)
       (-> (listof exact-integer?) exact-integer?)
       (displayln "Adder run...")
       (for/fold ((acc 0)) ((x lst))
	 (+ acc x)))
     (ay/adder '(1 2 3 4 5)) ;; 15
     
     (define (ay/dedup lst)
       (-> (listof exact-integer?) (listof exact-integer?))
       (define freq (for/fold ((h1 (make-hash)))
			((x lst))
		      (hash-set! h1 x (add1 (hash-ref h1 x 0)))
		      h1))
       for (((k v) (in-hash freq)))
       (printf "~ax[~a]\n" v k))
     (hash-keys freq))
(ay/dedup '(1 1 1 1 2 3 2 3 4 6))

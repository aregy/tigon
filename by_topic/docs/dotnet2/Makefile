build:
	openssl genrsa 2048 > key2.key
	chmod 400 key2.key
	openssl req -new -x509 -noes -sha256 -days 365 -key key2.key -out cert.cert
	openssl pkcs12 -export -out pfx2.pfx -inkey key2.key -in cert.cert

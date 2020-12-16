rm *.pem
rm *.srl
rm *.cert
rm *.key

# 1. Generate CA's private key and self-signed certificate
openssl genrsa -out ca.key 4096
openssl req -new -x509 -key ca.key -sha256 -subj "/C=GE/ST=LowerSaxony/L=Braunschweig/O=TUD/OU=Education/CN=sconedocs.github.io/emailAddress=yateenderk@gmail.com" -days 365 -out ca.cert
openssl genrsa -out service.key 4096
openssl req -new -key service.key -out service.csr -config certificate.conf
openssl x509 -req -in service.csr -CA ca.cert -CAkey ca.key -CAcreateserial \
		-out service.pem -days 365 -sha256 -extfile certificate.conf -extensions req_ext
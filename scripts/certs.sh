#!/usr/bin/env bash

echo "Generating certs..."
find ../ -type f \( -name "*.pem" -o -name "*.srl" \) -exec rm -rf {} \;

# Generate CA's private key and self-signed certificate
openssl req \
    -x509 \
    -newkey rsa:4096 \
    -days 365 \
    -nodes \
    -keyout ../certificates/assets/ca-key.pem \
    -out ../certificates/assets/ca-cert.pem \
    -subj "/C=US/ST=California/L=San Francisco/O=GitHub, Inc./OU=GitHub/CN=*/emailAddress=elulcao@icloud.com"

echo "CA's self-signed certificate"
openssl x509 \
    -in ../certificates/assets/ca-cert.pem \
    -noout \
    -text

# Generate web server's private key and certificate signing request (CSR)
openssl req \
    -newkey rsa:4096 \
    -nodes \
    -keyout ../certificates/assets/server-key.pem \
    -out ../certificates/assets/server-req.pem \
    -subj "/C=US/ST=California/L=Los Angeles/O=GitHub, Inc./OU=GitHub/CN=*/emailAddress=elulcao@icloud.com"

# Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 \
    -req \
    -in ../certificates/assets/server-req.pem \
    -days 365 \
    -CA ../certificates/assets/ca-cert.pem \
    -CAkey ../certificates/assets/ca-key.pem \
    -CAserial ../certificates/assets/CAserial.srl \
    -CAcreateserial \
    -out ../certificates/assets/server-cert.pem \
    -extfile certificates/server-ext.cnf

echo "Server's signed certificate"
openssl x509 \
    -in ../certificates/assets/server-cert.pem \
    -noout \
    -text

# Generate client's private key and certificate signing request (CSR)
openssl req \
    -newkey rsa:4096 \
    -nodes \
    -keyout ../certificates/assets/client-key.pem \
    -out ../certificates/assets/client-req.pem \
    -subj "/C=US/ST=California/L=Los Angeles/O=GitHub, Inc./OU=GitHub/CN=*/emailAddress=elulcao@icloud.com"

# Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 \
    -req \
    -in ../certificates/assets/client-req.pem \
    -days 60 \
    -CA ../certificates/assets/ca-cert.pem \
    -CAkey ../certificates/assets/ca-key.pem \
    -CAserial ../certificates/assets/CAserial.srl \
    -CAcreateserial \
    -out ../certificates/assets/client-cert.pem \
    -extfile certificates/client-ext.cnf

echo "Client's signed certificate"
openssl x509 \
    -in ../certificates/assets/client-cert.pem \
    -noout \
    -text

# Use CA's to validate certificate
echo "Verifying certificates"
openssl verify -CAfile ../certificates/assets/ca-cert.pem ../certificates/assets/server-cert.pem
openssl verify -CAfile ../certificates/assets/ca-cert.pem ../certificates/assets/client-cert.pem
openssl verify -CAfile ../certificates/assets/ca-cert.pem ../certificates/assets/ca-cert.pem

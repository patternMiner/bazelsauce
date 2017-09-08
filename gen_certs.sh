#!/bin/bash

cd $HOME/.ssh

# generate the private key and the public certificate.
openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out cert.pem

# review the certificate
openssl x509 -text -noout -in cert.pem

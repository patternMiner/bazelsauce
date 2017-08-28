#!/bin/bash

#
# Created by jbisa on 8/26/17.
#

cp $HOME/.ssh/*.pem data

/usr/local/go/bin/go run /Users/jbisa/go/src/github.com/patternMiner/bazelsauce/main.go \
    --cert_path=data/cert.pem --key_path=data/key.pem


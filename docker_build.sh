#!/usr/bin/env bash

cp $HOME/.ssh/*.pem data

docker build -t bazelsauce:v1 .

# docker images list
# docker run -d -p 8080:8080 bazelsauce:v1
# docker ps
# docker stop <container-id>

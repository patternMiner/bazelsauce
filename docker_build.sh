#!/usr/bin/env bash


docker build -t bazelsauce:v1 .

# docker images list
# docker run -d -p 8080:8080 bazelsauce:v1
# docker ps
# docker stop <container-id>

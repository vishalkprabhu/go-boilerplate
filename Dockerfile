FROM golang:1.15-alpine
LABEL maintainer="Vishalkprabhu"


# Setting up Dev environment

WORKDIR /echo_app/
# note this file, go.mod exists locally. and contain reference 
# to direct/indirect dependencies. this step allows to download 
# dependencies and speedup build for docker images (if it used 
# to build artifacts, and not as dev env).  
COPY go.mod  /echo_app/go.mod
RUN go mod download


EXPOSE 1323
# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

#########################################################
RUN     mkdir /app
WORKDIR /app

# Download Go modules
COPY    go.mod ./
COPY    go.sum ./
RUN     go mod download
RUN     go mod verify

COPY    . ./

# Build
RUN     go build -o /main
CMD     [ "/main" ]

# Expose ports
EXPOSE  9091
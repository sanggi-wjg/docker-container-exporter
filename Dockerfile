# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

EXPOSE  9091

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
RUN     go build -o /docker_container_exporter
CMD     [ "/docker_container_exporter" ]

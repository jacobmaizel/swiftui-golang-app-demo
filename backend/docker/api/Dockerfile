# FROM golang:1.18 as base

# FROM base as dev

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# WORKDIR /opt/app/api

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download



RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
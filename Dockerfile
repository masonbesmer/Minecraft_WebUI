FROM golang:latest as base

FROM base as dev

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go install github.com/mitranim/gow@latest
RUN export GOPATH="$HOME/go"
RUN export PATH="$GOPATH/bin:$PATH"

WORKDIR /app

CMD ["gow", "-e=go,html", "run", "."]
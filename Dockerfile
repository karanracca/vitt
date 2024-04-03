FROM golang:1.21-alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/vitt main.go

RUN echo $PATH

CMD ["vitt", "web"]
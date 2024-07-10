FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY *.go .

RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 18080

CMD ["app"]

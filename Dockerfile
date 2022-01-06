FROM golang:1.17

WORKDIR /go/src/data-structures
COPY . .

RUN go get -d ./...
RUN go install ./...

CMD ["data-structures"]
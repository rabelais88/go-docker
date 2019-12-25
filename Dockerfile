FROM golang:latest
WORKDIR /go/src/app
COPY . .

RUN go get
CMD go build main.go && ./main
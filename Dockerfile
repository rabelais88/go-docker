FROM golang:latest
WORKDIR /go/src/app
COPY . .

# RUN go get -d -v ...
CMD go build main.go && ./main
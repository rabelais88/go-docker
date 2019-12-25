# golang package example /w docker
- easier to solve GOPATH issues
- can crosscompile for multiple OS environments on a single computer

```sh
# throwaway build
docker run --rm -v "$PWD":/go/src/app -w /go/src/app golang:latest go build -v

# throwaway run
docker run --rm -v "$PWD":/go/src/app -w /go/src/app golang:latest go run main.go

# /w package
docker build -t ${PREFERRED_NAME}:${PREFERRED_TAG} .
docker run ${PREFERRED_NAME}:${PREFERRED_TAG}
```
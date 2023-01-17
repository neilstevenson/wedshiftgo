FROM registry.access.redhat.com/ubi9/go-toolset:latest

WORKDIR /build

RUN echo Version 7

COPY go.mod /build/
COPY go.sum /build/
COPY cli/   /build/cli/

RUN go mod download
RUN go build -o ./client

CMD ["./client"]

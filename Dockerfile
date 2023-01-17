FROM registry.access.redhat.com/ubi9/go-toolset:latest

WORKDIR /build

RUN echo Version 2

COPY go.mod /build/
COPY go.sum /build/
COPY cli/   /build/cli/

RUN chmod u+w /build/*

RUN go mod tidy
RUN go mod download
RUN go build ./cli/client

CMD ["./client"]

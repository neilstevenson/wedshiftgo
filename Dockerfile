FROM registry.access.redhat.com/ubi9/go-toolset:latest

WORKDIR /build

RUN echo Version 6

COPY go.mod /build/
COPY go.sum /build/
COPY cli/   /build/cli/

RUN ls -al /build

RUN go mod download
RUN go build -x ./cli/client

CMD ["./client"]

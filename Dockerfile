FROM golang:1.16.4-buster AS builder
ARG VERSION=dev  
WORKDIR /go/src/app
COPY main.go .
COPY go.mod .
RUN go build -o main.go .

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]


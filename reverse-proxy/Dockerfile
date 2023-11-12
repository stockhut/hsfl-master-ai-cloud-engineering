FROM golang:1.21-alpine

WORKDIR /go/src/app

COPY reverse-proxy/go.mod .
COPY reverse-proxy/go.sum .


RUN go mod download github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy

COPY . .

RUN go build -o main github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy/cmd

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/src/app/main ./

EXPOSE 8081

CMD ["/app/main"]
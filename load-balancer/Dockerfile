FROM golang:1.21-alpine

WORKDIR /go/src/app

COPY load-balancer/go.mod .
COPY load-balancer/go.sum .


RUN go mod download github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer

COPY . .

RUN go build -o main github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer/cmd

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/src/app/main ./

EXPOSE 5001

CMD ["/app/main"]
FROM golang:1.21-alpine

WORKDIR /go/src/app

COPY go.mod .

RUN go mod download github.com/stockhut/hsfl-master-ai-cloud-engineering/web-service

COPY . .

RUN go build -o main github.com/stockhut/hsfl-master-ai-cloud-engineering/web-service

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/src/app/main ./
COPY --from=0 /go/src/app/public/* ./public

EXPOSE 3000

CMD ["/app/main"]
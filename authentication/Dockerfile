FROM golang:1.21-alpine

WORKDIR /go/src/app

COPY authentication/go.mod .
COPY authentication/go.sum .


RUN go mod download github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication

COPY . .

RUN go build -o main github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/src/app/main ./

EXPOSE 8081

CMD ["/app/main"]
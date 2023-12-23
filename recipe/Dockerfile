FROM golang:1.21-alpine

WORKDIR /go/src/app

COPY recipe/go.mod .
COPY recipe/go.sum .

RUN go mod download github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

COPY . .

RUN go build -o main github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/src/app/main ./

EXPOSE 8081

CMD ["/app/main"]

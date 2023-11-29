FROM golang:1.21

WORKDIR /go/src/app

# for go-sqlite3
#RUN apk --no-cache add gcc g++

COPY recipe/go.mod .
COPY recipe/go.sum .

# install and go-sqlite3 driver in an early layer as it takes forever to build
RUN go install github.com/mattn/go-sqlite3

RUN go mod download github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

COPY . .

RUN go build -o main github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/src/app/main ./

EXPOSE 8081

CMD ["/app/main"]
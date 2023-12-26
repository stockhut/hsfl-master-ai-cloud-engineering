FROM golang:1.21-alpine

WORKDIR /go/src/app

RUN apk add protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY recipe/go.mod .
COPY recipe/go.sum .

RUN go mod download github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

COPY . .

RUN go generate authentication/auth-proto/gen.go

RUN go build -o main github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/src/app/main ./

EXPOSE 8081

CMD ["/app/main"]

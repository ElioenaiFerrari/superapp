FROM golang:1.22-alpine AS builder

ARG APP

WORKDIR /app

# install protoc
RUN apk add --no-cache make git protobuf
COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go install github.com/go-delve/delve/cmd/dlv
COPY . .
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN make protos
RUN go build -o ./main ./cmd/${APP}

# use distroless
FROM alpine:3.9.6

ARG APP
ENV APP=${APP}

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /go/bin/dlv /usr/local/bin/dlv
COPY --from=builder /app/entrypoint.sh .

RUN chmod +x ./entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]

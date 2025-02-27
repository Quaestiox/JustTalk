FROM golang:1.22.12-alpine AS builder
WORKDIR /usr/app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz
RUN curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip && \
    unzip -d /tmp/protoc /tmp/protoc.zip && \
    mv /tmp/protoc/bin/protoc /usr/local/bin/

FROM alpine:latest
WORKDIR /usr/app
COPY --from=builder /usr/app/main .
COPY --from=builder /usr/app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration


EXPOSE 8080
EXPOSE 9101
CMD ["/usr/app/main"]
ENTRYPOINT ["/usr/app/start.sh"]
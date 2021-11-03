FROM golang:1.17.2-alpine3.14

COPY . service/

WORKDIR service/

RUN go mod vendor
RUN go build -o ./bin/service-entrypoint ./cmd/service

CMD ["./bin/service-entrypoint"]
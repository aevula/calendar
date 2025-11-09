FROM golang:1.25.1-alpine3.22 AS base

RUN apk update && apk add --no-cache build-base

WORKDIR /usr/src/

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./


FROM base AS dev

WORKDIR /usr/src/

ENV CONFIG_PATH=./config/dev.yaml

CMD ["go", "run", "--race", "./cmd/server.go"]


FROM base AS prod_base

WORKDIR /usr/src/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/server.go


FROM alpine:3.22 AS prod

WORKDIR /opt/app/

COPY --from=prod_base /usr/src/server ./server
COPY --from=prod_base /usr/src/config/prod.yaml ./config.yaml

# TODO: use ENVS or vault instead of file copying
ENV CONFIG_PATH=./config.yaml

CMD ["./server"]

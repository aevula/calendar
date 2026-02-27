FROM golang:1.25.1-alpine3.22 AS base

RUN apk update && apk add --no-cache build-base

WORKDIR /usr/src/

COPY ./go.mod ./go.sum ./
RUN go mod download
RUN go mod tidy

COPY ./ ./


FROM base AS dev

WORKDIR /usr/src/

ENV CONFIG_PATH=./config/dev.yaml

CMD ["go", "run", "--race", "./cmd/server/server.go"]


FROM base AS prod_build

WORKDIR /usr/src/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/server/server.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./backjobs ./cmd/backjobs/backjobs.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./consumer ./cmd/consumer/consumer.go


FROM alpine:3.22 AS prod

WORKDIR /opt/app/

COPY --from=prod_build /usr/src/server ./
COPY --from=prod_build /usr/src/config/prod.yaml ./config.yaml

# TODO: use ENVS or vault instead of file copying
ENV CONFIG_PATH=./config.yaml

CMD ["./server"]


FROM alpine:3.22 AS prod_backjobs

WORKDIR /opt/app/

COPY --from=prod_build /usr/src/backjobs ./
COPY --from=prod_build /usr/src/config/prod.yaml ./config.yaml

# TODO: use ENVS or vault instead of file copying
ENV CONFIG_PATH=./config.yaml

CMD ["./backjobs"]


FROM alpine:3.22 AS prod_consumer

WORKDIR /opt/app/

COPY --from=prod_build /usr/src/consumer ./
COPY --from=prod_build /usr/src/config/prod.yaml ./config.yaml

# TODO: use ENVS or vault instead of file copying
ENV CONFIG_PATH=./config.yaml

CMD ["./consumer"]

FROM golang:1.25.1-alpine3.22 AS builder

WORKDIR /opt/app/
COPY ./ ./
RUN go mod download

# TODO: rewrite separate builds (e.g. -race for dev)
# TODO: check why always rebuilding without file changes 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server.go


FROM alpine:3.22 AS base

WORKDIR /opt/app/
COPY --from=builder /opt/app/server ./

ENV CONFIG_PATH=./config.yaml
CMD [ "./server" ]


FROM base AS dev

WORKDIR /opt/app/
COPY ./config/dev.yaml ./config.yaml


# TODO: use ENVS or vault instead of file copying
FROM base AS prod

WORKDIR /opt/app/
COPY ./config/prod.yaml ./config.yaml

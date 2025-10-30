FROM golang:1.25.1-alpine3.22 AS dev_base

RUN apk update && apk add --no-cache build-base


FROM golang:1.25.1-alpine3.22 AS prod_base


FROM dev_base AS dev_builder

WORKDIR /opt/app/

COPY ./go.mod ./go.sum ./
RUN go mod download

# TODO: exclude config dir (can not add to .dockerignore rn)
COPY ./ ./

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -race -o ./server ./cmd/server.go


FROM prod_base AS prod_builder

WORKDIR /opt/app/

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/server.go


FROM alpine:3.22 AS dev

WORKDIR /opt/app/

COPY --from=dev_builder /opt/app/server ./server
COPY ./config/dev.yaml ./config.yaml

ENV CONFIG_PATH=./config.yaml

CMD ["./server"]


FROM alpine:3.22 AS prod

WORKDIR /opt/app/

COPY --from=prod_builder ./opt/app/server ./server
COPY ./config/prod.yaml ./config.yaml
# TODO: use ENVS or vault instead of file copying
ENV CONFIG_PATH=./config.yaml

CMD ["./server"]

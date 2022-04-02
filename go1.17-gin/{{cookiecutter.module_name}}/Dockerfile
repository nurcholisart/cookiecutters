# Builder Stage
FROM golang:{{cookiecutter.go_version}}-alpine AS builder

RUN apk update
RUN apk add --no-cache curl bash nano

ARG VERSION
ARG CGO_ENABLED

WORKDIR /app/src
COPY . .

RUN go mod tidy
RUN go build -ldflags "-s -w -X main.version=${VERSION}" -o /app/{{cookiecutter.module_name}} ./cmd

WORKDIR /app

RUN rm -rf src/
ENTRYPOINT ["/app/{{cookiecutter.module_name}}"]


# Production Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/{{cookiecutter.module_name}} ./app
CMD ["./app"]

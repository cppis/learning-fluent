## learning-fluent Root 경로에서 실행
## docker build --rm --no-cache --tag {repo}/fluentlogger --file ./assets/docker/Dockerfile .

ARG APP_IMAGE=golang:1.19-alpine

## build stage
FROM $APP_IMAGE AS builder

WORKDIR /build
COPY . .

WORKDIR /build/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o fluentlogger .

## final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/app/fluentlogger .

CMD ["/app/fluentlogger"]

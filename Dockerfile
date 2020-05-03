FROM golang:alpine as builder

ENV GO111MODULE=on \
  GOOS=linux \
  GOARCH=amd64 \
  APP_HOME=/app

RUN apk update && apk upgrade && \
  apk add --no-cache git

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -ldflags "-s -w" -o shinkan-server cmd/service/main.go

# Run container
FROM alpine:latest

ENV APP_HOME /app

RUN apk --no-cache add ca-certificates && \
  apk add mysql-client

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY --chown=0:0 --from=builder $APP_HOME/shinkan-server $APP_HOME

EXPOSE 8080
ENTRYPOINT ["./shinkan-server"]

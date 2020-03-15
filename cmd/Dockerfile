FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/github.com/AirouTUS/shinkan-server
RUN apk update --no-cache
ADD . $APPDIR/
WORKDIR $APPDIR
RUN go build -mod=vendor -ldflags "-s -w" -o shinkan-server  cmd/main.go
RUN mv shinkan-server /

FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add mysql-client
COPY --from=builder /shinkan-server ./
ENTRYPOINT ["./shinkan-server"]

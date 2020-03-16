FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/github.com/AirouTUS/shinkan-server
ENV GO111MODULE on
RUN \
  apk update --no-cache && \
  mkdir -p $APPDIR
ADD . $APPDIR/
WORKDIR $APPDIR
RUN ls . && ls .. && pwd
RUN go build --mod=vendor -ldflags "-s -w" -o shinkan-server cmd/main.go
RUN mv shinkan-server /

FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add mysql-client
COPY --from=builder /shinkan-server ./
ENTRYPOINT ["./shinkan-server"]

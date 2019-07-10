# docker build --rm --no-cache -t deepdivenow/simple-server:0.1 .

FROM golang:1.12.5-alpine3.9 AS builder

COPY src /go/src
RUN apk add git && \
    go get -v simple-server && \
    go build -v simple-server

FROM alpine:3.9
COPY --from=builder /go/bin/simple-server /bin/
RUN mkdir /app
USER 1001
ENTRYPOINT ["/bin/simple-server"]


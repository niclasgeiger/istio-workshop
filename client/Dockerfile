FROM golang:1.16-alpine3.12 AS build

# Add User
RUN addgroup -g 2000 server
RUN adduser -D -u 1000 -G server server -s /sbin/nologin

RUN mkdir -p /code
WORKDIR /code

ADD cmd ./cmd

FROM build as compile

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o bin/server \
    ./cmd/main.go

FROM scratch

COPY --from=compile \
    /etc/passwd \
    /etc/passwd
COPY --from=compile /code/bin/ /

ADD assets ./assets

USER server

EXPOSE 3000
ENTRYPOINT ["/server"]



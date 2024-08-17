FROM golang:alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build
COPY ../../../go.mod .
COPY ../../../go.sum .
RUN go mod download

COPY ../../../app/chat/cmd/api ./app/chat/cmd/api
COPY ../../../app/chat/cmd/api/etc /api/etc
COPY ../../../app/common ./app/common
COPY ../../../app/internal ./app/internal
RUN go build -ldflags="-s -w" -o /out/socket_server ./app/chat/cmd/api/chatapi.go
CMD ["/out/socket_server","-f","/api/etc/chatapi_docker.yaml"]

FROM scratch
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV TZ Asia/Hong_Kong
WORKDIR /app
COPY --from=builder /out/socket_server /app/socket_server
COPY --from=builder /api/etc /app/etc

CMD ["./socket_server","-f","etc/chatapi_docker.yaml"]
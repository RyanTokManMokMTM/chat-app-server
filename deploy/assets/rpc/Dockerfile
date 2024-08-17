FROM golang:alpine as builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR build
ADD ../../../go.sum .
ADD ../../../go.mod .
RUN go mod download

COPY ../../../app/assets/cmd/rpc ./app/assets/cmd/rpc
COPY ../../../app/common ./app/common
COPY ../../../app/assets/cmd/rpc/etc /rpc/etc
COPY ../../../resources/default /app/resources/default

RUN go build -ldflags "-s -w" -o /out/assetsrpc ./app/assets/cmd/rpc/assets.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV TZ Asia/Hong_Kong
WORKDIR /app

COPY --from=builder /out/assetsrpc /app/assetsrpc
COPY --from=builder /rpc/etc /app/etc
COPY --from=builder /app/resources /app/resources
CMD ["./assetsrpc","-f","etc/assets_docker.yaml"]
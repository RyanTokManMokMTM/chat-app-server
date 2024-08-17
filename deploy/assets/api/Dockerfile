FROM golang:alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build
COPY ../../../go.mod .
COPY ../../../go.sum .
RUN go mod download

COPY ../../../app/assets/cmd/api ./app/assets/cmd/api
COPY ../../../app/assets/cmd/rpc ./app/assets/cmd/rpc
COPY ../../../app/assets/cmd/api/etc /api/etc
COPY ../../../app/common ./app/common
COPY ../../../resources/default /app/resources/default

RUN go build -ldflags="-s -w" -o /out/assetsapi ./app/assets/cmd/api/assetsapi.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV TZ Asia/Hong_Kong
WORKDIR /app
COPY --from=builder /out/assetsapi /app/assetsapi
COPY --from=builder /api/etc /app/etc
COPY --from=builder /app/resources /app/resources
CMD ["./assetsapi","-f","etc/assetsapi_docker.yaml"]
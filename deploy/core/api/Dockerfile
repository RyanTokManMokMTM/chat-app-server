FROM golang:alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build
COPY ../../../go.mod .
COPY ../../../go.sum .
RUN go mod download
#
COPY ../../../app/core/cmd/api ./app/core/cmd/api
COPY ../../../app/core/cmd/rpc ./app/core/cmd/rpc
COPY ../../../app/assets/cmd/rpc ./app/assets/cmd/rpc
COPY ../../../app/core/cmd/api/etc /api/etc
COPY ../../../app/common ./app/common
RUN go build -ldflags="-s -w" -o /out/coreapi ./app/core/cmd/api/coreapi.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV TZ Asia/Hong_Kong
WORKDIR /app
COPY --from=builder /out/coreapi /app/coreapi
COPY --from=builder /api/etc /app/etc
CMD ["./coreapi","-f","etc/coreapi_docker.yaml"]

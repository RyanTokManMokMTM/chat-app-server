FROM golang:alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build
ADD ../../../go.mod .
ADD ../../../go.sum .
RUN go mod download

COPY ../../../app/assets/cmd/api ./app/assets/cmd/api
COPY ../../../app/assets/cmd/rpc ./app/assets/cmd/rpc
COPY ../../../app/assets/cmd/api/etc /api/etc
COPY ../../../app/common ./app/common
RUN go build -ldflags="-s -w" -o /out/assetsapi ./app/assets/cmd/api/assetsapi.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asis/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV TZ Asia/Hong_Kong
WORKDIR /app
COPY --from=builder /out/assetsapi /app/assetsapi
COPY --from=builder /api/etc /app/etc
CMD ["./assetapi","-f","etc/assetsapi.yaml"]
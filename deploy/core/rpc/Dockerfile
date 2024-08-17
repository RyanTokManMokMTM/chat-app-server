FROM golang:alpine as builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
RUN apk update --no-cache && apk add --no-cache tzdata

COPY ../../../go.sum .
COPY ../../../go.mod .
RUN go mod download

COPY ../../../app/core/cmd/rpc ./app/core/cmd/rpc
COPY ../../../app/core/cmd/rpc/etc /rpc/etc
COPY ../../../app/assets/cmd/rpc ./app/assets/cmd/rpc
COPY ../../../app/common ./app/common
COPY ../../../app/internal ./app/internal

RUN go build -ldflags="-s -w" -o /out/corerpc ./app/core/cmd/rpc/core.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV TZ Asia/Hong_Kong
WORKDIR /app
COPY --from=builder /out/corerpc /app/corerpc
COPY --from=builder /rpc/etc /app/etc
CMD ["./corerpc","-f","etc/core_docker.yaml"]

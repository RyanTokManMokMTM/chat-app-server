FROM golang:alpine as builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0

RUN apt update --no-cache && apt update --no-cache tzdata

COPY ../../../go.sum .
COPY ../../../go.mod .
RUN go mod download

COPY ../../../app/chat/cmd/api ./app/chat/cmd/api
COPY ../../../app/chat/cmd/api/etc ./api/etc
COPY ../../../app/common ./app/common
RUN go build -ldflags="-s -w" /out/chatapi ./app/chat/cmd/api/chatapi.go

FROM scratch
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /app
ENV TA Asia/Hong_Kong
COPY --from=builder /out/chatapi /app/chatapi
COPY --from=builder /api/etc /app/etc
CMD ["./chatapi","-f","etc/chatapi.yaml"]
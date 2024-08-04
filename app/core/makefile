api:
	goctl api go -api ./cmd/api/core/core.api --dir=./cmd/api --home=../../tools
rpc:
	goctl rpc protoc ./cmd/rpc/proto/core.proto --go_out=./cmd/rpc/types --go-grpc_out=./cmd/rpc/types --zrpc_out=./cmd/rpc -m
runAPI:
	go run ./cmd/api/*.go
runRPC:
	go run ./cmd/rpc/*.go
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./socket-proto/socket.proto
gen-api:
	goctl api plugin -plugin goctl-swagger="swagger" -api ./apis/chat-app.api -dir ./datas
mock:
	mockgen -destination internal/dao/mock/store.go github.com/ryantokmanmokmtm/chat-app-server/internal/dao Store
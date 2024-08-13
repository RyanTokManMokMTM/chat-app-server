api:
	goctl api go -api ./cmd/api/assets/assets.api --dir=./cmd/api --home=../../tools
rpc:
	goctl rpc protoc ./cmd/rpc/proto/assets.proto --go_out=./cmd/rpc/types --go-grpc_out=./cmd/rpc/types --zrpc_out=./cmd/rpc
up-api:
	go run ./cmd/api/*.go -f "./cmd/api/etc/assetsapi.yaml"
up-rpc:
	go run ./cmd/rpc/*.go -f "./cmd/rpc/etc/assets.yaml"
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./socket-proto/socket.proto
gen-api:
	goctl api plugin -plugin goctl-swagger="swagger" -api ./apis/chat-app.api -dir ./datas

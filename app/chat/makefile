api:
	goctl api go -api ./cmd/api/chat/chat.api --dir=./cmd/api --home=../../tools
socketTypeGen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./cmd/api/socket-proto/socket.proto
up-api:
	go run ./cmd/api/*.go -f "./cmd/api/etc/chatapi.yaml"

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./socket-proto/socket.proto
gen-api:
	goctl api plugin -plugin goctl-swagger="swagger" -api ./apis/chat-app.api -dir ./datas

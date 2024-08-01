api:
	goctl api go -api ./cmd/api/chat/chat.api --dir=./cmd/api --home=../../tools
run:
	go run chatapp.go
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./socket-proto/socket.proto
gen-api:
	goctl api plugin -plugin goctl-swagger="swagger" -api ./apis/chat-app.api -dir ./datas
mock:
	mockgen -destination internal/dao/mock/store.go github.com/ryantokmanmokmtm/chat-app-server/internal/dao Store
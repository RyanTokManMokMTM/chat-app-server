api:
	goctl api go -api ./apis/chat-app.api --dir=./ --home=./tools
run:
	go run chatapp.go
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./socket-proto/socket.proto

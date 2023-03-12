BINARY_NAME=server.exe
USER_NAME=vova


run:	
	go run cmd/server/main.go


test:
	go test .



build: 
	go build -o ${BINARY_NAME} cmd/server/main.go

swag:	
	/home/${USER_NAME}/go/bin/swag init -g cmd/server/main.go

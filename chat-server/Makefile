LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate:
	make generate-chat-server-apimake 

generate-chat-server-apimake:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux server/main.go

copy-to-server:
	scp service_linux root@185.91.54.199:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/olegnemlii/test-server:v0.0.1 .
	docker login -u token -p CRgAAAAAVij4ICYGpgAme-SmVbGflbElpzriL7a- cr.selcloud.ru/olegnemlii
	docker push cr.selcloud.ru/olegnemlii/test-server:v0.0.1

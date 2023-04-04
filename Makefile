fmt:
	@go fmt ./...

setup:
	@go install honnef.co/go/tools/cmd/staticcheck@latest

lint:
	@staticcheck ./...

build:
	@go build -o dist/gueue ./cmd/gueue

example:
	@go build -o dist/examples/$(name)/producer ./examples/$(name)/producer
	@go build -o dist/examples/$(name)/consumer ./examples/$(name)/consumer

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. --go-grpc_opt=paths=source_relative \
         $(file)
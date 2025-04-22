.PHONY: generate run build pprof

generate:
	@mkdir -p gen
	@protoc -I proto -I protovalidate/proto/protovalidate  \
		--go_out=gen \
		--go_opt=paths=source_relative \
		proto/example/v1/example.proto

build: generate
	@go build -o app main.go

run: build
	@./app

pprof:
	@go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=30
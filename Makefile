wasm:
	GOOS=js GOARCH=wasm go build -o ./static/json.wasm ./cmd/wasm/main.go
serve: wasm
	go run ./cmd/server
clean:
	rm -rf ./static/json.wasm
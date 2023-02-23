wasm:
	GOOS=js GOARCH=wasm go build -o ./docs/json.wasm ./cmd/wasm/main.go
serve:
	go run ./cmd/server
clean:
	rm -rf ./docs/json.wasm
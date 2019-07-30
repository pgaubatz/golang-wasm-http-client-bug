build:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .
	GOARCH=wasm GOOS=js go build -o main.wasm main.go

run:
	go run server.go

clean:
	rm -f wasm_exec.js main.wasm

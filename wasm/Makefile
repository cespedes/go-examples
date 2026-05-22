all:	main.wasm wasm_exec.js

run:	all
	gohttp

GOROOT := $(shell go env GOROOT)

main.wasm:	main.go
		GOOS=js GOARCH=wasm go build -o main.wasm main.go

wasm_exec.js:	$(GOROOT)/lib/wasm/wasm_exec.js
		cp $(GOROOT)/lib/wasm/wasm_exec.js .

.PHONY:

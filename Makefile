test:
	go test ./... -v

build:
	go build -o bmsniffer ./cmd/bmsniffer

clean: 
	rm bmsniffer
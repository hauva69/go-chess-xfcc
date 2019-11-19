build:
	$(MAKE) -C bin/

test:
	go test -race ./...
	
build:
	$(MAKE) -C bin/

test:
	go test -race ./...

cover:
	go test -cover ./...
	
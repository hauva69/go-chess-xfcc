build:
	$(MAKE) -C bin/

test:
	go test -race ./...

cover:
	go test -cover ./...

coverprofile:
	go test -coverprofile=profile.out

html: coverprofile
	go tool cover -html=profile.out

	
build:
	go build -mod=vendor

test:
	cd xfcc && go test -race
	
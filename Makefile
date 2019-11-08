build:
	$(MAKE) -C bin/

test:
	cd xfcc && go test -race
	
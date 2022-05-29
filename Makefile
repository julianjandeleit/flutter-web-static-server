.PHONY: install
install:
	go build -buildvcs=false
	cp static-server /usr/local/bin/

.PHONY: install
install:
	go build
	cp static-server /usr/local/bin/

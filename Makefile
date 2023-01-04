FLAGS=-ldflags "-s -w"

.PHONY: all
all: make-install

.PHONY: make-install
make-install:
	go build send.go
	sudo cp send /usr/local/bin

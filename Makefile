build:
	go build -o bin/$(shell basename $(PWD)) cmd/main.go

test:
	go test -v ./...

install: build
	sudo install -m 755 bin/ecat /usr/local/bin/ecat

uninstall:
	sudo rm -f /usr/local/bin/ecat

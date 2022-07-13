.PHONY: build
.PHONY: coverage

build:
	mkdir -p build
	go build -o build ./...

test:
	go test ./...

coverage:
	mkdir -p coverage
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html
	go tool cover -func=coverage/coverage.out

upgrade:
	go install github.com/oligot/go-mod-upgrade@latest
	go-mod-upgrade

clean:
	rm -rf build
	rm -rf coverage

all:
	@make -s clean
	@make -s build
	@make -s test
	@make -s coverage

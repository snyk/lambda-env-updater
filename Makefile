.PHONY: test

test:
	go test ./pkg/...

test-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./pkg/...

install-tools:
	go install github.com/mitchellh/gox@v1.0.1

build:
	go build -o ./bin/lambda-env-updater

release: install-tools
	gox -parallel=2\
	  -osarch="linux/386 linux/amd64 linux/arm linux/arm64 darwin/amd64 darwin/arm64 windows/386 windows/amd64"\
	  -output "bin/lambda-env-updater_{{.OS}}_{{.Arch}}"\
	  ./
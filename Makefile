### Local Environment Variables
BIN_NAME=profiles-generator-cli
BIN_FOLDER=releases

PKG_NAME=github.com/abdulhannanali/protonvpn-ikev2-generator/cli/main

# Multiple architectures
AMD64_ARCH=amd64
386_ARCH=386
ARM_ARCH=arm

# Operating Systems
GOOS_WINDOWS=windows
GOOS_LINUX=linux
GOOS_DARWIN=darwin


lint: golangci-lint run .



compile-linux-amd64: 
	@GOOS=$(GOOS_LINUX) GOARCH=$(AMD64_ARCH) go build $(PKG_NAME) -o $(bash ./make_scripts/binary_name.sh)

test: lint
	go test ./...

clean:
	@go clean
	rm -rf $(BIN_FOLDER)


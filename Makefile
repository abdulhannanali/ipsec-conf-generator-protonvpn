export BIN_NAME=ipsec-protonvpn-cli
export BIN_FOLDER=releases


export PKG_NAME=github.com/abdulhannanali/ipsec-conf-generator-protonvpn/cli/main

generate_bin_name_fn = compile-ipsec-generator-$(0)-$(1)


install:
	@go install ./...




compile-all: compile-linux compile-darwin
compile-linux: compile-linux-amd64 compile-linux-386
compile-darwin: compile-darwin-amd64

compile-linux-amd64: create-releases-directory
	@echo Building Linux AMD64
	GOOS=linux GOARCH=amd64 go build -o \
		$(BIN_FOLDER)/$(BIN_NAME)-linux-amd64 \
		$(PKG_NAME)

compile-linux-386:	create-releases-directory
	@$(echo Building Linux 386)
	GOOS=linux GOARCH=amd64 go build -o \
		$(BIN_FOLDER)/$(BIN_NAME)-linux-386 \
		$(PKG_NAME)

compile-darwin-amd64:	create-releases-directory
	@$(echo "Building Darwin AMD64")
	GOOS=darwin GOARCH=amd64 go build -o \
		$(BIN_FOLDER)/$(BIN_NAME)-darwin-amd64 \
		$(PKG_NAME)

example: clean compile-all
	bash example.sh

clean:
	go clean
	rm -rf $(BIN_FOLDER)

create-releases-directory:
	@mkdir -p $(BIN_FOLDER)

install-lint:
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0


lint: 
	echo "Linting"
	@golangci-lint run cli generator protonvpn

test: lint
	@go test ./...

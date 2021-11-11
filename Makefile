PKG := $(shell head -n 1 go.mod | sed -e 's/^module\ \(.*\)/\1/g')
PKG_LIST := $(shell go list -mod=vendor ${PKG}/...)

APP_LIST = $(shell ls cmd)
APP_BIN_LIST = $(patsubst %, bin/%, $(APP_LIST))

GO_ENV = "GOPROXY=direct GOSUMDB=off"

.PHONY: all lint build clean

all: build

lint: ## Lint the files
	@go vet ${PKG_LIST}
	
$(APP_BIN_LIST): lint
	@CGO_ENABLED=0 go build -mod=vendor -o $@ -ldflags '-w -s' $(subst bin/,cmd/,$@)/*.go

build: $(APP_BIN_LIST) ## Build the binary files

clean: ## Remove previous build
	@rm -rf bin


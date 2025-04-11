# golang settings
GO ?= go
GOLINT ?= golangci-lint
GOOS := $(shell $(GO) env GOOS)
GOARCH := $(shell $(GO) env GOARCH)
BIN := $(abspath ./bin/$(GOOS)_$(GOARCH))
GO_ENV ?= GOPRIVATE=github.com/tusmasoma GOBIN=$(BIN)
GOPATH := $(shell go env GOPATH)

# tools
$(shell mkdir -p $(BIN))

PROTOC_VERSION := 24.4
PROTOC_ZIP := protoc-$(PROTOC_VERSION)-linux-x86_64.zip
$(BIN)/protoc-$(PROTOC_VERSION):
	@if ! command -v protoc &> /dev/null; then \
		echo "Installing protoc..."; \
		curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP); \
		unzip -o $(PROTOC_ZIP) -d $(HOME)/.local; \
		rm -f $(PROTOC_ZIP); \
	fi

PROTOC_GEN_GO_VERSION := v1.31.0
$(BIN)/protoc-gen-go-$(PROTOC_GEN_GO_VERSION):
	unlink $(BIN)/protoc-gen-go || true
	$(GO_ENV) ${GO} install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	mv $(BIN)/protoc-gen-go $(BIN)/protoc-gen-go-$(PROTOC_GEN_GO_VERSION)
	ln -s $(BIN)/protoc-gen-go-$(PROTOC_GEN_GO_VERSION) $(BIN)/protoc-gen-go

PROTOC_GEN_GO_GRPC_VERSION := v1.3.0
$(BIN)/protoc-gen-go-grpc-$(PROTOC_GEN_GO_GRPC_VERSION):
	unlink $(BIN)/protoc-gen-go-grpc || true
	$(GO_ENV) ${GO} install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	mv $(BIN)/protoc-gen-go-grpc $(BIN)/protoc-gen-go-grpc-$(PROTOC_GEN_GO_GRPC_VERSION)
	ln -s $(BIN)/protoc-gen-go-grpc-$(PROTOC_GEN_GO_GRPC_VERSION) $(BIN)/protoc-gen-go-grpc

proto_tools: $(BIN)/protoc-$(PROTOC_VERSION) $(BIN)/protoc-gen-go-$(PROTOC_GEN_GO_VERSION) $(BIN)/protoc-gen-go-grpc-$(PROTOC_GEN_GO_GRPC_VERSION)

.PHONY: proto_gen
proto_gen: proto_tools
	protoc --proto_path=. \
		--proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./proto/*.proto; \

format-proto:
	clang-format -i proto/*.proto
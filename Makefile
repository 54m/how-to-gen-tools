API_GEN_VERSION := v2.4.0
VOLCAGO_VERSION := v1.0.0
SWAG_VERSION := 1.7.6
GOLANGCI_LINT_VERSION := 1.41.0
MOCKGEN_VERSION := 1.6.0

OS_NAME := `echo $(shell uname -s) | tr A-Z a-z`
MACHINE_TYPE := $(shell uname -m)

.PHONY: bootstrap
bootstrap: bootstrap_api_gen bootstrap_volcago bootstrap_swag bootstrap_golangci_lint bootstrap_mockgen

.PHONY: bootstrap_with_curl
bootstrap_with_curl: bootstrap_api_gen_with_curl bootstrap_volcago_with_curl bootstrap_swag_with_curl bootstrap_golangci_lint_with_curl bootstrap_mockgen_with_curl

.PHONY: bootstrap_api_gen
bootstrap_api_gen:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/go-generalize/api_gen/v2/cmd/api_gen@$(API_GEN_VERSION)

.PHONY: bootstrap_api_gen_with_curl
bootstrap_api_gen_with_curl:
	mkdir -p ./bin
	curl -s -L -o ./bin/api_gen.tar.gz https://github.com/go-generalize/api_gen/releases/download/$(API_GEN_VERSION)/api_gen_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf api_gen.tar.gz && rm *.tar.gz

.PHONY: bootstrap_volcago
bootstrap_volcago:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/go-generalize/volcago/cmd/volcago@$(VOLCAGO_VERSION)

.PHONY: bootstrap_volcago_with_curl
bootstrap_volcago_with_curl:
	mkdir -p ./bin
	curl -s -L -o ./bin/volcago.tar.gz https://github.com/go-generalize/volcago/releases/download/$(VOLCAGO_VERSION)/volcago_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf volcago.tar.gz && rm *.tar.gz

.PHONY: bootstrap_swag
bootstrap_swag:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/swaggo/swag/cmd/swag@v$(SWAG_VERSION)

.PHONY: bootstrap_swag_with_curl
bootstrap_swag_with_curl:
	mkdir -p ./bin
	curl -s -L -o ./bin/swag.tar.gz https://github.com/swaggo/swag/releases/download/v$(SWAG_VERSION)/swag_$(SWAG_VERSION)_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf swag.tar.gz && rm *.tar.gz

.PHONY: bootstrap_golangci_lint
bootstrap_golangci_lint:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VERSION)

.PHONY: bootstrap_golangci_lint_with_curl
bootstrap_golangci_lint_with_curl:
	mkdir -p ./bin
	curl -s -L -o ./bin/golangci-lint.tar.gz https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS_NAME)-$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf golangci-lint.tar.gz && mv golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS_NAME)-$(MACHINE_TYPE)/golangci-lint golangci-lint && rm -rf golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS_NAME)-$(MACHINE_TYPE) *.tar.gz

.PHONY: bootstrap_mockgen
bootstrap_mockgen:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/golang/mock/mockgen@v$(MOCKGEN_VERSION)

.PHONY: bootstrap_mockgen_with_curl
bootstrap_mockgen_with_curl:
	mkdir -p ./bin
	curl -s -L -o ./bin/mockgen.tar.gz https://github.com/golang/mock/releases/download/v$(MOCKGEN_VERSION)/mock_$(MOCKGEN_VERSION)_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf mockgen.tar.gz && mv mock_$(MOCKGEN_VERSION)_$(OS_NAME)_$(MACHINE_TYPE)/mockgen ./mockgen && rm -rf mock_$(MOCKGEN_VERSION)_$(OS_NAME)_$(MACHINE_TYPE) *.tar.gz

.PHONY: server_generate
server_generate:
	cd ./back && ../bin/api_gen server -p server -o server ./interfaces

.PHONY: client_generate
client_generate:
	mkdir -p ./front/lib/api
	cd ./front/lib/api && ../../../bin/api_gen client ts ../../../back/interfaces

.PHONY: go_generate
go_generate:
	go generate ./...

.PHONY: generate
generate: server_generate client_generate go_generate

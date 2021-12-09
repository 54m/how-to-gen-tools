API_GEN_VERSION := v2.4.0
VOLCAGO_VERSION := v1.0.0
SWAG_VERSION := 1.7.6

OS_NAME := `echo $(shell uname -s) | tr A-Z a-z`
MACHINE_TYPE := $(shell uname -m)

.PHONY: bootstrap
bootstrap: bootstrap_api_gen bootstrap_volcago

.PHONY: bootstrap_with_curl
bootstrap_with_curl: bootstrap_api_gen_with_curl bootstrap_volcago_with_curl

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
	mkdir -p bin
	GOBIN=${PWD}/bin go install github.com/go-generalize/volcago/cmd/volcago@$(VOLCAGO_VERSION)

.PHONY: bootstrap_volcago_with_curl
bootstrap_volcago_with_curl:
	mkdir -p bin
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

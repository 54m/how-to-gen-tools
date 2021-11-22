API_GEN_VERSION := v2.3.1
FIRESTORE_REPO_VERSION := v1.10.0

OS_NAME := `echo $(shell uname -s) | tr A-Z a-z`
MACHINE_TYPE := $(shell uname -m)

.PHONY: bootstrap
bootstrap: bootstrap_api_gen bootstrap_firestore_repo

.PHONY: bootstrap_api_gen
bootstrap_api_gen:
	mkdir -p ./bin
	curl -s -L -o ./bin/api_gen.tar.gz https://github.com/go-generalize/api_gen/releases/download/$(API_GEN_VERSION)/api_gen_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf api_gen.tar.gz && rm *.tar.gz

.PHONY: bootstrap_firestore_repo
bootstrap_firestore_repo:
	mkdir -p bin
	curl -s -L -o ./bin/firestore-repo.tar.gz https://github.com/go-generalize/firestore-repo/releases/download/$(FIRESTORE_REPO_VERSION)/firestore-repo_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf firestore-repo.tar.gz && rm *.tar.gz

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

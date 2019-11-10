SHELL = bash

.ONESHELL:

.PHONY: local_test_env test test_integration
.SHELLFLAGS: -ec

local_test_env:
	mkdir -p .local_test_env/datastore
	gcloud beta emulators datastore start --data-dir=.local_test_env/datastore/

generate:
	export PATH=$$PATH:~/go/bin
	go generate ./...

test_unit:
	go test ./... -test.short

test_all:
	$(shell gcloud beta emulators datastore env-init --data-dir=.local_test_env/datastore)
	go test ./...

local_server:
	$(shell gcloud beta emulators datastore env-init --data-dir=.local_test_env/datastore)
	go run cmd/server/server.go

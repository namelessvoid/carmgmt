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

test:
	go test ./... -test.short

test_integration:
	$(shell gcloud beta emulators datastore env-init --data-dir=.local_test_env/datastore)
	go test ./...

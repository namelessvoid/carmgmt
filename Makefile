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

test_unit_coverage:
	go test ./... -test.short -coverprofile=coverage.out
	go tool cover -html=coverage.out

test_all:
	$(shell gcloud beta emulators datastore env-init --data-dir=.local_test_env/datastore)
	go test ./...

test_all_coverage:
	$(shell gcloud beta emulators datastore env-init --data-dir=.local_test_env/datastore)
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

local_server: export AUTH_ISSUER = "https://dev-fleetmgmt.eu.auth0.com/"
local_server: export AUTH_CLIENTID = "Fleet Management - Local"
local_server:
	$(shell gcloud beta emulators datastore env-init --data-dir=.local_test_env/datastore)
	go run cmd/server/server.go cmd/server/configuration.go

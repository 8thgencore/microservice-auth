MAKEFLAGS += --no-print-directory

# Check if the ENV variable is set
ifneq ($(ENV),)
	include .env.$(ENV)
endif
CONFIG=.env.$(ENV)

# Set the path to the local bin directory
LOCAL_BIN:=$(CURDIR)/bin

# Migration settings
LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost \
	port=$(POSTGRES_PORT) \
	dbname=$(POSTGRES_DB) \
	user=$(POSTGRES_USER) \
	password=$(POSTGRES_PASSWORD) \
	sslmode=disable"

# Tests settings
TESTS_PATH=./internal/service/...,./internal/api/...
TESTS_ATTEMPTS=5
TESTS_COVERAGE_FILE=coverage.out

# TLS settings
TLS_PATH=tls

# Warning message to ensure correct environment export
.PHONY: check-env
check-env:
ifndef ENV
	$(error "Please run 'export ENV=dev|stage|prod' and 'export $$(xargs < .env.$(ENV))' before executing make")
else 
	@echo "[INFO] Running make with environment: $(ENV)"
endif

# #################### #
# DEPENDENCIES & TOOLS #
# #################### #

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/air-verse/air@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.1.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.22.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.22.0
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
	GOBIN=$(LOCAL_BIN) go install mvdan.cc/gofumpt@latest
	GOBIN=$(LOCAL_BIN) go install github.com/yoheimuta/protolint/cmd/protolint@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
	GOBIN=$(LOCAL_BIN) go install mvdan.cc/gofumpt@latest
	GOBIN=$(LOCAL_BIN) go install github.com/yoheimuta/protolint/cmd/protolint@latest
	GOBIN=$(LOCAL_BIN) go install github.com/bojand/ghz/cmd/ghz@v0.120.0

# Fetch Go dependencies
get-protoc-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Linting
lint:
	GOBIN=$(LOCAL_BIN) bin/golangci-lint run ./... --config .golangci.pipeline.yaml

# Formating
format:
	GOBIN=$(LOCAL_BIN) bin/gofumpt -l -w .

# Protolint
protolint:
	GOBIN=$(LOCAL_BIN) bin/protolint lint api/*   

# ############### #
# CODE GENERATION #
# ############### #
generate-api:
	make generate-user-api
	make generate-auth-api
	make generate-access-api

generate-user-api:
	mkdir -p pkg/user/v1 pkg/swagger
	protoc --proto_path api/user/v1 --proto_path vendor.protogen \
	--go_out=pkg/user/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go-grpc_out=pkg/user/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/user/v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
	--validate_out lang=go:pkg/user/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
	api/user/v1/user.proto

generate-auth-api:
	mkdir -p pkg/auth/v1
	protoc --proto_path api/auth/v1 --proto_path vendor.protogen \
	--go_out=pkg/auth/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go-grpc_out=pkg/auth/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/auth/v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--validate_out lang=go:pkg/auth/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
	api/auth/v1/auth.proto

generate-access-api:
	mkdir -p pkg/access/v1
	protoc --proto_path api/access/v1 --proto_path api/user/v1 --proto_path vendor.protogen \
	--go_out=pkg/access/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go-grpc_out=pkg/access/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/access/v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--validate_out lang=go:pkg/access/v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
	api/access/v1/access.proto

# Update swagger environment variables
update-swagger: check-env
	sed -i -e 's/{HTTP_HOST}/$(HTTP_HOST)/g' pkg/swagger/api.swagger.json
	sed -i -e 's/{HTTP_PORT}/$(HTTP_PORT)/g' pkg/swagger/api.swagger.json

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone --depth=1 https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone --depth=1 https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone --depth=1 https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi

generate-mocks:
	go generate ./internal/repository
	go generate ./internal/service
	go generate ./internal/tokens

# Generation of a CA (Certification Authority)
generate-cert-ca: 
	mkdir -p tls
	openssl genpkey -algorithm ed25519 -out $(TLS_PATH)/ca.key
	openssl req -new -x509 -key $(TLS_PATH)/ca.key -out $(TLS_PATH)/ca.crt -days 365 -sha256 -subj "/CN=My CA"

# Generating a CA-signed certificate
generate-cert-auth: $(TLS_PATH)/ca.key $(TLS_PATH)/ca.crt
	openssl genpkey -algorithm ed25519 -out $(TLS_PATH)/auth.key
	openssl req -new -key $(TLS_PATH)/auth.key -config openssl.cnf -out $(TLS_PATH)/auth.csr
	openssl x509 -req -in $(TLS_PATH)/auth.csr -CA $(TLS_PATH)/ca.crt -CAkey $(TLS_PATH)/ca.key \
	-extfile openssl.cnf -extensions req_ext \
	-out $(TLS_PATH)/auth.crt -days 365 -sha256
	rm -rf $(TLS_PATH)/auth.csr

# ##### #
# TESTS #
# ##### #

test:
	go clean -testcache
	-go test ./... -v -covermode count -coverpkg=$(TESTS_PATH) -count $(TESTS_ATTEMPTS)

test-coverage:
	go clean -testcache
	-go test ./... -v -coverprofile=$(TESTS_COVERAGE_FILE).tmp -covermode count -coverpkg=$(TESTS_PATH) -count $(TESTS_ATTEMPTS)
	grep -v "mocks/" $(TESTS_COVERAGE_FILE).tmp > $(TESTS_COVERAGE_FILE)
	rm $(TESTS_COVERAGE_FILE).tmp
	go tool cover -html=$(TESTS_COVERAGE_FILE) -o coverage.html
	go tool cover -func=$(TESTS_COVERAGE_FILE) | grep "total"

load-test: check-env
	$(LOCAL_BIN)/ghz \
		--proto api/user/v1/user.proto \
		--import-paths=vendor.protogen/ \
		--cacert=tls/ca.crt \
		--call user_v1.UserV1.Get \
		--data '{"id": "9f80dfbf-2ae2-4a9c-a490-3921ca7f2b65"}' \
		--rps 100 \
		--total 3000 \
		${GRPC_HOST}:${GRPC_PORT}

load-test-error: check-env
	$(LOCAL_BIN)/ghz \
		--proto api/user_v1/user.proto \
		--import-paths=vendor.protogen/ \
		--cacert=tls/ca.crt \
		--call user_v1.UserV1.Get \
		--data '{"id": "9f80dfbf-2ae2-4a9c-a490-3921ca7f2b65"}' \
		--rps 100 \
		--total 3000 \
		${GRPC_HOST}:${GRPC_PORT}

# ##### #
# BUILD #
# ##### #

build-app:
	GOOS=linux GOARCH=amd64 go build -o $(LOCAL_BIN)/main cmd/user/main.go

docker-net:
	docker network create -d bridge service-net

docker-build: docker-build-app docker-build-migrator

docker-build-app: check-env
	docker buildx build --no-cache --platform linux/amd64 -t auth:${APP_IMAGE_TAG} --build-arg="ENV=${ENV}" --build-arg="CONFIG=${CONFIG}" .

docker-build-migrator: check-env
	docker buildx build --no-cache --platform linux/amd64 -t migrator-auth:${MIGRATOR_IMAGE_TAG} -f migrator.Dockerfile --build-arg="ENV=${ENV}" .

# ###### #
# DEPLOY #
# ###### #

docker-deploy: check-env docker-build
	docker compose --env-file=.env.$(ENV) up -d

# ######### #
# MIGRATION #
# ######### #

local-migration-status: check-env
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up: check-env
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down: check-env
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

# #### #
# STOP #
# #### #

docker-stop: check-env
	docker compose --env-file=.env.$(ENV) down

# ########### #
# DEVELOPMENT #
# ########### #

dev:
	$(LOCAL_BIN)/air

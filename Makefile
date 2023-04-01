SHELL = /bin/bash

# ------------------------------------------------
# Docker compose
# ------------------------------------------------
up:
	docker compose -f ./docker-compose.yml up
	
down:
	docker compose -f ./docker-compose.yml down

# ------------------------------------------------
# CLI commands
# ------------------------------------------------
# - DOT
dot:
	go run main.go host-group dot

dot-file:
	go run main.go host-group dot --file build/render.dot.txt

dot-color:
	go run main.go host-group dot --color

dot-file-color:
	go run main.go host-group dot --file build/render.dot.txt --color

svg:
	go run main.go host-group dot | dot -Tsvg > build/output.svg

svg-color:
	go run main.go host-group dot --color | dot -Tsvg > build/output.svg

# - JSON
json:
	go run main.go host-group json

json-file:
	go run main.go host-group json --file build/render.json

# - SHELL
shell:
	go run main.go host-group shell

shell-file:
	go run main.go host-group shell --file build/render.shell.txt

# - FIXTURES
# export-services:
# 	go run fixtures/service.go --url ${ZABBIX_URL} --user ${ZABBIX_USER} --password ${ZABBIX_PWD} --action export

import-services:
	go run fixtures/service.go --url http://localhost:4444/api_jsonrpc.php --user ${ZABBIX_USER} --password ${ZABBIX_PWD} --action import

# - HELPER
help:
	go run main.go --help

# ------------------------------------------------
# Tests
# ------------------------------------------------
test:
	@echo "Running container stack..."
	docker compose -f ./docker-compose.test.yml up -d
	@TIMER=40; \
	echo "Waiting $$TIMER\\s for Zabbix server to initialize"; \
	i=1; \
	while [[ $$i -ne $$TIMER ]]; \
	do \
		echo "$$i / $$TIMER"; \
		sleep 1; \
		i=$$((i+1)); \
	done;
	@echo "Running test..."
	go test ./...
	@echo "Destroying container stack"
	docker compose -f ./docker-compose.test.yml down

coverage:
	@echo "Running container stack..."
	docker compose -f ./docker-compose.test.yml up -d
	@TIMER=40; \
	echo "Waiting $$TIMER\\s for Zabbix server to initialize"; \
	i=1; \
	while [[ $$i -ne $$TIMER ]]; \
	do \
		echo "$$i / $$TIMER"; \
		sleep 1; \
		i=$$((i+1)); \
	done;
	@echo "Running test..."
	go test -coverprofile=coverage.out ./...
	@echo "Formatting coverage report to HTML..."
	go tool cover -html=coverage.out -o=coverage.html
	@echo "Destroying container stack"
	docker compose -f ./docker-compose.test.yml down

down-test:
	docker compose -f ./docker-compose.test.yml down

clean-test-cache:
	go clean -testcache

rm-test-files:
	rm internal/render/test_render.*
	rm internal/app/test_render.*

# ------------------------------------------------
# Build
# ------------------------------------------------
build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/zabbix-tree-cli_linux_amd64 --ldflags '-extldflags "-static"'

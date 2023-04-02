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
# - HostGroup
# -- DOT
hostgroup-dot:
	go run main.go host-group dot

hostgroup-dot-file:
	go run main.go host-group dot --file build/render.dot.txt

hostgroup-dot-color:
	go run main.go host-group dot --color

hostgroup-dot-file-color:
	go run main.go host-group dot --file build/render.dot.txt --color

hostgroup-svg:
	go run main.go host-group dot | dot -Tsvg > build/output.svg

hostgroup-svg-color:
	go run main.go host-group dot --color | dot -Tsvg > build/output-color.svg

# -- JSON
hostgroup-json:
	go run main.go host-group json

hostgroup-json-file:
	go run main.go host-group json --file build/render.json

# -- SHELL
hostgroup-shell:
	go run main.go host-group shell

hostgroup-shell-file:
	go run main.go host-group shell --file build/render.shell.txt

# - Service
# -- DOT
service-dot:
	go run main.go service dot

service-dot-file:
	go run main.go service dot --file build/render.dot.txt

service-dot-color:
	go run main.go service dot --color

service-dot-file-color:
	go run main.go service dot --file build/render.dot.txt --color

service-svg:
	go run main.go service dot | dot -Tsvg > build/output.svg

service-svg-color:
	go run main.go service dot --color | dot -Tsvg > build/output-color.svg

# -- JSON
service-json:
	go run main.go service json

service-json-file:
	go run main.go service json --file build/render.json

# -- SHELL
service-shell:
	go run main.go service shell

service-shell-file:
	go run main.go service shell --file build/render.shell.txt

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

#!/bin/bash

install-tools:
	go install github.com/pressly/goose/v3/cmd/goose@latest

create-migration:
	goose -dir files/sql create $(FILENAME) sql

migrate:
	goose -dir files/sql mysql "root:@/oms?parseTime=true" up

setup-env:
	docker compose up -d

sync-dep:
	go mod tidy
	go mod vendor

init-config:
	cp files/config/local.yaml.example files/config/local.yaml
THIS_FILE := $(lastword $(MAKEFILE_LIST))
.PHONY: help build up start down destroy stop restart logs logs-api ps login-timescale login-api db-shell
help:
		make -pRrq  -f $(THIS_FILE) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
build:
		docker-compose build $(c)
up:
		docker-compose up -d $(c)
start:
		docker-compose start $(c)
down:
		docker-compose down $(c)
destroy:
		docker-compose down -v $(c)
stop:
		docker-compose stop $(c)
restart:
		docker-compose stop $(c)
		docker-compose up -d $(c)
logs:
		docker-compose logs --tail=100 -f $(c)
logs-bot:
		docker-compose logs --tail=100 -f bot
ps:
		docker-compose ps
ssh-bot:
		docker-compose exec bot /bin/bash
generate:
		go generate ./...
db-shell:
		docker-compose exec db psql postgres://local-dev@db/tibiago?sslmode=disable
db-migrate:
		docker-compose exec bot /bin/bash -c "go run main.go migrate"
db-seed:
		docker-compose exec bot /bin/bash -c "go run main.go seed"
db-reset:
		docker-compose exec db psql postgres://local-dev@db/tibiago?sslmode=disable -c 'DROP SCHEMA public CASCADE; CREATE SCHEMA public;'
		docker-compose exec bot /bin/bash -c "go run main.go migrate"
		docker-compose exec bot /bin/bash -c "go run main.go seed"
install:
	npm install -g dynamodb-admin
	go install cmd/telescope/telescope.go

build:
	GOARCH=amd64 GOOS=linux go build -o main cmd/telescope/telescope.go

run:
	sam local invoke --docker-network telescope_local -t ./template.yml ScrapeData

setup:
	# Set ip alias for local host, this is required because
	# SAM cannot connect to localhost by default 
	sudo ifconfig lo0 alias 172.16.123.1 
	# Build and run docker-compose in background
	# This will setup network and dependecies that
	# SAM will access
	docker-compose build
	docker-compose up -d
	# Startup dynamodb-admin for debugging purposes
	DYNAMO_ENDPOINT=http://localhost:8000 dynamodb-admin

teardown: 
	docker-compose down
install:
	npm install -g dynamodb-admin
	go install cmd/telescope/telescope.go
	go install cmd/lens/lens.go

build:
	GOARCH=amd64 GOOS=linux go build -o bin/telescope cmd/telescope/telescope.go
	GOARCH=amd64 GOOS=linux go build -o bin/lens cmd/lens/lens.go

run:
	sam local start-lambda --docker-network telescope_local -t ./template.yml 

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
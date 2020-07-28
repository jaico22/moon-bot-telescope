build:
	go build -o main cmd/telescope/telescope.go

run:
	./main

compose:
	docker-compose build
	docker-compose up
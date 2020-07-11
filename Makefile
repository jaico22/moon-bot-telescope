build:
	go build -o bin/telescope cmd/telescope/telescope.go

run:
	./bin/telescope

compose:
	docker-compose build
	docker-compose up
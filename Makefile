build:
	GOARCH=amd64 GOOS=linux go build -o main cmd/telescope/telescope.go

run:
	sam local invoke --docker-network telescope_local -t ./template.yml ScrapeData

setup:
	sudo ifconfig lo0 alias 172.16.123.1
	docker-compose build
	docker-compose up
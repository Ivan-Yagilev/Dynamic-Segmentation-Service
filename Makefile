build:
	docker build -t back:1 .
	docker-compose build
run: build
	docker-compose up -d
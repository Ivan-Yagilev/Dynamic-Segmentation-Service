build:
	docker build -t back:1 .
	docker-compose build
run: build
	docker-compose up -d
	migrate -path ./schema -database 'postgres://bruh:bruh@0.0.0.0:5432/db?sslmode=disable' up
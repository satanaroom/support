.PHONY: run migrate

run:
	@./support

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

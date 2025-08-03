.PHONY: build
build:
	@docker build -t pg-uuid-ulid:latest .

.PHONY: start
start:
	@docker compose up -d

.PHONY: stop
stop:
	@docker compose down

.PHONY: analyze
analyze:
	@cat analyze.sql | docker exec -i pg17 psql -U pg -d mydb

.PHONY: clean
clean:
	@cat clean.sql | docker exec -i pg17 psql -U pg -d mydb

.PHONY: size
size:
	@cat size.sql | docker exec -i pg17 psql -U pg -d mydb
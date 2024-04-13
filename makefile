DC = docker-compose -f docker-compose.yml
# DCT = docker-compose -f docker-compose.test.yml
MIGRATE = ./run_migrations.sh

dev:
	$(DC) up --build

env_local:
	cp .env.local .env

migrate:
	$(MIGRATE) up development
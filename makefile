DC = docker-compose -f docker-compose.yml
# DCT = docker-compose -f docker-compose.test.yml
dev:
	$(DC) up --build

env_local:
	cp .env.local .env
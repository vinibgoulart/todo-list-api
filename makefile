DC = docker-compose -f docker-compose.yml
DCT = docker-compose -f docker-compose.test.yml
dev:
	$(DC) up --build

test: 
	$(DCT) up --build --abort-on-container-exit

env_local:
	cp .env.local .env
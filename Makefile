run-db: remove-db
	docker run --name cinema-db -e POSTGRES_USER=cinema-user \
	-e POSTGRES_PASSWORD=cinema-password -e POSTGRES_DB=cinema-db \
	-p 5432:5432 -d postgres

remove-db:
	docker rm -f cinema-db

run-memphis:
	curl -s https://memphisdev.github.io/memphis-docker/docker-compose.yml \
	-o docker-compose.yml && docker compose -f docker-compose.yml -p memphis up -d

remove-memphis:
	docker rm -f memphis-mongo-1 memphis-memphis-cluster-1

svc-booking:
	source ./.env && go run ./cmd/booking-service

svc-cinema:
	source ./.env && go run ./cmd/cinema-service

svc-payment:
	source ./.env && go run ./cmd/payment-service
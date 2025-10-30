setup:
	docker compose up -d db
	docker compose exec -it db psql -U pguser -c "CREATE DATABASE calendar_dev"

run:
	docker compose run --rm --service-ports app


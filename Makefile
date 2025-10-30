.PHONY:
run: setup
	docker compose run --rm --service-ports app

.PHONY:
setup: build
	docker compose up -d db
	docker compose exec -it db \
		psql -U pguser -d postgres -tc "SELECT 1 FROM pg_database WHERE datname='calendar_prod';" | grep -q 1 \
		|| docker compose exec -it db psql -U pguser -d postgres -c "CREATE DATABASE calendar_prod;"

.PHONY:
build:
	docker compose build

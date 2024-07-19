DB_URL="postgres://postgres:3f4f2770c42a8efddc80e61da8d7c9f71cfe2eb03ff6040542ad24a42192731f4af7e875ca5fc8736240@127.0.0.1:55432/postgres?sslmode=disable"

.PHONY: migrate

reset:
	docker compose down --volumes --remove-orphans
	docker compose up -d

migrate:
	docker run --rm -it --network=host -v "./app/database:/db" ghcr.io/amacneil/dbmate -u "$(DB_URL)" -no-dump-schema migrate

run_scenario_1:
	./tests-damien.sh
	./tests-qazar.sh

all:
	make reset
	sleep 4
	make migrate
	echo 'waiting for api to boot'
	sleep 20
	make run_scenario_1
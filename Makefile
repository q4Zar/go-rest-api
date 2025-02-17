DB_URL="postgres://postgres:3f4f2770c42a8efddc80e61da8d7c9f71cfe2eb03ff6040542ad24a42192731f4af7e875ca5fc8736240@127.0.0.1:55432/postgres?sslmode=disable"

.PHONY: migrate

reset:
	docker compose down --volumes --remove-orphans
	docker compose up postgres -d

migrate:
	docker run --rm -it --network=host -v "./app/database:/db" ghcr.io/amacneil/dbmate -u "$(DB_URL)" -no-dump-schema migrate

run_go_api:
	docker compose up go-api -d

run_curl_tests:
	docker compose up curl-tests

all:
	make reset
	echo "wait postgres to be up"
	sleep 4
	make migrate
	# echo "wait for api to boot"
	# make run_go_api
	# docker compose logs -f
	# make run_curl_tests
# run-migrations

`docker run --rm -it --network=host -v "$(pwd)/app/database:/db" ghcr.io/amacneil/dbmate -u postgres://postgres:3f4f2770c42a8efddc80e61da8d7c9f71cfe2eb03ff6040542ad24a42192731f4af7e875ca5fc8736240@136.243.124.144:5432/postgres?sslmode=disable -no-dump-schema migrate`
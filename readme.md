# github.com/q4Zar/go-rest-api

- Overall i decided to use goyave because i like the way of building with it and it's integration with postgresql with native GORM i think it's super pertinent ;)
    - I never used it before starting the test i did a lil research and found it
    - So i've added some learning curve but that was cool and stimulated my brain

- I switched basic auth to JWT because it's allowing me to recognize the owner of any ressources
    - If we want to build a good product with a long therm vision that's the way
    - It's not productive to build something with basicAuth in this case because i can't recognise user directly, i'm doing it with JWT
    - Hope you'll see the benefits :)

# step 1 

- creating users in DB
- using JWT by login users

# step 2
- currency is the balance (could done another table but i'll use enum and unicity)
- set unicity owner-id and currency-name 


## run-migrations

`docker run --rm -it --network=host -v "$(pwd)/app/database:/db" ghcr.io/amacneil/dbmate -u postgres://postgres:3f4f2770c42a8efddc80e61da8d7c9f71cfe2eb03ff6040542ad24a42192731f4af7e875ca5fc8736240@136.243.124.144:5432/postgres?sslmode=disable -no-dump-schema migrate`

## curl
- create-user : `curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/users`
- login-user : `curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/login`

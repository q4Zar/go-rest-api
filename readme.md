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
- ⚠️currency index jump when unicity constraint error


# run

## run-migrations
- #1 make migrate

## curl
- #2 create-user : `curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/users`
- #3 login-user : `curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/login`
- #4 create-currency-balance-by-user: 
    - EUR `curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjEwMzg4MzYsIm5iZiI6MTcyMTAzNTgzNiwic3ViIjoiZGFtcyJ9.cgMhkTL5zAiLCL4WzAztTI_O4qu4qXLZ14u65UnRtM0" -d '{"name": "EUR", "amount": 1000}' -H "Content-Type: application/json" http://127.0.0.1:8080/currencies`
    - USD `curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjEwMzg4MzYsIm5iZiI6MTcyMTAzNTgzNiwic3ViIjoiZGFtcyJ9.cgMhkTL5zAiLCL4WzAztTI_O4qu4qXLZ14u65UnRtM0" -d '{"name": "USD", "amount": 1000}' -H "Content-Type: application/json" http://127.0.0.1:8080/currencies`
- #5 list-currencies-from-user : `curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjEwMzg4MzYsIm5iZiI6MTcyMTAzNTgzNiwic3ViIjoiZGFtcyJ9.cgMhkTL5zAiLCL4WzAztTI_O4qu4qXLZ14u65UnRtM0" -H "Content-Type: application/json" http://127.0.0.1:8080/currency`
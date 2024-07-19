# github.com/q4Zar/go-rest-api


## State of Mind

- Overall, I decided to use Goyave because I like the way it is built and its integration with PostgreSQL through native GORM. I think it's super pertinent ;)
    - I had never used it before starting the test. I did a little research and discovered it.
    - So, I encountered a bit of a learning curve, but it was cool and stimulating for my brain.

- I switched from basic auth to JWT because it allows me to recognize the owner of any resource.
    - If we want to build a good product with a long-term vision, this is the way to go.
    - I hope you'll see the benefits :)

---

## Tiny Docs
### Create User
`curl -s -X POST -d '{"username": "qazar", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" "$go_api/users"`
### Login User
`token=$(curl -s -X POST -d '{"username": "qazar", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" "$go_api/login" | jq -r '.token')`
### Create Asset
#### EUR
`asset_euro_success=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "EUR", "balance" : 10000}' -H "Content-Type: application/json" "$go_api/assets")`
#### USD
`asset_dollar_success=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "USD", "balance" : 10000}' -H "Content-Type: application/json" "$go_api/assets")`

### Create Order

---

## Running (2 Terminals for better readability)
**I automatized everything but there is some output defaults so it's easier to just trigger it manually)

### Terminal 1
- sudo make all (run postgres, migrate, boot api)
**Wait for**
```
go-api-1    |
go-api-1    |  INFO  2024/07/19 12:28:02.40583 (/app/main.go:39)
go-api-1    | Registering hooks
go-api-1    |
go-api-1    |  INFO  2024/07/19 12:28:02.406913 (/app/main.go:62)
go-api-1    | Registering services
go-api-1    |
go-api-1    |  INFO  2024/07/19 12:28:02.406986 (/app/main.go:52)
go-api-1    | Registering routes
go-api-1    |
go-api-1    |  INFO  2024/07/19 12:28:02.407393 (/app/main.go:43)
go-api-1    | Server is listening
go-api-1    | host: 0.0.0.0:8080
```
[![asciicast](https://asciinema.org/a/3KcuGWwv3CRJ2skpI6q6SRNyQ.svg)](https://asciinema.org/a/3KcuGWwv3CRJ2skpI6q6SRNyQ)

### Terminal 2
- cd curl-tests
- ./tests-damien.sh http://localhost:8080
- ./tests-qazar.sh http://localhost:8080
[![asciicast](https://asciinema.org/a/WTw7DkXugJ6xVXHWBBFvUFIp7.svg)](https://asciinema.org/a/WTw7DkXugJ6xVXHWBBFvUFIp7)

---

## ScreenShots

### Users
![Users](.screenshots/users.png)
### Assets
![Assets](.screenshots/assets.png)
### Orders
![Orders](.screenshots/orders.png)

---

## Quick Optimization
- Check Balance Before Creating Order (i started but i'm running out of time so i just stick to the initial subject)
- Go Client to query API
- GET /assets /orders only returns Owner data like PATCH & DELETE
- ...

---

## Outro
- Hope you like i enjoyed doing it, asynchronously some hours there and here for a week
- I think it's a great implementation to start a good system
- I could have done an easie micro-framework and sqlx with RAW SQL but i wanted to try something else and i found goyave really pertinent and recent espacially with the files segmentations
- Don't hesitate if you have any questions 


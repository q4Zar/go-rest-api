# Back-end Engineer Golang - Assignment

## Introduction

Welcome!

This exercise is part of the application process of the back-end engineer role.

Although the product aspect for this exercise is not directly linked to the Tranched products, we want to highlight challenges similar to the ones you would encounter working at Tranched.

Through assessing your solution, we hope to get a better picture of how you solve problems and whether you'd be a good fit.

There is no time limit for the completiong of the task, but we expect you to spend around 3 hours on it.
Do not hesitate to ask questions by emailing us directly.

## Assignment

The idea is to create a REST API that allows users to deposit, withdraw assets (currencies) and trade them with other users.

### Ground rules

- We would like to see the solution implemented in Golang because it is what we use for our backend systems at Tranched.
- Application data to be stored in a PostgreSQL database (a local docker container would be ideal).
- You can use any utility library and/or framework that you see fit, but remember that we want to evaluate how you design code, not how many libraries you know.
- Many of the design requirements presented here are naive and greatly simplified, the point is to focus on an end-to-end story that makes sense and allow us to assess your design choices.
- Try to write clean, maintainable and understandable code. We don't expect the code to be unit tested, but you can write some if that helps you structure your code better!

### Deliverables

- Source code for an HTTP REST API back-end application.
- Instructions on how to run the application locally.

### API & Components

#### 1. Users & authentication

- All users of the system have a username and a password, stored in the database.
- Users can authenticate each HTTP requests using the Basic Authentication scheme.
- Usernames (alice, bob for example) and passwords can be hardcoded in a separate configuration file or directly in the application source code.

---

---

#### 2. Asset Management

- Only two types of currency assets are supported so far: EUR and USD.
- Each user has a balance for each asset they own, recorded in the database.
- Set arbitrary balances for each user, at application startup.
- A user can query their assets balance at any time:
  - Endpoint: `GET /assets`
  - Returning JSON payload containing an array of simple objects containing the `asset_type` (string) and `balance` (float), for each asset they own.
  - The field `asset_type` would be either EUR or USD.

---

- WIP
  - 

---

#### 3. Orders & Trading

- A user can create limit buy or sell orders for a valid asset pair, by calling the endpoint:
  - Endpoint: `POST /orders`
  - Providing a JSON payload containing the fields `side` (string), `asset_pair` (string), `amount` (float) and a `price` (float).
  - The `side` can be either `BUY` or `SELL`.
  - A valid `asset_pair` would be `EUR-USD` in our case.
  - Order semantics for the `EUR-USD` pair:
    - The `price` represents the amount of USD required to buy one unit of EUR (1.2 USD = 1 EUR).
    - The `amount` is a quantity of `USD` to include in the order.
    - A `BUY` side for `EUR-USD` pair, for an amount of `1200.0` and price `1.2` would mean for a user to "buy 1000.0 EUR spending 1200.0 USD"
    - A `SELL` side for `EUR-USD` pair, for an amount of `1300.0` and price `1.3` would mean for a user to "sell 1000 EUR to receive 1300.0 USD"
  - The new order is stored in the database, for that user, with a `pending` status.
  - Order statuses can be `pending` or `filled`.
- Write a simplistic order matching service that takes care of matching orders as they come in through the above endpoint (no database storage here, all pending orders are stored in memory).
  - As soon as two pending orders match exactly, a few things should happen:
    - Update the status of both orders to `filled` in the database.
    - Update the assets balance for both users in the database.
    - Remove the pending orders from the order matching service's set of pending orders.
  - Example of two matching orders:
    - (`BUY`, `EUR-USD`, `1200.0`, `1.2`)
    - (`SELL`, `EUR-USD`, `1200.0`, `1.2`)
  - Naturally, after an order is executed the users should see their updated asset balances reflected when calling the `GET /assets` endpoint.
- A user can query their orders:
  - Endpoint: `GET /orders`
  - Returning JSON payload containing an array of order objects containing the fields as above (`side`, `asset_pair`, `amount`, `price`) as well as the `status` (string) field.

---

- WIP
  - 

---
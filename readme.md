# github.com/q4Zar/go-rest-api

## State of Mind

- Overall, I decided to use Goyave because I like the way it is built and its integration with PostgreSQL through native GORM. I think it's super pertinent ;)
    - I had never used it before starting the test. I did a little research and discovered it.
    - So, I encountered a bit of a learning curve, but it was cool and stimulating for my brain.

- I switched from basic auth to JWT because it allows me to recognize the owner of any resource.
    - If we want to build a good product with a long-term vision, this is the way to go.
    - I hope you'll see the benefits :)

## Running

- Two terminals (I automatized everything but it goes too fast it's better to just trigger the two curl script separately to check the outputs)

### terminal 1

- ` sudo make all`
        - run postgres
        - do migration
        - boot api

### terminal 2

- ./tests-damien.sh http://localhost:8080
- ./tests-qazar.sh http://localhost:8080

## ScreenShots

- Users
- Assets Before Orders
- Orders Pending
- Orders Filled
- Asset After Orders


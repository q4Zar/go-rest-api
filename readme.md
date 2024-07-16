# github.com/q4Zar/go-rest-api

## state-of-mind
- Overall i decided to use goyave because i like the way of building with it and it's integration with postgresql with native GORM i think it's super pertinent ;)
    - I never used it before starting the test i did a lil research and found it
    - So i've added some learning curve but that was cool and stimulated my brain

- I switched basic auth to JWT because it's allowing me to recognize the owner of any ressources
    - If we want to build a good product with a long therm vision that's the way
    - It's not productive to build something with basicAuth in this case because i can't recognise user directly, i'm doing it with JWT
    - Hope you'll see the benefits :)

---

### exercice

#### step 1 
- signup user in DB
- login user and get JWT

#### step 2
- currency is a table


#### step 3

---

## debug

### make all
```zsh
qazarcloud:go-rest-api/ (main) $ make all                                                                                                                                                                                          [11:10:56]
make reset_db
make[1]: Entering directory '/home/qazarcloud/organizations/tranched/go-rest-api'
docker compose down --volumes --remove-orphans
[+] Running 1/1
 ✔ Container go-rest-api-postgres-1  Removed                                                                                                                                                                                             0.3s
docker compose up -d
[+] Running 1/1
 ✔ Container go-rest-api-postgres-1  Started                                                                                                                                                                                             0.3s
make[1]: Leaving directory '/home/qazarcloud/organizations/tranched/go-rest-api'
sleep 4
make migrate
make[1]: Entering directory '/home/qazarcloud/organizations/tranched/go-rest-api'
docker run --rm -it --network=host -v "./app/database:/db" ghcr.io/amacneil/dbmate -u ""postgres://postgres:3f4f2770c42a8efddc80e61da8d7c9f71cfe2eb03ff6040542ad24a42192731f4af7e875ca5fc8736240@127.0.0.1:55432/postgres?sslmode=disable"" -no-dump-schema migrate
Applying: 20240714133046_create_users_table.sql
Applied: 20240714133046_create_users_table.sql in 8.023944ms
Applying: 20240714133055_create_currencies_table.sql
Applied: 20240714133055_create_currencies_table.sql in 12.534957ms
Applying: 20240714143055_create_assets_table.sql
Applied: 20240714143055_create_assets_table.sql in 12.659585ms
make[1]: Leaving directory '/home/qazarcloud/organizations/tranched/go-rest-api'
sleep 4
make run_scenario_1
make[1]: Entering directory '/home/qazarcloud/organizations/tranched/go-rest-api'
./tests.sh
creating user
login user
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   203  100   158  100    45   3434    978 --:--:-- --:--:-- --:--:--  4413
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjExMjQwOTgsIm5iZiI6MTcyMTEyMTA5OCwic3ViIjoiZGFtcyJ9.VI7wDvon-T_UcB3f_Zc4wiANX9EkTH3v3uaRH43IJdA
creates currencies
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    15    0     0  100    15      0   7500 --:--:-- --:--:-- --:--:--  7500

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    15    0     0  100    15      0  15000 --:--:-- --:--:-- --:--:-- 15000

creates assets
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    54    0     0  100    54      0  27000 --:--:-- --:--:-- --:--:-- 27000

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   187  100   133  100    54   129k  54000 --:--:-- --:--:-- --:--:--  182k
{"error":"ERROR: insert or update on table \"assets\" violates foreign key constraint \"assets_currency_id_fkey\" (SQLSTATE 23503)"}
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   172  100   118  100    54   115k  54000 --:--:-- --:--:-- --:--:--  167k
{"error":"ERROR: duplicate key value violates unique constraint \"assets_user_id_currency_id_key\" (SQLSTATE 23505)"}
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    54    0     0  100    54      0  54000 --:--:-- --:--:-- --:--:-- 54000

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   187  100   133  100    54   129k  54000 --:--:-- --:--:-- --:--:--  182k
{"error":"ERROR: insert or update on table \"assets\" violates foreign key constraint \"assets_currency_id_fkey\" (SQLSTATE 23503)"}
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   172  100   118  100    54   115k  54000 --:--:-- --:--:-- --:--:--  167k
{"error":"ERROR: duplicate key value violates unique constraint \"assets_user_id_currency_id_key\" (SQLSTATE 23505)"}
make[1]: Leaving directory '/home/qazarcloud/organizations/tranched/go-rest-api'
qazarcloud:go-rest-api/ (main) $
```

### go run .
```zsh
qazarcloud:app/ (main) $ go run .                                                                                                                                                                                                  [11:11:24]

 INFO  2024/07/16 11:11:26.608175 (/home/qazarcloud/organizations/tranched/go-rest-api/app/main.go:40)
Registering hooks

 INFO  2024/07/16 11:11:26.608248 (/home/qazarcloud/organizations/tranched/go-rest-api/app/main.go:63)
Registering services

 INFO  2024/07/16 11:11:26.608259 (/home/qazarcloud/organizations/tranched/go-rest-api/app/main.go:53)
Registering routes

 INFO  2024/07/16 11:11:26.608396 (/home/qazarcloud/organizations/tranched/go-rest-api/app/main.go:44)
Server is listening
host: 0.0.0.0:8080

 DEBUG  2024/07/16 11:11:38.605065 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:36)
[0.860ms] [rows:1] INSERT INTO "users" ("username","password","created_at","updated_at") VALUES ('dams','$2a$10$f8P5X59S9v5qVN5cCek5tek7CMzkI1VFaFoET2N3J/vlnpqk/ODWW','2024-07-16 11:11:38.604','2024-07-16 11:11:38.604') RETURNING "id"

 INFO  2024/07/16 11:11:38.624989 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/users" HTTP/1.1" 201 0 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.628662 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.772ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 INFO  2024/07/16 11:11:38.674312 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/login" HTTP/1.1" 200 158 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.677264 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.309ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 DEBUG  2024/07/16 11:11:38.679211 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/currency.go:37)
[1.865ms] [rows:1] INSERT INTO "currencies" ("created_at","updated_at","deleted_at","name") VALUES ('2024-07-16 11:11:38.677','2024-07-16 11:11:38.677',NULL,'EUR') RETURNING "id"

 INFO  2024/07/16 11:11:38.679225 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/currencies" HTTP/1.1" 201 0 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.681745 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.121ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 DEBUG  2024/07/16 11:11:38.683362 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/currency.go:37)
[1.534ms] [rows:1] INSERT INTO "currencies" ("created_at","updated_at","deleted_at","name") VALUES ('2024-07-16 11:11:38.681','2024-07-16 11:11:38.681',NULL,'USD') RETURNING "id"

 INFO  2024/07/16 11:11:38.683376 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/currencies" HTTP/1.1" 201 0 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.685822 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.111ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 DEBUG  2024/07/16 11:11:38.687743 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:55)
[1.837ms] [rows:1] INSERT INTO "assets" ("created_at","updated_at","deleted_at","amount","currency_id","user_id") VALUES ('2024-07-16 11:11:38.685','2024-07-16 11:11:38.685',NULL,12356.54897,1,1) RETURNING "id"

 INFO  2024/07/16 11:11:38.687755 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/assets" HTTP/1.1" 201 0 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.690215 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.132ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 ERROR  2024/07/16 11:11:38.690622 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:55)
ERROR: insert or update on table "assets" violates foreign key constraint "assets_currency_id_fkey" (SQLSTATE 23503)
[0.351ms] [rows:0] INSERT INTO "assets" ("created_at","updated_at","deleted_at","amount","currency_id","user_id") VALUES ('2024-07-16 11:11:38.69','2024-07-16 11:11:38.69',NULL,12356.54897,3,1) RETURNING "id"

 ERROR  2024/07/16 11:11:38.690633 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/response.go:316)
ERROR: insert or update on table "assets" violates foreign key constraint "assets_currency_id_fkey" (SQLSTATE 23503)
trace:
github.com/q4Zar/go-rest-api/database/repository.(*Asset).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:56
github.com/q4Zar/go-rest-api/service/asset.(*Service).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/service/asset/asset.go:38
github.com/q4Zar/go-rest-api/http/controller/asset.(*Controller).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/http/controller/asset/asset.go:49
goyave.dev/goyave/v5.(*validateRequestMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:202
goyave.dev/goyave/v5/middleware/parse.(*Middleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware/parse/parse.go:92
goyave.dev/goyave/v5/auth.(*Handler[...]).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/auth/authenticator.go:84
goyave.dev/goyave/v5/log.(*AccessMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:112
goyave.dev/goyave/v5.(*corsMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:237
goyave.dev/goyave/v5.(*languageMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:118
goyave.dev/goyave/v5.(*recoveryMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:90
goyave.dev/goyave/v5.(*Router).requestHandler
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:507
goyave.dev/goyave/v5.(*Router).ServeHTTP
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:280
net/http.serverHandler.ServeHTTP
        /usr/local/go/src/net/http/server.go:3142
net/http.(*conn).serve
        /usr/local/go/src/net/http/server.go:2044

 INFO  2024/07/16 11:11:38.691256 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/assets" HTTP/1.1" 500 133 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.693747 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.130ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 ERROR  2024/07/16 11:11:38.694057 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:55)
ERROR: duplicate key value violates unique constraint "assets_user_id_currency_id_key" (SQLSTATE 23505)
[0.250ms] [rows:0] INSERT INTO "assets" ("created_at","updated_at","deleted_at","amount","currency_id","user_id") VALUES ('2024-07-16 11:11:38.693','2024-07-16 11:11:38.693',NULL,12356.54897,1,1) RETURNING "id"

 ERROR  2024/07/16 11:11:38.694069 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/response.go:316)
ERROR: duplicate key value violates unique constraint "assets_user_id_currency_id_key" (SQLSTATE 23505)
trace:
github.com/q4Zar/go-rest-api/database/repository.(*Asset).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:56
github.com/q4Zar/go-rest-api/service/asset.(*Service).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/service/asset/asset.go:38
github.com/q4Zar/go-rest-api/http/controller/asset.(*Controller).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/http/controller/asset/asset.go:49
goyave.dev/goyave/v5.(*validateRequestMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:202
goyave.dev/goyave/v5/middleware/parse.(*Middleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware/parse/parse.go:92
goyave.dev/goyave/v5/auth.(*Handler[...]).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/auth/authenticator.go:84
goyave.dev/goyave/v5/log.(*AccessMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:112
goyave.dev/goyave/v5.(*corsMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:237
goyave.dev/goyave/v5.(*languageMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:118
goyave.dev/goyave/v5.(*recoveryMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:90
goyave.dev/goyave/v5.(*Router).requestHandler
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:507
goyave.dev/goyave/v5.(*Router).ServeHTTP
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:280
net/http.serverHandler.ServeHTTP
        /usr/local/go/src/net/http/server.go:3142
net/http.(*conn).serve
        /usr/local/go/src/net/http/server.go:2044

 INFO  2024/07/16 11:11:38.694108 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/assets" HTTP/1.1" 500 118 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.696524 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.096ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 DEBUG  2024/07/16 11:11:38.698121 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:55)
[1.542ms] [rows:1] INSERT INTO "assets" ("created_at","updated_at","deleted_at","amount","currency_id","user_id") VALUES ('2024-07-16 11:11:38.696','2024-07-16 11:11:38.696',NULL,12356.54897,2,1) RETURNING "id"

 INFO  2024/07/16 11:11:38.698134 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/assets" HTTP/1.1" 201 0 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.700581 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.083ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 ERROR  2024/07/16 11:11:38.700872 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:55)
ERROR: insert or update on table "assets" violates foreign key constraint "assets_currency_id_fkey" (SQLSTATE 23503)
[0.246ms] [rows:0] INSERT INTO "assets" ("created_at","updated_at","deleted_at","amount","currency_id","user_id") VALUES ('2024-07-16 11:11:38.7','2024-07-16 11:11:38.7',NULL,12356.54897,3,1) RETURNING "id"

 ERROR  2024/07/16 11:11:38.700881 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/response.go:316)
ERROR: insert or update on table "assets" violates foreign key constraint "assets_currency_id_fkey" (SQLSTATE 23503)
trace:
github.com/q4Zar/go-rest-api/database/repository.(*Asset).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:56
github.com/q4Zar/go-rest-api/service/asset.(*Service).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/service/asset/asset.go:38
github.com/q4Zar/go-rest-api/http/controller/asset.(*Controller).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/http/controller/asset/asset.go:49
goyave.dev/goyave/v5.(*validateRequestMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:202
goyave.dev/goyave/v5/middleware/parse.(*Middleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware/parse/parse.go:92
goyave.dev/goyave/v5/auth.(*Handler[...]).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/auth/authenticator.go:84
goyave.dev/goyave/v5/log.(*AccessMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:112
goyave.dev/goyave/v5.(*corsMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:237
goyave.dev/goyave/v5.(*languageMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:118
goyave.dev/goyave/v5.(*recoveryMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:90
goyave.dev/goyave/v5.(*Router).requestHandler
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:507
goyave.dev/goyave/v5.(*Router).ServeHTTP
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:280
net/http.serverHandler.ServeHTTP
        /usr/local/go/src/net/http/server.go:3142
net/http.(*conn).serve
        /usr/local/go/src/net/http/server.go:2044

 INFO  2024/07/16 11:11:38.700926 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/assets" HTTP/1.1" 500 133 "" "curl/7.61.1"

 DEBUG  2024/07/16 11:11:38.703388 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/user.go:31)
[0.083ms] [rows:1] SELECT * FROM "users" WHERE "username" = 'dams' ORDER BY "users"."id" LIMIT 1

 ERROR  2024/07/16 11:11:38.703623 (/home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:55)
ERROR: duplicate key value violates unique constraint "assets_user_id_currency_id_key" (SQLSTATE 23505)
[0.189ms] [rows:0] INSERT INTO "assets" ("created_at","updated_at","deleted_at","amount","currency_id","user_id") VALUES ('2024-07-16 11:11:38.703','2024-07-16 11:11:38.703',NULL,12356.54897,2,1) RETURNING "id"

 ERROR  2024/07/16 11:11:38.703632 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/response.go:316)
ERROR: duplicate key value violates unique constraint "assets_user_id_currency_id_key" (SQLSTATE 23505)
trace:
github.com/q4Zar/go-rest-api/database/repository.(*Asset).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/database/repository/asset.go:56
github.com/q4Zar/go-rest-api/service/asset.(*Service).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/service/asset/asset.go:38
github.com/q4Zar/go-rest-api/http/controller/asset.(*Controller).Create
        /home/qazarcloud/organizations/tranched/go-rest-api/app/http/controller/asset/asset.go:49
goyave.dev/goyave/v5.(*validateRequestMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:202
goyave.dev/goyave/v5/middleware/parse.(*Middleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware/parse/parse.go:92
goyave.dev/goyave/v5/auth.(*Handler[...]).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/auth/authenticator.go:84
goyave.dev/goyave/v5/log.(*AccessMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:112
goyave.dev/goyave/v5.(*corsMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:237
goyave.dev/goyave/v5.(*languageMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:118
goyave.dev/goyave/v5.(*recoveryMiddleware).Handle.func1
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/middleware.go:90
goyave.dev/goyave/v5.(*Router).requestHandler
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:507
goyave.dev/goyave/v5.(*Router).ServeHTTP
        /home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/router.go:280
net/http.serverHandler.ServeHTTP
        /usr/local/go/src/net/http/server.go:3142
net/http.(*conn).serve
        /usr/local/go/src/net/http/server.go:2044

 INFO  2024/07/16 11:11:38.703669 (/home/qazarcloud/go/pkg/mod/goyave.dev/goyave/v5@v5.1.1/log/log.go:88)
127.0.0.1 - - [16/Jul/2024:11:11:38 +0200] "POST "/assets" HTTP/1.1" 500 118 "" "curl/7.61.1"
```
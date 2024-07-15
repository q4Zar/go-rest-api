token=$(curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/login | jq -r '.token')
echo $token

# currencies=$(curl -H "Authorization: Bearer $token" -H "Content-Type: application/json" http://127.0.0.1:8080/currencies)
# echo $currencies

# curl -X POST -H "Authorization: Bearer $token" -d '{"name": "EUR", "amount": 1000}' -H "Content-Type: application/json" http://127.0.0.1:8080/currencies

curl -H "Authorization: Bearer $token" -H "Content-Type: application/json" http://127.0.0.1:8080/users/1
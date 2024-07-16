# 0
echo 'creating user'
curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/users

# 1
echo 'login user'

token=$(curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/login | jq -r '.token')
echo $token

# 2
echo 'creates currencies'

currency_euro=$(curl -X POST -H "Authorization: Bearer $token" -d '{"name": "EUR"}' -H "Content-Type: application/json" http://127.0.0.1:8080/currencies)
echo $currency_euro

currency_dollar=$(curl -X POST -H "Authorization: Bearer $token" -d '{"name": "USD"}' -H "Content-Type: application/json" http://127.0.0.1:8080/currencies)
echo $currency_dollar

# 3
# echo 'creates balances'

# balance_euro=$(curl -X POST -H "Authorization: Bearer $token" -d '{"currency": "EUR", "amount" : 1000}' -H "Content-Type: application/json" http://127.0.0.1:8080/balances)
# echo $currency_euro

# balance_dollar=$(curl -X POST -H "Authorization: Bearer $token" -d '{"currency": "USD", "amount" : 1000}' -H "Content-Type: application/json" http://127.0.0.1:8080/balances)
# echo $currency_dollar

# curl -X POST -H "Authorization: Bearer $token" -d '{"name": "EUR", "amount": 1000}' -H "Content-Type: application/json" http://127.0.0.1:8080/currencies
# curl -X PATCH -H "Authorization: Bearer $token" -d '{"name": "EUR", "amount": 222000}' -H "Content-Type: application/json" http://127.0.0.1:8080/currencies/1

# curl -H "Authorization: Bearer $token" -H "Content-Type: application/json" http://127.0.0.1:8080/users/profile
# 0
echo 'creating user'
curl -s -X POST -d '{"username": "qazar", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/users

# 1
echo 'login user'

token=$(curl -s -X POST -d '{"username": "qazar", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/login | jq -r '.token')
# echo $token

# 2
echo 'creates assets'

## EUR
asset_euro_success=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "EUR", "balance" : 2000}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
# echo $asset_euro_success

### duplicate on index currency & user
# asset_euro_fails_2=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "EUR", "balance" : 22356.54897}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
# echo $asset_euro_fails_2

##USD
asset_dollar_success=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "USD", "balance" : 2000}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
# echo $asset_dollar_success

# ### duplicate on index currency & user
# asset_dollar_fails_1=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "USD", "balance" : 22356.54897}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
# echo $asset_dollar_fails_1

# 3
echo 'get user assets'

assets=$(curl -s -H "Authorization: Bearer $token" "http://127.0.0.1:8080/assets?fields=balance,asset_type,user_id")
echo "$assets" | jq '.'

# 4
echo 'creates orders'

order_buy_usdeur=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.2, "side":"SELL", "assetPair" : "USD-EUR"}' -H "Content-Type: application/json" http://127.0.0.1:8080/orders)
# echo $order_sell_usdeur

order_buy_eurusd=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.2, "side":"BUY", "assetPair" : "EUR-USD"}' -H "Content-Type: application/json" http://127.0.0.1:8080/orders)
# echo $order_buy_eurusd

order_buy_usdeur=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.2, "side":"SELL", "assetPair" : "USD-EUR"}' -H "Content-Type: application/json" http://127.0.0.1:8080/orders)
# echo $order_sell_usdeur

order_buy_eurusd=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.2, "side":"BUY", "assetPair" : "EUR-USD"}' -H "Content-Type: application/json" http://127.0.0.1:8080/orders)
# echo $order_buy_eurusd

order_buy_usdeur=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 500, "price" : 1.4, "side":"SELL", "assetPair" : "USD-EUR"}' -H "Content-Type: application/json" http://127.0.0.1:8080/orders)
# echo $order_sell_usdeur

order_buy_eurusd=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 500, "price" : 1.4, "side":"BUY", "assetPair" : "EUR-USD"}' -H "Content-Type: application/json" http://127.0.0.1:8080/orders)
# echo $order_buy_eurusd

# 5 check for orders
echo "GET /orders"
orders=$(curl -s -H "Authorization: Bearer $token" "http://127.0.0.1:8080/orders")
echo "$orders" | jq '.'

sleep 4

# 5 check for new balances

assets=$(curl -s -H "Authorization: Bearer $token" "http://127.0.0.1:8080/assets?fields=balance,asset_type,user_id")
echo "$assets" | jq '.'

sleep 2

# 5 check for orders
orders=$(curl -s -H "Authorization: Bearer $token" "http://127.0.0.1:8080/orders")
echo "$orders" | jq '.'

sleep 2

# 5 check for new balances

assets=$(curl -s -H "Authorization: Bearer $token" "http://127.0.0.1:8080/assets?fields=balance,asset_type,user_id")
echo "$assets" | jq '.'

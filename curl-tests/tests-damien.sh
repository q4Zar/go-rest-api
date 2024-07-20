#!/bin/bash

set -e

go_api="$1"

# 0
echo 'Creating user'
curl -s -X POST -d '{"username": "damien", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" "$go_api/users"

# 1
echo 'Logging in user'

token=$(curl -s -X POST -d '{"username": "damien", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" "$go_api/login" | jq -r '.token')
echo "Token: $token"

# 2
echo 'Creating assets'

## EUR
asset_euro_success=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "EUR", "balance" : 10000}' -H "Content-Type: application/json" "$go_api/assets")
echo "Asset Euro Success: $asset_euro_success"

### duplicate on index currency & user
# asset_euro_fails_2=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "EUR", "balance" : 22356.54897}' -H "Content-Type: application/json" "$go_api/assets")
# echo "Asset Euro Fails 2: $asset_euro_fails_2"

## USD
asset_dollar_success=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "USD", "balance" : 10000}' -H "Content-Type: application/json" "$go_api/assets")
echo "Asset Dollar Success: $asset_dollar_success"

# ### duplicate on index currency & user
# asset_dollar_fails_1=$(curl -s -X POST -H "Authorization: Bearer $token" -d '{"assetType": "USD", "balance" : 22356.54897}' -H "Content-Type: application/json" "$go_api/assets")
# echo "Asset Dollar Fails 1: $asset_dollar_fails_1"

# 3
echo 'Getting user assets'

assets=$(curl -s -H "Authorization: Bearer $token" "$go_api/assets?fields=balance,asset_type,user_id")
echo "$assets" | jq '.'

# 4
# echo 'Creating orders'

# curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.2, "side":"BUY", "assetPair" : "USD-EUR"}' -H "Content-Type: application/json" "$go_api/orders"

# # curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.2, "side":"SELL", "assetPair" : "USD-EUR"}' -H "Content-Type: application/json" "$go_api/orders"

# # curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.1, "side":"BUY", "assetPair" : "EUR-USD"}' -H "Content-Type: application/json" "$go_api/orders"

# curl -s -X POST -H "Authorization: Bearer $token" -d '{"amount": 1000, "price" : 1.1, "side":"SELL", "assetPair" : "EUR-USD"}' -H "Content-Type: application/json" "$go_api/orders"

# # 5 Check for orders
# echo 'Checking for orders'
# orders=$(curl -s -H "Authorization: Bearer $token" "$go_api/orders")
# echo "$orders" | jq '.'

# sleep 4

# # 5 Check for new balances
# echo 'Checking for new balances'
# assets=$(curl -s -H "Authorization: Bearer $token" "$go_api/assets?fields=balance,asset_type,user_id")
# echo "$assets" | jq '.'

# sleep 2

# # 5 Check for orders
# echo 'Checking for orders again'
# orders=$(curl -s -H "Authorization: Bearer $token" "$go_api/orders")
# echo "$orders" | jq '.'

# sleep 2

# # 5 Check for new balances
# echo 'Checking for new balances again'
# assets=$(curl -s -H "Authorization: Bearer $token" "$go_api/assets?fields=balance,asset_type,user_id")
# echo "$assets" | jq '.'

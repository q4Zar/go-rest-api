# 0
echo 'creating user'
curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/users

# 1
echo 'login user'

token=$(curl -X POST -d '{"username": "dams", "password": "p4ssW0rd_"}' -H "Content-Type: application/json" http://127.0.0.1:8080/login | jq -r '.token')
echo $token

# 2
echo 'creates assets'

## EUR
asset_euro_success=$(curl -X POST -H "Authorization: Bearer $token" -d '{"assetType": "EUR", "balance" : 12356.54897}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
echo $asset_euro_success

### duplicate on index currency & user
asset_euro_fails_2=$(curl -X POST -H "Authorization: Bearer $token" -d '{"assetType": "EUR", "balance" : 22356.54897}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
echo $asset_euro_fails_2

##USD
asset_dollar_success=$(curl -X POST -H "Authorization: Bearer $token" -d '{"assetType": "USD", "balance" : 12356.54897}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
echo $asset_dollar_success

# ### duplicate on index currency & user
asset_dollar_fails_1=$(curl -X POST -H "Authorization: Bearer $token" -d '{"assetType": "USD", "balance" : 22356.54897}' -H "Content-Type: application/json" http://127.0.0.1:8080/assets)
echo $asset_dollar_fails_1

# 4 
# echo 'get user assets'

assets=$(curl -H "Authorization: Bearer $token" "http://127.0.0.1:8080/assets?fields=balance,asset_type")
echo "$assets" | jq '.'
    
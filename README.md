# go-lambda

curl -X POST (AWS_API_GATEWAY_URL)/test -d '{"token": "asdf"}' -H "Content-Type: application/json" => 403
curl -X POST (AWS_API_GATEWAY_URL)/test -d '{"token": "test"}' -H "Content-Type: application/json" => 200
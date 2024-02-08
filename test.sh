#! /usr/bin/env bash

# This is a short, very barebones script that can easily check
# that the API is working as intended.
# A postman collection or equivalent is probably a better choice
# for long-term maintainability.

receipt1=$(cat << JSON
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
JSON
)

receipt2=$(cat << JSON
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
JSON
)

# Send the first receipt
response=$(curl \
  -s \
  -H "Content-Type: application/json" \
  --request POST \
  --data "$receipt1" \
  http://localhost:8080/receipts/process \
)
id=$(echo "console.log(JSON.parse('$response').id)" | node)
response=$(curl -s "http://localhost:8080/receipts/$id/points")
echo "Receipt 1: $response"

# Send the second receipt
response=$(curl \
  -s \
  -H "Content-Type: application/json" \
  --request POST \
  --data "$receipt2" \
  http://localhost:8080/receipts/process \
)
id=$(echo "console.log(JSON.parse('$response').id)" | node)
response=$(curl -s "http://localhost:8080/receipts/$id/points")
echo "Receipt 2: $response"

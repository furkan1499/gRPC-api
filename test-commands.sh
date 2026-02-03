#!/bin/bash

echo "=== Creating users ==="
# Kullanıcı oluştur
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com"}'

echo -e "\n\n=== Getting user by ID ==="
# Kullanıcıyı getir
curl http://localhost:8080/users/1

echo -e "\n\n=== Listing all users ==="
# Tüm kullanıcıları listele
curl http://localhost:8080/users

echo -e "\n\n=== Testing non-existent user ==="
# Olmayan kullanıcı test et
curl http://localhost:8080/users/999
#!/bin/bash

# 设置基础URL
BASE_URL="http://localhost:8000/v1/users"

# 创建用户
echo "创建用户..."
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }')
echo "创建用户响应: $CREATE_RESPONSE"

# 获取用户ID
USER_ID=$(echo $CREATE_RESPONSE | grep -o '"id":[0-9]*' | cut -d':' -f2)
echo "创建的用户ID: $USER_ID"

# 获取所有用户
echo "获取所有用户..."
curl -s -X GET "$BASE_URL"

# 获取特定用户
echo "获取特定用户..."
curl -s -X GET "$BASE_URL/$USER_ID"

# 更新用户
echo "更新用户..."
curl -s -X PUT "$BASE_URL/$USER_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "email": "zhangsan_new@example.com",
    "age": 26
  }'

# 删除用户
echo "删除用户..."
curl -s -X DELETE "$BASE_URL/$USER_ID"

# 再次获取所有用户，确认删除
echo "获取所有用户（确认删除）..."
curl -s -X GET "$BASE_URL"
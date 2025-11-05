#!/bin/bash

# ERP System API 测试脚本

echo "Starting ERP System API Tests..."

# 启动服务器 (在后台运行)
echo "Starting server..."
go run main.go &

# 等待服务器启动
sleep 3

# 测试健康检查
echo "Testing health check..."
curl -X GET http://localhost:8080/health
echo ""

# 测试创建会计科目
echo "Creating account..."
curl -X POST http://localhost:8080/api/finance/accounts \
  -H "Content-Type: application/json" \
  -d '{
    "code": "1003",
    "name": "银行存款-工商银行",
    "account_type": "ASSET",
    "balance_direction": "DEBIT",
    "parent_code": "1002",
    "is_leaf": true
  }'
echo ""

# 测试获取会计科目
echo "Getting account..."
curl -X GET http://localhost:8080/api/finance/accounts/1003
echo ""

# 测试创建凭证
echo "Creating voucher..."
curl -X POST http://localhost:8080/api/finance/vouchers \
  -H "Content-Type: application/json" \
  -d '{
    "voucher_date": "2023-01-01T00:00:00Z",
    "description": "收到投资款",
    "entries": [
      {
        "account_code": "1002",
        "debit_amount": 1000000,
        "credit_amount": 0,
        "description": "银行存款增加"
      },
      {
        "account_code": "3001",
        "debit_amount": 0,
        "credit_amount": 1000000,
        "description": "实收资本增加"
      }
    ]
  }'
echo ""

# 测试创建采购订单
echo "Creating purchase order..."
curl -X POST http://localhost:8080/api/purchase/orders \
  -H "Content-Type: application/json" \
  -d '{
    "supplier_code": "SUP001",
    "supplier_name": "供应商A",
    "order_date": "2023-01-01T00:00:00Z",
    "delivery_date": "2023-01-10T00:00:00Z",
    "description": "采购原材料",
    "items": [
      {
        "product_code": "MAT001",
        "product_name": "原材料A",
        "unit": "公斤",
        "quantity": 1000,
        "unit_price": 10,
        "amount": 10000,
        "description": "原材料采购"
      }
    ]
  }'
echo ""

# 测试创建销售订单
echo "Creating sales order..."
curl -X POST http://localhost:8080/api/sales/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_code": "CUS001",
    "customer_name": "客户A",
    "order_date": "2023-01-01T00:00:00Z",
    "delivery_date": "2023-01-10T00:00:00Z",
    "description": "销售产品",
    "items": [
      {
        "product_code": "PRD001",
        "product_name": "产品A",
        "unit": "件",
        "quantity": 100,
        "unit_price": 50,
        "amount": 5000,
        "description": "产品销售"
      }
    ]
  }'
echo ""

echo "API tests completed."
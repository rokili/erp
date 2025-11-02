# ERP System

基于 Go 和 PostgreSQL 的企业资源规划系统

## 功能模块

1. 财务管理
   - 会计科目管理
   - 凭证录入与查询

2. 采购管理
   - 采购订单创建、审批、关闭
   - 供应商管理

3. 销售管理
   - 销售订单创建、审批、关闭
   - 客户管理

## 技术栈

- Go 1.21
- Gin Web框架
- PostgreSQL 数据库
- sqlx 数据库操作库
- Viper 配置管理

## 项目结构

```
.
├── internal/
│   ├── config/        # 配置加载
│   ├── handler/       # HTTP路由处理
│   ├── model/         # 数据模型
│   ├── repository/    # 数据访问层
│   ├── server/        # 服务启动
│   └── service/       # 业务逻辑层
├── config.yaml        # 配置文件
├── init.sql           # 数据库初始化脚本
└── main.go            # 程序入口
```

## 快速开始

### 1. 数据库设置

首先确保已安装并运行 PostgreSQL 数据库，然后执行以下步骤：

```bash
# 连接到PostgreSQL
psql -U postgres

# 执行初始化脚本
\i init.sql
```

### 2. 配置文件

编辑 [config.yaml](file:///d:/soft/pg/erp/config.yaml) 文件，根据实际情况修改数据库连接信息：

```yaml
server:
  host: localhost
  port: 8080

database:
  host: localhost
  port: 5432
  user: erp_user
  password: erp_password
  name: erp_system
  sslmode: disable
```

### 3. 运行应用

```bash
# 安装依赖
go mod tidy

# 运行应用
go run main.go
```

## API 接口

### 财务管理

#### 会计科目
- `POST /api/finance/accounts` - 创建会计科目
- `GET /api/finance/accounts/{code}` - 获取会计科目
- `GET /api/finance/accounts` - 获取所有会计科目

#### 凭证
- `POST /api/finance/vouchers` - 创建凭证
- `GET /api/finance/vouchers/{id}` - 获取凭证
- `GET /api/finance/vouchers` - 获取凭证列表

### 采购管理

#### 采购订单
- `POST /api/purchase/orders` - 创建采购订单
- `GET /api/purchase/orders/{id}` - 获取采购订单
- `GET /api/purchase/orders` - 获取采购订单列表
- `PUT /api/purchase/orders/{id}/approve` - 审批采购订单
- `PUT /api/purchase/orders/{id}/close` - 关闭采购订单

### 销售管理

#### 销售订单
- `POST /api/sales/orders` - 创建销售订单
- `GET /api/sales/orders/{id}` - 获取销售订单
- `GET /api/sales/orders` - 获取销售订单列表
- `PUT /api/sales/orders/{id}/approve` - 审批销售订单
- `PUT /api/sales/orders/{id}/close` - 关闭销售订单

## 示例请求

### 创建会计科目
```json
POST /api/finance/accounts
{
  "code": "1003",
  "name": "银行存款-工商银行",
  "account_type": "ASSET",
  "balance_direction": "DEBIT",
  "parent_code": "1002",
  "is_leaf": true
}
```

### 创建凭证
```json
POST /api/finance/vouchers
{
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
}
```

### 创建采购订单
```json
POST /api/purchase/orders
{
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
}
```

### 创建销售订单
```json
POST /api/sales/orders
{
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
}
```
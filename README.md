# ERP System

基于 Go 和 PostgreSQL 的企业资源规划系统

## 功能模块

1. 财务管理
   - 会计科目管理
   - 凭证录入与查询
   - 系统开账管理（开账前允许单方面调整，开账后严格执行借贷平衡）

2. 采购管理
   - 采购订单创建、审批、关闭
   - 供应商管理

3. 销售管理
   - 销售订单创建、审批、关闭
   - 客户管理

4. 库存管理
   - 商品分类管理
   - 商品管理
   - 库存跟踪
   - 成本核算（加权平均法、先进先出法）

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
├── web/               # 前端静态文件
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
  user: postgres
  password: Aa123456
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

## 系统开账功能

系统支持开账管理，确保财务数据的严谨性：

1. **开账前**：允许单方面调整初始数据，便于录入期初余额
2. **开账后**：严格执行"有借必有贷，借贷必相等"的会计原则
3. **操作流程**：
   - 系统初始化并录入期初数据
   - 通过前端界面或API调用开账接口
   - 开账后所有凭证必须借贷平衡

## API 接口

### 系统管理

#### 系统状态
- `GET /api/system/status` - 获取系统状态
- `POST /api/system/open` - 系统开账

### 财务管理

#### 会计科目
- `POST /api/finance/accounts` - 创建会计科目
- `GET /api/finance/accounts/{code}` - 获取会计科目
- `GET /api/finance/accounts` - 获取所有会计科目

#### 凭证
- `POST /api/finance/vouchers` - 创建凭证
- `GET /api/finance/vouchers/{id}` - 获取凭证
- `GET /api/finance/vouchers` - 获取凭证列表

### 商品管理

#### 商品分类
- `POST /api/categories` - 创建商品分类
- `GET /api/categories/{id}` - 获取商品分类
- `GET /api/categories` - 获取所有商品分类

#### 商品
- `POST /api/products` - 创建商品
- `GET /api/products/{id}` - 获取商品
- `GET /api/products` - 获取所有商品

### 采购管理

#### 供应商
- `POST /api/suppliers` - 创建供应商
- `GET /api/suppliers/{id}` - 获取供应商
- `GET /api/suppliers` - 获取所有供应商

#### 采购订单
- `POST /api/purchase/orders` - 创建采购订单
- `GET /api/purchase/orders/{id}` - 获取采购订单
- `GET /api/purchase/orders` - 获取采购订单列表
- `PUT /api/purchase/orders/{id}/approve` - 审批采购订单
- `PUT /api/purchase/orders/{id}/close` - 关闭采购订单

### 销售管理

#### 客户
- `POST /api/customers` - 创建客户
- `GET /api/customers/{id}` - 获取客户
- `GET /api/customers` - 获取所有客户

#### 销售订单
- `POST /api/sales/orders` - 创建销售订单
- `GET /api/sales/orders/{id}` - 获取销售订单
- `GET /api/sales/orders` - 获取销售订单列表
- `PUT /api/sales/orders/{id}/approve` - 审批销售订单
- `PUT /api/sales/orders/{id}/close` - 关闭销售订单

## 示例请求

### 系统开账
```bash
# 开账前检查系统状态
curl -X GET http://localhost:8080/api/system/status

# 执行开账操作
curl -X POST http://localhost:8080/api/system/open
```

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

### 创建凭证（开账后必须借贷平衡）
```json
POST /api/finance/vouchers
{
  "voucher_date": "2025-11-05T00:00:00Z",
  "description": "收到投资款",
  "entries": [
    {
      "account_code": "1001",
      "debit_amount": 1000000,
      "credit_amount": 0,
      "description": "库存现金增加"
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

### 创建商品分类
```json
POST /api/categories
{
  "code": "CAT001",
  "name": "电子产品",
  "description": "电子类产品分类"
}
```

### 创建商品
```json
POST /api/products
{
  "code": "PRD001",
  "name": "智能手机",
  "category_id": 1,
  "unit": "台",
  "specification": "64GB",
  "description": "高性能智能手机"
}
```

### 创建供应商
```json
POST /api/suppliers
{
  "code": "SUP001",
  "name": "供应商A",
  "contact_person": "张三",
  "phone": "13800138000",
  "email": "zhangsan@supplier.com",
  "address": "北京市朝阳区xxx街道",
  "tax_number": "911101087890123456",
  "bank_name": "中国银行",
  "bank_account": "1234567890123456"
}
```

### 创建客户
```json
POST /api/customers
{
  "code": "CUS001",
  "name": "客户A",
  "contact_person": "李四",
  "phone": "13900139000",
  "email": "lisi@customer.com",
  "address": "上海市浦东新区xxx街道",
  "tax_number": "913101156789012345",
  "bank_name": "工商银行",
  "bank_account": "6543210987654321"
}
```

### 创建采购订单
```json
POST /api/purchase/orders
{
  "supplier_code": "SUP001",
  "supplier_name": "供应商A",
  "order_date": "2025-11-05T00:00:00Z",
  "delivery_date": "2025-11-10T00:00:00Z",
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
  "order_date": "2025-11-05T00:00:00Z",
  "delivery_date": "2025-11-10T00:00:00Z",
  "description": "销售产品",
  "items": [
    {
      "product_code": "PRD001",
      "product_name": "智能手机",
      "unit": "台",
      "quantity": 100,
      "unit_price": 5000,
      "amount": 500000,
      "description": "产品销售"
    }
  ]
}
```
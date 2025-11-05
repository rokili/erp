<<<<<<< HEAD
-- 创建数据库
CREATE DATABASE erp_system;

-- 连接到数据库
\c erp_system;

-- 创建会计科目表
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    account_type VARCHAR(20) NOT NULL,
    balance_direction VARCHAR(10) NOT NULL,
    parent_code VARCHAR(20),
    is_leaf BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建凭证表
CREATE TABLE vouchers (
    id SERIAL PRIMARY KEY,
    voucher_date DATE NOT NULL,
    description TEXT,
    status VARCHAR(20) DEFAULT 'DRAFT',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建凭证明细表
CREATE TABLE voucher_entries (
    id SERIAL PRIMARY KEY,
    voucher_id INTEGER REFERENCES vouchers(id) ON DELETE CASCADE,
    account_code VARCHAR(20) NOT NULL,
    debit_amount DECIMAL(15,2) DEFAULT 0,
    credit_amount DECIMAL(15,2) DEFAULT 0,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建商品分类表
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建商品表
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    category_id INTEGER REFERENCES categories(id),
    unit VARCHAR(20),
    specification VARCHAR(100),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建供应商表
CREATE TABLE suppliers (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    contact_person VARCHAR(50),
    phone VARCHAR(20),
    email VARCHAR(100),
    address TEXT,
    tax_number VARCHAR(50),
    bank_name VARCHAR(100),
    bank_account VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建客户表
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    contact_person VARCHAR(50),
    phone VARCHAR(20),
    email VARCHAR(100),
    address TEXT,
    tax_number VARCHAR(50),
    bank_name VARCHAR(100),
    bank_account VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建采购订单表
CREATE TABLE purchase_orders (
    id SERIAL PRIMARY KEY,
    supplier_code VARCHAR(20) NOT NULL,
    supplier_name VARCHAR(100) NOT NULL,
    order_date DATE NOT NULL,
    delivery_date DATE,
    description TEXT,
    status VARCHAR(20) DEFAULT 'DRAFT',
    total_amount DECIMAL(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建采购订单明细表
CREATE TABLE purchase_order_items (
    id SERIAL PRIMARY KEY,
    purchase_order_id INTEGER REFERENCES purchase_orders(id) ON DELETE CASCADE,
    product_code VARCHAR(20) NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    unit VARCHAR(20),
    quantity DECIMAL(10,2) NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建销售订单表
CREATE TABLE sales_orders (
    id SERIAL PRIMARY KEY,
    customer_code VARCHAR(20) NOT NULL,
    customer_name VARCHAR(100) NOT NULL,
    order_date DATE NOT NULL,
    delivery_date DATE,
    description TEXT,
    status VARCHAR(20) DEFAULT 'DRAFT',
    total_amount DECIMAL(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建销售订单明细表
CREATE TABLE sales_order_items (
    id SERIAL PRIMARY KEY,
    sales_order_id INTEGER REFERENCES sales_orders(id) ON DELETE CASCADE,
    product_code VARCHAR(20) NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    unit VARCHAR(20),
    quantity DECIMAL(10,2) NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建用户表
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(20),
    status VARCHAR(20) DEFAULT 'ACTIVE',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建角色表
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建用户角色关联表
CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, role_id)
);

-- 创建权限表
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    resource VARCHAR(100),
    action VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建角色权限关联表
CREATE TABLE role_permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(role_id, permission_id)
);

-- 插入默认角色
INSERT INTO roles (name, description) VALUES 
('采购员', '负责采购订单的创建和管理'),
('销售', '负责销售订单的创建和管理'),
('库房', '负责库存管理和出入库操作'),
('出纳', '负责现金和银行存款的收支操作'),
('会计', '负责财务凭证的录入和审核'),
('经理', '拥有系统所有权限');

-- 插入默认权限
INSERT INTO permissions (name, description, resource, action) VALUES 
('创建采购订单', '创建采购订单', 'purchase_orders', 'create'),
('审批采购订单', '审批采购订单', 'purchase_orders', 'approve'),
('关闭采购订单', '关闭采购订单', 'purchase_orders', 'close'),
('查看采购订单', '查看采购订单', 'purchase_orders', 'read'),
('创建销售订单', '创建销售订单', 'sales_orders', 'create'),
('审批销售订单', '审批销售订单', 'sales_orders', 'approve'),
('关闭销售订单', '关闭销售订单', 'sales_orders', 'close'),
('查看销售订单', '查看销售订单', 'sales_orders', 'read'),
('管理商品', '管理商品信息', 'products', 'manage'),
('管理库存', '管理库存', 'inventory', 'manage'),
('现金收支', '现金收支操作', 'cash', 'manage'),
('银行存取', '银行存取操作', 'bank', 'manage'),
('凭证录入', '财务凭证录入', 'vouchers', 'create'),
('凭证审核', '财务凭证审核', 'vouchers', 'approve'),
('系统管理', '系统管理权限', 'system', 'manage');

-- 为经理角色分配所有权限
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM roles r, permissions p 
WHERE r.name = '经理';

-- 为其他角色分配相应权限
-- 采购员权限
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM roles r, permissions p 
WHERE r.name = '采购员' 
AND p.name IN ('创建采购订单', '查看采购订单');

-- 销售权限
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM roles r, permissions p 
WHERE r.name = '销售' 
AND p.name IN ('创建销售订单', '查看销售订单');

-- 库房权限
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM roles r, permissions p 
WHERE r.name = '库房' 
AND p.name IN ('管理商品', '管理库存');

-- 出纳权限
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM roles r, permissions p 
WHERE r.name = '出纳' 
AND p.name IN ('现金收支', '银行存取');

-- 会计权限
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM roles r, permissions p 
WHERE r.name = '会计' 
AND p.name IN ('凭证录入', '凭证审核');

-- 插入默认会计科目
INSERT INTO accounts (code, name, account_type, balance_direction, is_leaf) VALUES
-- 资产类科目
('1001', '库存现金', 'ASSET', 'DEBIT', true),
('1002', '银行存款', 'ASSET', 'DEBIT', true),
('1101', '应收账款', 'ASSET', 'DEBIT', true),
('1201', '其他应收款', 'ASSET', 'DEBIT', true),
('1401', '原材料', 'ASSET', 'DEBIT', true),
('1402', '库存商品', 'ASSET', 'DEBIT', true),
('1403', '半成品', 'ASSET', 'DEBIT', true),
('1501', '固定资产', 'ASSET', 'DEBIT', true),
('1502', '累计折旧', 'ASSET', 'CREDIT', true),
('1601', '无形资产', 'ASSET', 'DEBIT', true),

-- 负债类科目
('2001', '短期借款', 'LIABILITY', 'CREDIT', true),
('2101', '应付账款', 'LIABILITY', 'CREDIT', true),
('2201', '应付职工薪酬', 'LIABILITY', 'CREDIT', true),
('2202', '应交税费', 'LIABILITY', 'CREDIT', true),
('2203', '应付股利', 'LIABILITY', 'CREDIT', true),
('2301', '其他应付款', 'LIABILITY', 'CREDIT', true),
('2401', '长期借款', 'LIABILITY', 'CREDIT', true),

-- 所有者权益类科目
('3001', '实收资本', 'EQUITY', 'CREDIT', true),
('3002', '资本公积', 'EQUITY', 'CREDIT', true),
('3101', '盈余公积', 'EQUITY', 'CREDIT', true),
('3102', '未分配利润', 'EQUITY', 'CREDIT', true),

-- 成本类科目
('4001', '生产成本', 'COST', 'DEBIT', true),
('4101', '制造费用', 'COST', 'DEBIT', true),
('4201', '主营业务成本', 'COST', 'DEBIT', true),
('4202', '其他业务成本', 'COST', 'DEBIT', true),

-- 损益类科目
('5001', '主营业务收入', 'INCOME', 'CREDIT', true),
('5002', '其他业务收入', 'INCOME', 'CREDIT', true),
('5101', '投资收益', 'INCOME', 'CREDIT', true),
('5201', '营业外收入', 'INCOME', 'CREDIT', true),
('5301', '主营业务成本', 'EXPENSE', 'DEBIT', true),
('5302', '其他业务成本', 'EXPENSE', 'DEBIT', true),
('5401', '税金及附加', 'EXPENSE', 'DEBIT', true),
('5402', '销售费用', 'EXPENSE', 'DEBIT', true),
('5403', '管理费用', 'EXPENSE', 'DEBIT', true),
('5404', '财务费用', 'EXPENSE', 'DEBIT', true),
('5501', '营业外支出', 'EXPENSE', 'DEBIT', true),
('5601', '所得税费用', 'EXPENSE', 'DEBIT', true);
=======
-- Create database
CREATE DATABASE erp_system;

-- Connect to the database
\c erp_system;

-- Create finance_account table
CREATE TABLE finance_account (
    code VARCHAR(32) PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    account_type VARCHAR(32) NOT NULL,
    balance_direction VARCHAR(8) NOT NULL,
    parent_code VARCHAR(32),
    is_leaf BOOLEAN DEFAULT FALSE,
    status VARCHAR(16) DEFAULT 'ACTIVE',
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create finance_voucher table
CREATE TABLE finance_voucher (
    id SERIAL PRIMARY KEY,
    voucher_no VARCHAR(32) UNIQUE NOT NULL,
    voucher_date DATE NOT NULL,
    description TEXT,
    total_amount NUMERIC(18,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create finance_voucher_entry table
CREATE TABLE finance_voucher_entry (
    id SERIAL PRIMARY KEY,
    voucher_id INTEGER NOT NULL REFERENCES finance_voucher(id) ON DELETE CASCADE,
    account_code VARCHAR(32) NOT NULL REFERENCES finance_account(code),
    debit_amount NUMERIC(18,2) DEFAULT 0,
    credit_amount NUMERIC(18,2) DEFAULT 0,
    description TEXT
);

-- Create purchase_order table
CREATE TABLE purchase_order (
    id SERIAL PRIMARY KEY,
    order_no VARCHAR(32) UNIQUE NOT NULL,
    supplier_code VARCHAR(32) NOT NULL,
    supplier_name VARCHAR(128) NOT NULL,
    order_date DATE NOT NULL,
    delivery_date DATE,
    total_amount NUMERIC(18,2) NOT NULL,
    status VARCHAR(16) DEFAULT 'DRAFT', -- DRAFT, APPROVED, RECEIVED, CLOSED
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create purchase_order_item table
CREATE TABLE purchase_order_item (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES purchase_order(id) ON DELETE CASCADE,
    product_code VARCHAR(32) NOT NULL,
    product_name VARCHAR(128) NOT NULL,
    unit VARCHAR(16) NOT NULL,
    quantity NUMERIC(18,3) NOT NULL,
    unit_price NUMERIC(18,2) NOT NULL,
    amount NUMERIC(18,2) NOT NULL,
    received_qty NUMERIC(18,3) DEFAULT 0,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create sales_order table
CREATE TABLE sales_order (
    id SERIAL PRIMARY KEY,
    order_no VARCHAR(32) UNIQUE NOT NULL,
    customer_code VARCHAR(32) NOT NULL,
    customer_name VARCHAR(128) NOT NULL,
    order_date DATE NOT NULL,
    delivery_date DATE,
    total_amount NUMERIC(18,2) NOT NULL,
    status VARCHAR(16) DEFAULT 'DRAFT', -- DRAFT, APPROVED, DELIVERED, CLOSED
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create sales_order_item table
CREATE TABLE sales_order_item (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES sales_order(id) ON DELETE CASCADE,
    product_code VARCHAR(32) NOT NULL,
    product_name VARCHAR(128) NOT NULL,
    unit VARCHAR(16) NOT NULL,
    quantity NUMERIC(18,3) NOT NULL,
    unit_price NUMERIC(18,2) NOT NULL,
    amount NUMERIC(18,2) NOT NULL,
    delivered_qty NUMERIC(18,3) DEFAULT 0,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create product_category table
CREATE TABLE product_category (
    id SERIAL PRIMARY KEY,
    code VARCHAR(32) UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    parent_id INTEGER REFERENCES product_category(id),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create product table
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    code VARCHAR(32) UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    category_id INTEGER REFERENCES product_category(id),
    unit VARCHAR(16) NOT NULL,
    specification VARCHAR(128),
    description TEXT,
    status VARCHAR(16) DEFAULT 'ACTIVE', -- ACTIVE, INACTIVE
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create supplier table
CREATE TABLE supplier (
    id SERIAL PRIMARY KEY,
    code VARCHAR(32) UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    contact_person VARCHAR(64),
    phone VARCHAR(32),
    email VARCHAR(128),
    address TEXT,
    tax_number VARCHAR(64),
    bank_name VARCHAR(128),
    bank_account VARCHAR(64),
    status VARCHAR(16) DEFAULT 'ACTIVE', -- ACTIVE, INACTIVE
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create customer table
CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    code VARCHAR(32) UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    contact_person VARCHAR(64),
    phone VARCHAR(32),
    email VARCHAR(128),
    address TEXT,
    tax_number VARCHAR(64),
    bank_name VARCHAR(128),
    bank_account VARCHAR(64),
    status VARCHAR(16) DEFAULT 'ACTIVE', -- ACTIVE, INACTIVE
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create inventory table
CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES product(id),
    warehouse VARCHAR(64) DEFAULT 'MAIN',
    quantity NUMERIC(18,3) DEFAULT 0,
    unit_cost NUMERIC(18,2) DEFAULT 0, -- 单位成本
    total_cost NUMERIC(18,2) DEFAULT 0, -- 总成本
    cost_method VARCHAR(16) DEFAULT 'WEIGHTED_AVERAGE', -- 成本核算方法: WEIGHTED_AVERAGE, FIFO
    last_updated TIMESTAMP DEFAULT NOW(),
    UNIQUE(product_id, warehouse)
);

-- Create inventory_transaction table
CREATE TABLE inventory_transaction (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES product(id),
    transaction_type VARCHAR(16) NOT NULL, -- IN, OUT
    quantity NUMERIC(18,3) NOT NULL,
    unit_cost NUMERIC(18,2) NOT NULL, -- 交易时的单位成本
    total_cost NUMERIC(18,2) NOT NULL, -- 交易时的总成本
    reference_type VARCHAR(32), -- PURCHASE_ORDER, SALES_ORDER, etc.
    reference_id INTEGER, -- 关联的订单ID
    warehouse VARCHAR(64) DEFAULT 'MAIN',
    transaction_date TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create inventory_fifo_layer table (用于先进先出法的成本核算)
CREATE TABLE inventory_fifo_layer (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES product(id),
    quantity NUMERIC(18,3) NOT NULL,
    unit_cost NUMERIC(18,2) NOT NULL,
    total_cost NUMERIC(18,2) NOT NULL,
    remaining_quantity NUMERIC(18,3) NOT NULL,
    receipt_date TIMESTAMP DEFAULT NOW(), -- 入库日期
    warehouse VARCHAR(64) DEFAULT 'MAIN'
);

-- Create system_config table (用于存储系统配置，如开账状态)
CREATE TABLE system_config (
    id SERIAL PRIMARY KEY,
    config_key VARCHAR(64) UNIQUE NOT NULL,
    config_value TEXT,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_voucher_date ON finance_voucher(voucher_date);
CREATE INDEX idx_voucher_entry_account ON finance_voucher_entry(account_code);
CREATE INDEX idx_voucher_entry_voucher ON finance_voucher_entry(voucher_id);
CREATE INDEX idx_purchase_order_date ON purchase_order(order_date);
CREATE INDEX idx_purchase_order_status ON purchase_order(status);
CREATE INDEX idx_purchase_order_item_order ON purchase_order_item(order_id);
CREATE INDEX idx_sales_order_date ON sales_order(order_date);
CREATE INDEX idx_sales_order_status ON sales_order(status);
CREATE INDEX idx_sales_order_item_order ON sales_order_item(order_id);
CREATE INDEX idx_product_category ON product(category_id);
CREATE INDEX idx_inventory_product ON inventory(product_id);
CREATE INDEX idx_inventory_transaction_product ON inventory_transaction(product_id);
CREATE INDEX idx_inventory_transaction_date ON inventory_transaction(transaction_date);
CREATE INDEX idx_inventory_fifo_layer_product ON inventory_fifo_layer(product_id);

-- Insert sample data
INSERT INTO finance_account (code, name, account_type, balance_direction, is_leaf) VALUES
('1001', '库存现金', 'ASSET', 'DEBIT', true),
('1002', '银行存款', 'ASSET', 'DEBIT', true),
('1122', '应收账款', 'ASSET', 'DEBIT', true),
('1401', '原材料', 'ASSET', 'DEBIT', true),
('1601', '固定资产', 'ASSET', 'DEBIT', true),
('2201', '应付账款', 'LIABILITY', 'CREDIT', true),
('2202', '应付票据', 'LIABILITY', 'CREDIT', true),
('2211', '应付职工薪酬', 'LIABILITY', 'CREDIT', true),
('2221', '应交税费', 'LIABILITY', 'CREDIT', true),
('3001', '实收资本', 'EQUITY', 'CREDIT', true),
('3002', '资本公积', 'EQUITY', 'CREDIT', true),
('3101', '盈余公积', 'EQUITY', 'CREDIT', true),
('3103', '本年利润', 'EQUITY', 'CREDIT', true),
('3104', '利润分配', 'EQUITY', 'CREDIT', true),
('6001', '主营业务收入', 'INCOME', 'CREDIT', true),
('6051', '其他业务收入', 'INCOME', 'CREDIT', true),
('6101', '投资收益', 'INCOME', 'CREDIT', true),
('6301', '营业外收入', 'INCOME', 'CREDIT', true),
('6401', '主营业务成本', 'EXPENSE', 'DEBIT', true),
('6402', '其他业务成本', 'EXPENSE', 'DEBIT', true),
('6403', '营业税金及附加', 'EXPENSE', 'DEBIT', true),
('6601', '销售费用', 'EXPENSE', 'DEBIT', true),
('6602', '管理费用', 'EXPENSE', 'DEBIT', true),
('6603', '财务费用', 'EXPENSE', 'DEBIT', true),
('6701', '资产减值损失', 'EXPENSE', 'DEBIT', true),
('6711', '营业外支出', 'EXPENSE', 'DEBIT', true);

-- Insert system configuration
INSERT INTO system_config (config_key, config_value, description) VALUES
('system_opened', 'false', '系统是否已开账，false表示未开账，true表示已开账');
>>>>>>> eb04c5bc7c6a1998a4c109ada9e19202dab00b44

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
package repository

import (
	"erp/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// Voucher operations
func (r *Repository) CreateVoucher(voucher *model.Voucher) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 插入凭证头
	query := `INSERT INTO finance_voucher (voucher_no, voucher_date, description, total_amount, created_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRowx(query, voucher.VoucherNo, voucher.VoucherDate, voucher.Description, voucher.TotalAmount, time.Now()).Scan(&voucher.ID)
	if err != nil {
		return err
	}

	// 插入凭证明细
	for i := range voucher.Entries {
		entry := &voucher.Entries[i]
		query = `INSERT INTO finance_voucher_entry (voucher_id, account_code, debit_amount, credit_amount, description) 
		         VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err = tx.QueryRowx(query, voucher.ID, entry.AccountCode, entry.DebitAmount, entry.CreditAmount, entry.Description).Scan(&entry.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *Repository) GetVoucher(id int64) (*model.Voucher, error) {
	// 查询凭证头
	voucher := &model.Voucher{}
	query := `SELECT id, voucher_no, voucher_date, description, total_amount, created_at 
	          FROM finance_voucher WHERE id = $1`
	err := r.db.Get(voucher, query, id)
	if err != nil {
		return nil, err
	}

	// 查询凭证明细
	query = `SELECT id, voucher_id, account_code, debit_amount, credit_amount, description 
	         FROM finance_voucher_entry WHERE voucher_id = $1 ORDER BY id`
	err = r.db.Select(&voucher.Entries, query, id)
	if err != nil {
		return nil, err
	}

	return voucher, nil
}

func (r *Repository) ListVouchers(limit, offset int) ([]*model.Voucher, error) {
	vouchers := []*model.Voucher{}
	query := `SELECT id, voucher_no, voucher_date, description, total_amount, created_at 
	          FROM finance_voucher ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	err := r.db.Select(&vouchers, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

// Account operations
func (r *Repository) CreateAccount(account *model.Account) error {
	query := `INSERT INTO finance_account (code, name, account_type, balance_direction, parent_code, is_leaf, status, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, account.Code, account.Name, account.AccountType, account.BalanceDirection,
		account.ParentCode, account.IsLeaf, account.Status, time.Now())
	return err
}

func (r *Repository) GetAccount(code string) (*model.Account, error) {
	account := &model.Account{}
	query := `SELECT code, name, account_type, balance_direction, parent_code, is_leaf, status, created_at 
	          FROM finance_account WHERE code = $1`
	err := r.db.Get(account, query, code)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *Repository) ListAccounts() ([]*model.Account, error) {
	accounts := []*model.Account{}
	query := `SELECT code, name, account_type, balance_direction, parent_code, is_leaf, status, created_at 
	          FROM finance_account ORDER BY code`
	err := r.db.Select(&accounts, query)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// ProductCategory operations
func (r *Repository) CreateProductCategory(category *model.ProductCategory) error {
	query := `INSERT INTO product_category (code, name, parent_id, description, created_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRowx(query, category.Code, category.Name, category.ParentID, category.Description, time.Now()).Scan(&category.ID)
}

func (r *Repository) GetProductCategory(id int64) (*model.ProductCategory, error) {
	category := &model.ProductCategory{}
	query := `SELECT id, code, name, parent_id, description, created_at 
	          FROM product_category WHERE id = $1`
	err := r.db.Get(category, query, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *Repository) ListProductCategories() ([]*model.ProductCategory, error) {
	categories := []*model.ProductCategory{}
	query := `SELECT id, code, name, parent_id, description, created_at 
	          FROM product_category ORDER BY name`
	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// Product operations
func (r *Repository) CreateProduct(product *model.Product) error {
	query := `INSERT INTO product (code, name, category_id, unit, specification, description, status, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	return r.db.QueryRowx(query, product.Code, product.Name, product.CategoryID, product.Unit, product.Specification, product.Description, product.Status, time.Now()).Scan(&product.ID)
}

func (r *Repository) GetProduct(id int64) (*model.Product, error) {
	product := &model.Product{}
	query := `SELECT id, code, name, category_id, unit, specification, description, status, created_at 
	          FROM product WHERE id = $1`
	err := r.db.Get(product, query, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repository) ListProducts() ([]*model.Product, error) {
	products := []*model.Product{}
	query := `SELECT id, code, name, category_id, unit, specification, description, status, created_at 
	          FROM product ORDER BY name`
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Supplier operations
func (r *Repository) CreateSupplier(supplier *model.Supplier) error {
	query := `INSERT INTO supplier (code, name, contact_person, phone, email, address, tax_number, bank_name, bank_account, status, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	return r.db.QueryRowx(query, supplier.Code, supplier.Name, supplier.ContactPerson, supplier.Phone, supplier.Email, supplier.Address, supplier.TaxNumber, supplier.BankName, supplier.BankAccount, supplier.Status, time.Now()).Scan(&supplier.ID)
}

func (r *Repository) GetSupplier(id int64) (*model.Supplier, error) {
	supplier := &model.Supplier{}
	query := `SELECT id, code, name, contact_person, phone, email, address, tax_number, bank_name, bank_account, status, created_at 
	          FROM supplier WHERE id = $1`
	err := r.db.Get(supplier, query, id)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (r *Repository) ListSuppliers() ([]*model.Supplier, error) {
	suppliers := []*model.Supplier{}
	query := `SELECT id, code, name, contact_person, phone, email, address, tax_number, bank_name, bank_account, status, created_at 
	          FROM supplier ORDER BY name`
	err := r.db.Select(&suppliers, query)
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

// Customer operations
func (r *Repository) CreateCustomer(customer *model.Customer) error {
	query := `INSERT INTO customer (code, name, contact_person, phone, email, address, tax_number, bank_name, bank_account, status, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	return r.db.QueryRowx(query, customer.Code, customer.Name, customer.ContactPerson, customer.Phone, customer.Email, customer.Address, customer.TaxNumber, customer.BankName, customer.BankAccount, customer.Status, time.Now()).Scan(&customer.ID)
}

func (r *Repository) GetCustomer(id int64) (*model.Customer, error) {
	customer := &model.Customer{}
	query := `SELECT id, code, name, contact_person, phone, email, address, tax_number, bank_name, bank_account, status, created_at 
	          FROM customer WHERE id = $1`
	err := r.db.Get(customer, query, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *Repository) ListCustomers() ([]*model.Customer, error) {
	customers := []*model.Customer{}
	query := `SELECT id, code, name, contact_person, phone, email, address, tax_number, bank_name, bank_account, status, created_at 
	          FROM customer ORDER BY name`
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// SystemConfig operations
func (r *Repository) GetSystemConfig(key string) (*model.SystemConfig, error) {
	config := &model.SystemConfig{}
	query := `SELECT id, config_key, config_value, description, created_at, updated_at 
	          FROM system_config WHERE config_key = $1`
	err := r.db.Get(config, query, key)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (r *Repository) SetSystemConfig(key, value, description string) error {
	// 检查配置项是否存在
	var count int
	query := `SELECT COUNT(*) FROM system_config WHERE config_key = $1`
	err := r.db.Get(&count, query, key)
	if err != nil {
		return err
	}

	if count > 0 {
		// 更新现有配置
		query = `UPDATE system_config SET config_value = $1, description = $2, updated_at = $3 WHERE config_key = $4`
		_, err = r.db.Exec(query, value, description, time.Now(), key)
	} else {
		// 插入新配置
		query = `INSERT INTO system_config (config_key, config_value, description, created_at, updated_at) 
		         VALUES ($1, $2, $3, $4, $5)`
		_, err = r.db.Exec(query, key, value, description, time.Now(), time.Now())
	}

	return err
}

func (r *Repository) IsSystemOpened() (bool, error) {
	config, err := r.GetSystemConfig("system_opened")
	if err != nil {
		// 如果配置项不存在，返回默认值 false（未开账）
		return false, nil
	}

	return config.ConfigValue == "true", nil
}

func (r *Repository) OpenSystem() error {
	return r.SetSystemConfig("system_opened", "true", "系统已开账")
}

// PurchaseOrder operations
func (r *Repository) CreatePurchaseOrder(order *model.PurchaseOrder) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 插入采购订单头
	query := `INSERT INTO purchase_order (order_no, supplier_code, supplier_name, order_date, delivery_date, total_amount, status, description, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err = tx.QueryRowx(query, order.OrderNo, order.SupplierCode, order.SupplierName, order.OrderDate, order.DeliveryDate,
		order.TotalAmount, order.Status, order.Description, time.Now(), time.Now()).Scan(&order.ID)
	if err != nil {
		return err
	}

	// 插入采购订单明细
	for i := range order.Items {
		item := &order.Items[i]
		query = `INSERT INTO purchase_order_item (order_id, product_code, product_name, unit, quantity, unit_price, amount, description, created_at) 
		         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
		err = tx.QueryRowx(query, order.ID, item.ProductCode, item.ProductName, item.Unit, item.Quantity,
			item.UnitPrice, item.Amount, item.Description, time.Now()).Scan(&item.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *Repository) GetPurchaseOrder(id int64) (*model.PurchaseOrder, error) {
	// 查询采购订单头
	order := &model.PurchaseOrder{}
	query := `SELECT id, order_no, supplier_code, supplier_name, order_date, delivery_date, total_amount, status, description, created_at, updated_at 
	          FROM purchase_order WHERE id = $1`
	err := r.db.Get(order, query, id)
	if err != nil {
		return nil, err
	}

	// 查询采购订单明细
	query = `SELECT id, order_id, product_code, product_name, unit, quantity, unit_price, amount, received_qty, description, created_at 
	         FROM purchase_order_item WHERE order_id = $1 ORDER BY id`
	err = r.db.Select(&order.Items, query, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) ListPurchaseOrders(limit, offset int) ([]*model.PurchaseOrder, error) {
	orders := []*model.PurchaseOrder{}
	query := `SELECT id, order_no, supplier_code, supplier_name, order_date, delivery_date, total_amount, status, description, created_at, updated_at 
	          FROM purchase_order ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	err := r.db.Select(&orders, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// 查询每个订单的明细
	for _, order := range orders {
		query = `SELECT id, order_id, product_code, product_name, unit, quantity, unit_price, amount, received_qty, description, created_at 
		         FROM purchase_order_item WHERE order_id = $1 ORDER BY id`
		err = r.db.Select(&order.Items, query, order.ID)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (r *Repository) UpdatePurchaseOrderStatus(id int64, status string) error {
	query := `UPDATE purchase_order SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(query, status, time.Now(), id)
	return err
}

// SalesOrder operations
func (r *Repository) CreateSalesOrder(order *model.SalesOrder) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 插入销售订单头
	query := `INSERT INTO sales_order (order_no, customer_code, customer_name, order_date, delivery_date, total_amount, status, description, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err = tx.QueryRowx(query, order.OrderNo, order.CustomerCode, order.CustomerName, order.OrderDate, order.DeliveryDate,
		order.TotalAmount, order.Status, order.Description, time.Now(), time.Now()).Scan(&order.ID)
	if err != nil {
		return err
	}

	// 插入销售订单明细
	for i := range order.Items {
		item := &order.Items[i]
		query = `INSERT INTO sales_order_item (order_id, product_code, product_name, unit, quantity, unit_price, amount, description, created_at) 
		         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
		err = tx.QueryRowx(query, order.ID, item.ProductCode, item.ProductName, item.Unit, item.Quantity,
			item.UnitPrice, item.Amount, item.Description, time.Now()).Scan(&item.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *Repository) GetSalesOrder(id int64) (*model.SalesOrder, error) {
	// 查询销售订单头
	order := &model.SalesOrder{}
	query := `SELECT id, order_no, customer_code, customer_name, order_date, delivery_date, total_amount, status, description, created_at, updated_at 
	          FROM sales_order WHERE id = $1`
	err := r.db.Get(order, query, id)
	if err != nil {
		return nil, err
	}

	// 查询销售订单明细
	query = `SELECT id, order_id, product_code, product_name, unit, quantity, unit_price, amount, delivered_qty, description, created_at 
	         FROM sales_order_item WHERE order_id = $1 ORDER BY id`
	err = r.db.Select(&order.Items, query, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) ListSalesOrders(limit, offset int) ([]*model.SalesOrder, error) {
	orders := []*model.SalesOrder{}
	query := `SELECT id, order_no, customer_code, customer_name, order_date, delivery_date, total_amount, status, description, created_at, updated_at 
	          FROM sales_order ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	err := r.db.Select(&orders, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// 查询每个订单的明细
	for _, order := range orders {
		query = `SELECT id, order_id, product_code, product_name, unit, quantity, unit_price, amount, delivered_qty, description, created_at 
		         FROM sales_order_item WHERE order_id = $1 ORDER BY id`
		err = r.db.Select(&order.Items, query, order.ID)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (r *Repository) UpdateSalesOrderStatus(id int64, status string) error {
	query := `UPDATE sales_order SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(query, status, time.Now(), id)
	return err
}

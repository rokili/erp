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

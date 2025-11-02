package model

import (
	"time"
)

// Voucher represents a financial voucher
type Voucher struct {
	ID          int64          `db:"id" json:"id"`
	VoucherNo   string         `db:"voucher_no" json:"voucher_no"`
	VoucherDate time.Time      `db:"voucher_date" json:"voucher_date"`
	Description string         `db:"description" json:"description"`
	TotalAmount float64        `db:"total_amount" json:"total_amount"`
	Entries     []VoucherEntry `json:"entries"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
}

// VoucherEntry represents a voucher entry
type VoucherEntry struct {
	ID           int64   `db:"id" json:"id"`
	VoucherID    int64   `db:"voucher_id" json:"voucher_id"`
	AccountCode  string  `db:"account_code" json:"account_code"`
	DebitAmount  float64 `db:"debit_amount" json:"debit_amount"`
	CreditAmount float64 `db:"credit_amount" json:"credit_amount"`
	Description  string  `db:"description" json:"description"`
}

// Account represents a financial account
type Account struct {
	Code             string    `db:"code" json:"code"`
	Name             string    `db:"name" json:"name"`
	AccountType      string    `db:"account_type" json:"account_type"`
	BalanceDirection string    `db:"balance_direction" json:"balance_direction"`
	ParentCode       string    `db:"parent_code" json:"parent_code"`
	IsLeaf           bool      `db:"is_leaf" json:"is_leaf"`
	Status           string    `db:"status" json:"status"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

// CreateVoucherDTO represents the data transfer object for creating a voucher
type CreateVoucherDTO struct {
	VoucherDate time.Time      `json:"voucher_date"`
	Description string         `json:"description"`
	Entries     []VoucherEntry `json:"entries"`
}

// CreateAccountDTO represents the data transfer object for creating an account
type CreateAccountDTO struct {
	Code             string `json:"code"`
	Name             string `json:"name"`
	AccountType      string `json:"account_type"`
	BalanceDirection string `json:"balance_direction"`
	ParentCode       string `json:"parent_code"`
	IsLeaf           bool   `json:"is_leaf"`
}

// PurchaseOrder represents a purchase order
type PurchaseOrder struct {
	ID           int64               `db:"id" json:"id"`
	OrderNo      string              `db:"order_no" json:"order_no"`
	SupplierCode string              `db:"supplier_code" json:"supplier_code"`
	SupplierName string              `db:"supplier_name" json:"supplier_name"`
	OrderDate    time.Time           `db:"order_date" json:"order_date"`
	DeliveryDate time.Time           `db:"delivery_date" json:"delivery_date"`
	TotalAmount  float64             `db:"total_amount" json:"total_amount"`
	Status       string              `db:"status" json:"status"` // DRAFT, APPROVED, RECEIVED, CLOSED
	Description  string              `db:"description" json:"description"`
	Items        []PurchaseOrderItem `json:"items"`
	CreatedAt    time.Time           `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time           `db:"updated_at" json:"updated_at"`
}

// PurchaseOrderItem represents a purchase order item
type PurchaseOrderItem struct {
	ID          int64     `db:"id" json:"id"`
	OrderID     int64     `db:"order_id" json:"order_id"`
	ProductCode string    `db:"product_code" json:"product_code"`
	ProductName string    `db:"product_name" json:"product_name"`
	Unit        string    `db:"unit" json:"unit"`
	Quantity    float64   `db:"quantity" json:"quantity"`
	UnitPrice   float64   `db:"unit_price" json:"unit_price"`
	Amount      float64   `db:"amount" json:"amount"`
	ReceivedQty float64   `db:"received_qty" json:"received_qty"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// CreatePurchaseOrderDTO represents the data transfer object for creating a purchase order
type CreatePurchaseOrderDTO struct {
	SupplierCode string              `json:"supplier_code"`
	SupplierName string              `json:"supplier_name"`
	OrderDate    time.Time           `json:"order_date"`
	DeliveryDate time.Time           `json:"delivery_date"`
	Description  string              `json:"description"`
	Items        []PurchaseOrderItem `json:"items"`
}

// SalesOrder represents a sales order
type SalesOrder struct {
	ID           int64            `db:"id" json:"id"`
	OrderNo      string           `db:"order_no" json:"order_no"`
	CustomerCode string           `db:"customer_code" json:"customer_code"`
	CustomerName string           `db:"customer_name" json:"customer_name"`
	OrderDate    time.Time        `db:"order_date" json:"order_date"`
	DeliveryDate time.Time        `db:"delivery_date" json:"delivery_date"`
	TotalAmount  float64          `db:"total_amount" json:"total_amount"`
	Status       string           `db:"status" json:"status"` // DRAFT, APPROVED, DELIVERED, CLOSED
	Description  string           `db:"description" json:"description"`
	Items        []SalesOrderItem `json:"items"`
	CreatedAt    time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at" json:"updated_at"`
}

// SalesOrderItem represents a sales order item
type SalesOrderItem struct {
	ID           int64     `db:"id" json:"id"`
	OrderID      int64     `db:"order_id" json:"order_id"`
	ProductCode  string    `db:"product_code" json:"product_code"`
	ProductName  string    `db:"product_name" json:"product_name"`
	Unit         string    `db:"unit" json:"unit"`
	Quantity     float64   `db:"quantity" json:"quantity"`
	UnitPrice    float64   `db:"unit_price" json:"unit_price"`
	Amount       float64   `db:"amount" json:"amount"`
	DeliveredQty float64   `db:"delivered_qty" json:"delivered_qty"`
	Description  string    `db:"description" json:"description"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

// CreateSalesOrderDTO represents the data transfer object for creating a sales order
type CreateSalesOrderDTO struct {
	CustomerCode string           `json:"customer_code"`
	CustomerName string           `json:"customer_name"`
	OrderDate    time.Time        `json:"order_date"`
	DeliveryDate time.Time        `json:"delivery_date"`
	Description  string           `json:"description"`
	Items        []SalesOrderItem `json:"items"`
}

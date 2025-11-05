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
	ParentCode       *string   `db:"parent_code" json:"parent_code"`
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

// ProductCategory represents a product category
type ProductCategory struct {
	ID          int64     `db:"id" json:"id"`
	Code        string    `db:"code" json:"code"`
	Name        string    `db:"name" json:"name"`
	ParentID    *int64    `db:"parent_id" json:"parent_id"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// CreateProductCategoryDTO represents the data transfer object for creating a product category
type CreateProductCategoryDTO struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	ParentID    *int64 `json:"parent_id"`
	Description string `json:"description"`
}

// Product represents a product
type Product struct {
	ID            int64     `db:"id" json:"id"`
	Code          string    `db:"code" json:"code"`
	Name          string    `db:"name" json:"name"`
	CategoryID    *int64    `db:"category_id" json:"category_id"`
	Unit          string    `db:"unit" json:"unit"`
	Specification string    `db:"specification" json:"specification"`
	Description   string    `db:"description" json:"description"`
	Status        string    `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

// CreateProductDTO represents the data transfer object for creating a product
type CreateProductDTO struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	CategoryID    *int64 `json:"category_id"`
	Unit          string `json:"unit"`
	Specification string `json:"specification"`
	Description   string `json:"description"`
}

// Supplier represents a supplier
type Supplier struct {
	ID            int64     `db:"id" json:"id"`
	Code          string    `db:"code" json:"code"`
	Name          string    `db:"name" json:"name"`
	ContactPerson string    `db:"contact_person" json:"contact_person"`
	Phone         string    `db:"phone" json:"phone"`
	Email         string    `db:"email" json:"email"`
	Address       string    `db:"address" json:"address"`
	TaxNumber     string    `db:"tax_number" json:"tax_number"`
	BankName      string    `db:"bank_name" json:"bank_name"`
	BankAccount   string    `db:"bank_account" json:"bank_account"`
	Status        string    `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

// CreateSupplierDTO represents the data transfer object for creating a supplier
type CreateSupplierDTO struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	TaxNumber     string `json:"tax_number"`
	BankName      string `json:"bank_name"`
	BankAccount   string `json:"bank_account"`
}

// Customer represents a customer
type Customer struct {
	ID            int64     `db:"id" json:"id"`
	Code          string    `db:"code" json:"code"`
	Name          string    `db:"name" json:"name"`
	ContactPerson string    `db:"contact_person" json:"contact_person"`
	Phone         string    `db:"phone" json:"phone"`
	Email         string    `db:"email" json:"email"`
	Address       string    `db:"address" json:"address"`
	TaxNumber     string    `db:"tax_number" json:"tax_number"`
	BankName      string    `db:"bank_name" json:"bank_name"`
	BankAccount   string    `db:"bank_account" json:"bank_account"`
	Status        string    `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

// CreateCustomerDTO represents the data transfer object for creating a customer
type CreateCustomerDTO struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	TaxNumber     string `json:"tax_number"`
	BankName      string `json:"bank_name"`
	BankAccount   string `json:"bank_account"`
}

// Inventory represents product inventory
type Inventory struct {
	ID          int64     `db:"id" json:"id"`
	ProductID   int64     `db:"product_id" json:"product_id"`
	Warehouse   string    `db:"warehouse" json:"warehouse"`
	Quantity    float64   `db:"quantity" json:"quantity"`
	UnitCost    float64   `db:"unit_cost" json:"unit_cost"`
	TotalCost   float64   `db:"total_cost" json:"total_cost"`
	CostMethod  string    `db:"cost_method" json:"cost_method"`
	LastUpdated time.Time `db:"last_updated" json:"last_updated"`
}

// InventoryTransaction represents an inventory transaction
type InventoryTransaction struct {
	ID              int64     `db:"id" json:"id"`
	ProductID       int64     `db:"product_id" json:"product_id"`
	TransactionType string    `db:"transaction_type" json:"transaction_type"` // IN, OUT
	Quantity        float64   `db:"quantity" json:"quantity"`
	UnitCost        float64   `db:"unit_cost" json:"unit_cost"`
	TotalCost       float64   `db:"total_cost" json:"total_cost"`
	ReferenceType   string    `db:"reference_type" json:"reference_type"`
	ReferenceID     *int64    `db:"reference_id" json:"reference_id"`
	Warehouse       string    `db:"warehouse" json:"warehouse"`
	TransactionDate time.Time `db:"transaction_date" json:"transaction_date"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}

// InventoryFIFOLayer represents a FIFO cost layer
type InventoryFIFOLayer struct {
	ID           int64     `db:"id" json:"id"`
	ProductID    int64     `db:"product_id" json:"product_id"`
	Quantity     float64   `db:"quantity" json:"quantity"`
	UnitCost     float64   `db:"unit_cost" json:"unit_cost"`
	TotalCost    float64   `db:"total_cost" json:"total_cost"`
	RemainingQty float64   `db:"remaining_quantity" json:"remaining_quantity"`
	ReceiptDate  time.Time `db:"receipt_date" json:"receipt_date"`
	Warehouse    string    `db:"warehouse" json:"warehouse"`
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

// SystemConfig represents system configuration
type SystemConfig struct {
	ID          int64     `db:"id" json:"id"`
	ConfigKey   string    `db:"config_key" json:"config_key"`
	ConfigValue string    `db:"config_value" json:"config_value"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// CreateSystemConfigDTO represents the data transfer object for creating system configuration
type CreateSystemConfigDTO struct {
	ConfigKey   string `json:"config_key"`
	ConfigValue string `json:"config_value"`
	Description string `json:"description"`
}

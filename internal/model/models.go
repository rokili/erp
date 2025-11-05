package model

import (
	"time"
)

// BaseModel 基础模型
type BaseModel struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Account 会计科目
type Account struct {
	BaseModel
	Code             string  `db:"code" json:"code"`                           // 科目编码
	Name             string  `db:"name" json:"name"`                           // 科目名称
	AccountType      string  `db:"account_type" json:"account_type"`           // 科目类型 (ASSET, LIABILITY, EQUITY, INCOME, EXPENSE)
	BalanceDirection string  `db:"balance_direction" json:"balance_direction"` // 余额方向 (DEBIT, CREDIT)
	ParentCode       *string `db:"parent_code" json:"parent_code"`             // 上级科目编码
	IsLeaf           bool    `db:"is_leaf" json:"is_leaf"`                     // 是否末级科目
}

// Voucher 凭证
type Voucher struct {
	BaseModel
	VoucherDate time.Time      `db:"voucher_date" json:"voucher_date"` // 凭证日期
	Description string         `db:"description" json:"description"`   // 摘要
	Status      string         `db:"status" json:"status"`             // 状态 (DRAFT, APPROVED, CLOSED)
	Entries     []VoucherEntry `json:"entries"`                        // 凭证明细
}

// VoucherEntry 凭证明细
type VoucherEntry struct {
	ID           int     `db:"id" json:"id"`
	VoucherID    int     `db:"voucher_id" json:"voucher_id"`       // 凭证ID
	AccountCode  string  `db:"account_code" json:"account_code"`   // 科目编码
	DebitAmount  float64 `db:"debit_amount" json:"debit_amount"`   // 借方金额
	CreditAmount float64 `db:"credit_amount" json:"credit_amount"` // 贷方金额
	Description  string  `db:"description" json:"description"`     // 摘要
}

// Category 商品分类
type Category struct {
	BaseModel
	Code        string `db:"code" json:"code"`               // 分类编码
	Name        string `db:"name" json:"name"`               // 分类名称
	Description string `db:"description" json:"description"` // 描述
}

// Product 商品
type Product struct {
	BaseModel
	Code          string `db:"code" json:"code"`                   // 商品编码
	Name          string `db:"name" json:"name"`                   // 商品名称
	CategoryID    int    `db:"category_id" json:"category_id"`     // 分类ID
	Unit          string `db:"unit" json:"unit"`                   // 单位
	Specification string `db:"specification" json:"specification"` // 规格
	Description   string `db:"description" json:"description"`     // 描述
}

// Supplier 供应商
type Supplier struct {
	BaseModel
	Code          string `db:"code" json:"code"`                     // 供应商编码
	Name          string `db:"name" json:"name"`                     // 供应商名称
	ContactPerson string `db:"contact_person" json:"contact_person"` // 联系人
	Phone         string `db:"phone" json:"phone"`                   // 电话
	Email         string `db:"email" json:"email"`                   // 邮箱
	Address       string `db:"address" json:"address"`               // 地址
	TaxNumber     string `db:"tax_number" json:"tax_number"`         // 税号
	BankName      string `db:"bank_name" json:"bank_name"`           // 开户银行
	BankAccount   string `db:"bank_account" json:"bank_account"`     // 银行账户
}

// Customer 客户
type Customer struct {
	BaseModel
	Code          string `db:"code" json:"code"`                     // 客户编码
	Name          string `db:"name" json:"name"`                     // 客户名称
	ContactPerson string `db:"contact_person" json:"contact_person"` // 联系人
	Phone         string `db:"phone" json:"phone"`                   // 电话
	Email         string `db:"email" json:"email"`                   // 邮箱
	Address       string `db:"address" json:"address"`               // 地址
	TaxNumber     string `db:"tax_number" json:"tax_number"`         // 税号
	BankName      string `db:"bank_name" json:"bank_name"`           // 开户银行
	BankAccount   string `db:"bank_account" json:"bank_account"`     // 银行账户
}

// PurchaseOrder 采购订单
type PurchaseOrder struct {
	BaseModel
	SupplierCode string              `db:"supplier_code" json:"supplier_code"` // 供应商编码
	SupplierName string              `db:"supplier_name" json:"supplier_name"` // 供应商名称
	OrderDate    time.Time           `db:"order_date" json:"order_date"`       // 订单日期
	DeliveryDate time.Time           `db:"delivery_date" json:"delivery_date"` // 交货日期
	Description  string              `db:"description" json:"description"`     // 描述
	Status       string              `db:"status" json:"status"`               // 状态 (DRAFT, APPROVED, CLOSED)
	TotalAmount  float64             `db:"total_amount" json:"total_amount"`   // 总金额
	Items        []PurchaseOrderItem `json:"items"`                            // 订单明细
}

// PurchaseOrderItem 采购订单明细
type PurchaseOrderItem struct {
	ID              int     `db:"id" json:"id"`
	PurchaseOrderID int     `db:"purchase_order_id" json:"purchase_order_id"` // 采购订单ID
	ProductCode     string  `db:"product_code" json:"product_code"`           // 商品编码
	ProductName     string  `db:"product_name" json:"product_name"`           // 商品名称
	Unit            string  `db:"unit" json:"unit"`                           // 单位
	Quantity        float64 `db:"quantity" json:"quantity"`                   // 数量
	UnitPrice       float64 `db:"unit_price" json:"unit_price"`               // 单价
	Amount          float64 `db:"amount" json:"amount"`                       // 金额
	Description     string  `db:"description" json:"description"`             // 描述
}

// SalesOrder 销售订单
type SalesOrder struct {
	BaseModel
	CustomerCode string           `db:"customer_code" json:"customer_code"` // 客户编码
	CustomerName string           `db:"customer_name" json:"customer_name"` // 客户名称
	OrderDate    time.Time        `db:"order_date" json:"order_date"`       // 订单日期
	DeliveryDate time.Time        `db:"delivery_date" json:"delivery_date"` // 交货日期
	Description  string           `db:"description" json:"description"`     // 描述
	Status       string           `db:"status" json:"status"`               // 状态 (DRAFT, APPROVED, CLOSED)
	TotalAmount  float64          `db:"total_amount" json:"total_amount"`   // 总金额
	Items        []SalesOrderItem `json:"items"`                            // 订单明细
}

// SalesOrderItem 销售订单明细
type SalesOrderItem struct {
	ID           int     `db:"id" json:"id"`
	SalesOrderID int     `db:"sales_order_id" json:"sales_order_id"` // 销售订单ID
	ProductCode  string  `db:"product_code" json:"product_code"`     // 商品编码
	ProductName  string  `db:"product_name" json:"product_name"`     // 商品名称
	Unit         string  `db:"unit" json:"unit"`                     // 单位
	Quantity     float64 `db:"quantity" json:"quantity"`             // 数量
	UnitPrice    float64 `db:"unit_price" json:"unit_price"`         // 单价
	Amount       float64 `db:"amount" json:"amount"`                 // 金额
	Description  string  `db:"description" json:"description"`       // 描述
}

// User 用户
type User struct {
	BaseModel
	Username string `db:"username" json:"username"` // 用户名
	Password string `db:"password" json:"-"`        // 密码（不返回给前端）
	Name     string `db:"name" json:"name"`         // 姓名
	Email    string `db:"email" json:"email"`       // 邮箱
	Phone    string `db:"phone" json:"phone"`       // 电话
	Status   string `db:"status" json:"status"`     // 状态 (ACTIVE, INACTIVE)
}

// Role 角色
type Role struct {
	BaseModel
	Name        string `db:"name" json:"name"`               // 角色名称
	Description string `db:"description" json:"description"` // 描述
}

// UserRole 用户角色关联
type UserRole struct {
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"` // 用户ID
	RoleID int `db:"role_id" json:"role_id"` // 角色ID
}

// Permission 权限
type Permission struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`               // 权限名称
	Description string `db:"description" json:"description"` // 描述
	Resource    string `db:"resource" json:"resource"`       // 资源
	Action      string `db:"action" json:"action"`           // 操作
}

// RolePermission 角色权限关联
type RolePermission struct {
	ID           int `db:"id" json:"id"`
	RoleID       int `db:"role_id" json:"role_id"`             // 角色ID
	PermissionID int `db:"permission_id" json:"permission_id"` // 权限ID
}

// SystemStatusResponse 系统状态响应
type SystemStatusResponse struct {
	IsOpened bool   `json:"is_opened"`
	Message  string `json:"message"`
}

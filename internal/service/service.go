package service

import (
	"erp/internal/model"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

// Repository 接口定义
type Repository interface {
	// Voucher operations
	CreateVoucher(voucher *model.Voucher) error
	GetVoucher(id int64) (*model.Voucher, error)
	ListVouchers(limit, offset int) ([]*model.Voucher, error)

	// Account operations
	CreateAccount(account *model.Account) error
	GetAccount(code string) (*model.Account, error)
	ListAccounts() ([]*model.Account, error)

	// ProductCategory operations
	CreateProductCategory(category *model.ProductCategory) error
	GetProductCategory(id int64) (*model.ProductCategory, error)
	ListProductCategories() ([]*model.ProductCategory, error)

	// Product operations
	CreateProduct(product *model.Product) error
	GetProduct(id int64) (*model.Product, error)
	ListProducts() ([]*model.Product, error)

	// Supplier operations
	CreateSupplier(supplier *model.Supplier) error
	GetSupplier(id int64) (*model.Supplier, error)
	ListSuppliers() ([]*model.Supplier, error)

	// Customer operations
	CreateCustomer(customer *model.Customer) error
	GetCustomer(id int64) (*model.Customer, error)
	ListCustomers() ([]*model.Customer, error)

	// SystemConfig operations
	GetSystemConfig(key string) (*model.SystemConfig, error)
	SetSystemConfig(key, value, description string) error
	IsSystemOpened() (bool, error)
	OpenSystem() error

	// PurchaseOrder operations
	CreatePurchaseOrder(order *model.PurchaseOrder) error
	GetPurchaseOrder(id int64) (*model.PurchaseOrder, error)
	ListPurchaseOrders(limit, offset int) ([]*model.PurchaseOrder, error)
	UpdatePurchaseOrderStatus(id int64, status string) error

	// SalesOrder operations
	CreateSalesOrder(order *model.SalesOrder) error
	GetSalesOrder(id int64) (*model.SalesOrder, error)
	ListSalesOrders(limit, offset int) ([]*model.SalesOrder, error)
	UpdateSalesOrderStatus(id int64, status string) error
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

// Voucher operations
func (s *Service) CreateVoucher(dto *model.CreateVoucherDTO) (*model.Voucher, error) {
	// 检查系统是否已开账
	isOpened, err := s.repo.IsSystemOpened()
	if err != nil {
		return nil, err
	}

	// 验证凭证数据
	if err := s.validateVoucher(dto, isOpened); err != nil {
		return nil, err
	}

	// 计算总金额
	totalAmount := s.calculateVoucherTotal(dto.Entries)

	// 生成凭证号
	voucherNo := s.generateVoucherNo()

	// 构建凭证对象
	voucher := &model.Voucher{
		VoucherNo:   voucherNo,
		VoucherDate: dto.VoucherDate,
		Description: dto.Description,
		TotalAmount: totalAmount,
		Entries:     dto.Entries,
		CreatedAt:   time.Now(),
	}

	// 保存到数据库
	if err := s.repo.CreateVoucher(voucher); err != nil {
		return nil, err
	}

	return voucher, nil
}

func (s *Service) validateVoucher(dto *model.CreateVoucherDTO, systemOpened bool) error {
	// 检查必需字段
	if dto.VoucherDate.IsZero() {
		return errors.New("凭证日期不能为空")
	}

	if len(dto.Entries) == 0 {
		return errors.New("凭证至少需要一个分录")
	}

	// 如果系统已开账，需要验证借贷平衡
	if systemOpened {
		// 验证凭证明细
		totalDebit := 0.0
		totalCredit := 0.0
		for _, entry := range dto.Entries {
			if entry.AccountCode == "" {
				return errors.New("科目代码不能为空")
			}

			if entry.DebitAmount < 0 || entry.CreditAmount < 0 {
				return errors.New("借贷金额不能为负数")
			}

			if entry.DebitAmount > 0 && entry.CreditAmount > 0 {
				return errors.New("同一分录不能同时有借贷金额")
			}

			if entry.DebitAmount == 0 && entry.CreditAmount == 0 {
				return errors.New("借贷金额不能同时为零")
			}

			totalDebit += entry.DebitAmount
			totalCredit += entry.CreditAmount
		}

		// 检查借贷平衡
		if totalDebit != totalCredit {
			return errors.New("借贷金额不平衡")
		}
	} else {
		// 系统未开账时，允许单方面调整
		for _, entry := range dto.Entries {
			if entry.AccountCode == "" {
				return errors.New("科目代码不能为空")
			}

			if entry.DebitAmount < 0 || entry.CreditAmount < 0 {
				return errors.New("金额不能为负数")
			}

			if entry.DebitAmount > 0 && entry.CreditAmount > 0 {
				return errors.New("同一分录不能同时有借贷金额")
			}
		}
	}

	return nil
}

func (s *Service) calculateVoucherTotal(entries []model.VoucherEntry) float64 {
	total := 0.0
	for _, entry := range entries {
		total += entry.DebitAmount + entry.CreditAmount
	}
	return total
}

func (s *Service) generateVoucherNo() string {
	// 生成凭证号：V+日期+序号
	now := time.Now()
	return "V" + now.Format("20060102") + uuid.New().String()[:8]
}

func (s *Service) GetVoucher(id int64) (*model.Voucher, error) {
	return s.repo.GetVoucher(id)
}

func (s *Service) ListVouchers(limit, offset int) ([]*model.Voucher, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.ListVouchers(limit, offset)
}

// Account operations
func (s *Service) CreateAccount(dto *model.CreateAccountDTO) (*model.Account, error) {
	// 验证科目数据
	if err := s.validateAccount(dto); err != nil {
		return nil, err
	}

	// 构建科目对象
	account := &model.Account{
		Code:             dto.Code,
		Name:             dto.Name,
		AccountType:      dto.AccountType,
		BalanceDirection: dto.BalanceDirection,
		ParentCode:       &dto.ParentCode,
		IsLeaf:           dto.IsLeaf,
		Status:           "ACTIVE",
		CreatedAt:        time.Now(),
	}

	// 如果 ParentCode 为空字符串，则设置为 nil
	if dto.ParentCode == "" {
		account.ParentCode = nil
	}

	// 保存到数据库
	if err := s.repo.CreateAccount(account); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *Service) validateAccount(dto *model.CreateAccountDTO) error {
	// 检查必需字段
	if dto.Code == "" {
		return errors.New("科目代码不能为空")
	}

	if dto.Name == "" {
		return errors.New("科目名称不能为空")
	}

	if dto.AccountType == "" {
		return errors.New("科目类型不能为空")
	}

	if dto.BalanceDirection == "" {
		return errors.New("余额方向不能为空")
	}

	return nil
}

func (s *Service) GetAccount(code string) (*model.Account, error) {
	return s.repo.GetAccount(code)
}

func (s *Service) ListAccounts() ([]*model.Account, error) {
	return s.repo.ListAccounts()
}

// ProductCategory operations
func (s *Service) CreateProductCategory(dto *model.CreateProductCategoryDTO) (*model.ProductCategory, error) {
	// 验证商品分类数据
	if err := s.validateProductCategory(dto); err != nil {
		return nil, err
	}

	// 构建商品分类对象
	category := &model.ProductCategory{
		Code:        dto.Code,
		Name:        dto.Name,
		ParentID:    dto.ParentID,
		Description: dto.Description,
		CreatedAt:   time.Now(),
	}

	// 保存到数据库
	if err := s.repo.CreateProductCategory(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) validateProductCategory(dto *model.CreateProductCategoryDTO) error {
	// 检查必需字段
	if dto.Name == "" {
		return errors.New("分类名称不能为空")
	}

	return nil
}

func (s *Service) GetProductCategory(id int64) (*model.ProductCategory, error) {
	return s.repo.GetProductCategory(id)
}

func (s *Service) ListProductCategories() ([]*model.ProductCategory, error) {
	return s.repo.ListProductCategories()
}

// Product operations
func (s *Service) CreateProduct(dto *model.CreateProductDTO) (*model.Product, error) {
	// 验证商品数据
	if err := s.validateProduct(dto); err != nil {
		return nil, err
	}

	// 构建商品对象
	product := &model.Product{
		Code:          dto.Code,
		Name:          dto.Name,
		CategoryID:    dto.CategoryID,
		Unit:          dto.Unit,
		Specification: dto.Specification,
		Description:   dto.Description,
		Status:        "ACTIVE",
		CreatedAt:     time.Now(),
	}

	// 保存到数据库
	if err := s.repo.CreateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) validateProduct(dto *model.CreateProductDTO) error {
	// 检查必需字段
	if dto.Name == "" {
		return errors.New("商品名称不能为空")
	}

	return nil
}

func (s *Service) GetProduct(id int64) (*model.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *Service) ListProducts() ([]*model.Product, error) {
	return s.repo.ListProducts()
}

// Supplier operations
func (s *Service) CreateSupplier(dto *model.CreateSupplierDTO) (*model.Supplier, error) {
	// 验证供应商数据
	if err := s.validateSupplier(dto); err != nil {
		return nil, err
	}

	// 构建供应商对象
	supplier := &model.Supplier{
		Code:          dto.Code,
		Name:          dto.Name,
		ContactPerson: dto.ContactPerson,
		Phone:         dto.Phone,
		Email:         dto.Email,
		Address:       dto.Address,
		TaxNumber:     dto.TaxNumber,
		BankName:      dto.BankName,
		BankAccount:   dto.BankAccount,
		Status:        "ACTIVE",
		CreatedAt:     time.Now(),
	}

	// 保存到数据库
	if err := s.repo.CreateSupplier(supplier); err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *Service) validateSupplier(dto *model.CreateSupplierDTO) error {
	// 检查必需字段
	if dto.Name == "" {
		return errors.New("供应商名称不能为空")
	}

	return nil
}

func (s *Service) GetSupplier(id int64) (*model.Supplier, error) {
	return s.repo.GetSupplier(id)
}

func (s *Service) ListSuppliers() ([]*model.Supplier, error) {
	return s.repo.ListSuppliers()
}

// Customer operations
func (s *Service) CreateCustomer(dto *model.CreateCustomerDTO) (*model.Customer, error) {
	// 验证客户数据
	if err := s.validateCustomer(dto); err != nil {
		return nil, err
	}

	// 构建客户对象
	customer := &model.Customer{
		Code:          dto.Code,
		Name:          dto.Name,
		ContactPerson: dto.ContactPerson,
		Phone:         dto.Phone,
		Email:         dto.Email,
		Address:       dto.Address,
		TaxNumber:     dto.TaxNumber,
		BankName:      dto.BankName,
		BankAccount:   dto.BankAccount,
		Status:        "ACTIVE",
		CreatedAt:     time.Now(),
	}

	// 保存到数据库
	if err := s.repo.CreateCustomer(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) validateCustomer(dto *model.CreateCustomerDTO) error {
	// 检查必需字段
	if dto.Name == "" {
		return errors.New("客户名称不能为空")
	}

	return nil
}

func (s *Service) GetCustomer(id int64) (*model.Customer, error) {
	return s.repo.GetCustomer(id)
}

func (s *Service) ListCustomers() ([]*model.Customer, error) {
	return s.repo.ListCustomers()
}

// System operations
func (s *Service) IsSystemOpened() (bool, error) {
	return s.repo.IsSystemOpened()
}

func (s *Service) OpenSystem() error {
	// 检查系统是否已经开账
	isOpened, err := s.repo.IsSystemOpened()
	if err != nil {
		return err
	}

	if isOpened {
		return errors.New("系统已经开账，不能重复开账")
	}

	return s.repo.OpenSystem()
}

// PurchaseOrder operations
func (s *Service) CreatePurchaseOrder(dto *model.CreatePurchaseOrderDTO) (*model.PurchaseOrder, error) {
	// 验证采购订单数据
	if err := s.validatePurchaseOrder(dto); err != nil {
		return nil, err
	}

	// 计算总金额
	totalAmount := s.calculatePurchaseOrderTotal(dto.Items)

	// 生成订单号
	orderNo := s.generatePurchaseOrderNo()

	// 构建采购订单对象
	order := &model.PurchaseOrder{
		OrderNo:      orderNo,
		SupplierCode: dto.SupplierCode,
		SupplierName: dto.SupplierName,
		OrderDate:    dto.OrderDate,
		DeliveryDate: dto.DeliveryDate,
		TotalAmount:  totalAmount,
		Status:       "DRAFT",
		Description:  dto.Description,
		Items:        dto.Items,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// 保存到数据库
	if err := s.repo.CreatePurchaseOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Service) validatePurchaseOrder(dto *model.CreatePurchaseOrderDTO) error {
	// 检查必需字段
	if dto.SupplierCode == "" {
		return errors.New("供应商代码不能为空")
	}

	if dto.SupplierName == "" {
		return errors.New("供应商名称不能为空")
	}

	if dto.OrderDate.IsZero() {
		return errors.New("订单日期不能为空")
	}

	if len(dto.Items) == 0 {
		return errors.New("订单至少需要一个商品")
	}

	// 验证订单明细
	for _, item := range dto.Items {
		if item.ProductCode == "" {
			return errors.New("商品代码不能为空")
		}

		if item.ProductName == "" {
			return errors.New("商品名称不能为空")
		}

		if item.Unit == "" {
			return errors.New("单位不能为空")
		}

		if item.Quantity <= 0 {
			return errors.New("数量必须大于0")
		}

		if item.UnitPrice <= 0 {
			return errors.New("单价必须大于0")
		}
	}

	return nil
}

func (s *Service) calculatePurchaseOrderTotal(items []model.PurchaseOrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Amount
	}
	return total
}

func (s *Service) generatePurchaseOrderNo() string {
	// 生成订单号：PO+日期+序号
	now := time.Now()
	return "PO" + now.Format("20060102") + uuid.New().String()[:8]
}

func (s *Service) GetPurchaseOrder(id int64) (*model.PurchaseOrder, error) {
	return s.repo.GetPurchaseOrder(id)
}

func (s *Service) ListPurchaseOrders(limit, offset int) ([]*model.PurchaseOrder, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.ListPurchaseOrders(limit, offset)
}

func (s *Service) ApprovePurchaseOrder(id int64) error {
	return s.repo.UpdatePurchaseOrderStatus(id, "APPROVED")
}

func (s *Service) ClosePurchaseOrder(id int64) error {
	return s.repo.UpdatePurchaseOrderStatus(id, "CLOSED")
}

// SalesOrder operations
func (s *Service) CreateSalesOrder(dto *model.CreateSalesOrderDTO) (*model.SalesOrder, error) {
	// 验证销售订单数据
	if err := s.validateSalesOrder(dto); err != nil {
		return nil, err
	}

	// 计算总金额
	totalAmount := s.calculateSalesOrderTotal(dto.Items)

	// 生成订单号
	orderNo := s.generateSalesOrderNo()

	// 构建销售订单对象
	order := &model.SalesOrder{
		OrderNo:      orderNo,
		CustomerCode: dto.CustomerCode,
		CustomerName: dto.CustomerName,
		OrderDate:    dto.OrderDate,
		DeliveryDate: dto.DeliveryDate,
		TotalAmount:  totalAmount,
		Status:       "DRAFT",
		Description:  dto.Description,
		Items:        dto.Items,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// 保存到数据库
	if err := s.repo.CreateSalesOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Service) validateSalesOrder(dto *model.CreateSalesOrderDTO) error {
	// 检查必需字段
	if dto.CustomerCode == "" {
		return errors.New("客户代码不能为空")
	}

	if dto.CustomerName == "" {
		return errors.New("客户名称不能为空")
	}

	if dto.OrderDate.IsZero() {
		return errors.New("订单日期不能为空")
	}

	if len(dto.Items) == 0 {
		return errors.New("订单至少需要一个商品")
	}

	// 验证订单明细
	for _, item := range dto.Items {
		if item.ProductCode == "" {
			return errors.New("商品代码不能为空")
		}

		if item.ProductName == "" {
			return errors.New("商品名称不能为空")
		}

		if item.Unit == "" {
			return errors.New("单位不能为空")
		}

		if item.Quantity <= 0 {
			return errors.New("数量必须大于0")
		}

		if item.UnitPrice <= 0 {
			return errors.New("单价必须大于0")
		}
	}

	return nil
}

func (s *Service) calculateSalesOrderTotal(items []model.SalesOrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Amount
	}
	return total
}

func (s *Service) generateSalesOrderNo() string {
	// 生成订单号：SO+日期+序号
	now := time.Now()
	return "SO" + now.Format("20060102") + uuid.New().String()[:8]
}

func (s *Service) GetSalesOrder(id int64) (*model.SalesOrder, error) {
	return s.repo.GetSalesOrder(id)
}

func (s *Service) ListSalesOrders(limit, offset int) ([]*model.SalesOrder, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.ListSalesOrders(limit, offset)
}

func (s *Service) ApproveSalesOrder(id int64) error {
	return s.repo.UpdateSalesOrderStatus(id, "APPROVED")
}

func (s *Service) CloseSalesOrder(id int64) error {
	return s.repo.UpdateSalesOrderStatus(id, "CLOSED")
}

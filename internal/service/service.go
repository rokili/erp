package service

import (
	"erp/internal/model"
	"erp/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// Voucher operations
func (s *Service) CreateVoucher(dto *model.CreateVoucherDTO) (*model.Voucher, error) {
	// 验证凭证数据
	if err := s.validateVoucher(dto); err != nil {
		return nil, err
	}

	// 计算总金额
	totalAmount := s.calculateTotalAmount(dto.Entries)

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

func (s *Service) validateVoucher(dto *model.CreateVoucherDTO) error {
	// 检查必需字段
	if dto.VoucherDate.IsZero() {
		return errors.New("凭证日期不能为空")
	}

	if len(dto.Entries) < 2 {
		return errors.New("凭证至少需要两个分录")
	}

	// 验证借贷平衡
	var totalDebit, totalCredit float64
	for _, entry := range dto.Entries {
		if entry.DebitAmount < 0 || entry.CreditAmount < 0 {
			return errors.New("金额不能为负数")
		}

		if entry.DebitAmount > 0 && entry.CreditAmount > 0 {
			return errors.New("同一分录不能同时有借方和贷方金额")
		}

		totalDebit += entry.DebitAmount
		totalCredit += entry.CreditAmount
	}

	if totalDebit != totalCredit {
		return errors.New("凭证借贷不平衡")
	}

	return nil
}

func (s *Service) calculateTotalAmount(entries []model.VoucherEntry) float64 {
	var total float64
	for _, entry := range entries {
		if entry.DebitAmount > total {
			total = entry.DebitAmount
		}
		if entry.CreditAmount > total {
			total = entry.CreditAmount
		}
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
		ParentCode:       dto.ParentCode,
		IsLeaf:           dto.IsLeaf,
		Status:           "ACTIVE",
		CreatedAt:        time.Now(),
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

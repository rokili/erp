package service

import (
	"erp/internal/model"
	"erp/internal/repository"
)

type FinanceService struct{}

func NewFinanceService() *FinanceService {
	return &FinanceService{}
}

// CreateAccount 创建会计科目
func (s *FinanceService) CreateAccount(account *model.Account) error {
	// TODO: 实现创建会计科目逻辑
	return repository.CreateAccount(account)
}

// GetAccount 获取会计科目
func (s *FinanceService) GetAccount(code string) (*model.Account, error) {
	// TODO: 实现获取会计科目逻辑
	return repository.GetAccountByCode(code)
}

// ListAccounts 获取所有会计科目
func (s *FinanceService) ListAccounts() ([]*model.Account, error) {
	// TODO: 实现获取所有会计科目逻辑
	return repository.ListAccounts()
}

// CreateVoucher 创建凭证
func (s *FinanceService) CreateVoucher(voucher *model.Voucher) error {
	// TODO: 实现创建凭证逻辑
	return nil
}

// GetVoucher 获取凭证
func (s *FinanceService) GetVoucher(id int) (*model.Voucher, error) {
	// TODO: 实现获取凭证逻辑
	return &model.Voucher{}, nil
}

// ListVouchers 获取凭证列表
func (s *FinanceService) ListVouchers() ([]*model.Voucher, error) {
	// TODO: 实现获取凭证列表逻辑
	return []*model.Voucher{}, nil
}

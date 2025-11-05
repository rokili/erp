package repository

import (
	"erp/internal/model"
	"log"
)

// CreateAccount 创建会计科目
func CreateAccount(account *model.Account) error {
	db := GetDB()
	query := `INSERT INTO accounts (code, name, account_type, balance_direction, parent_code, is_leaf) 
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, account.Code, account.Name, account.AccountType, account.BalanceDirection, account.ParentCode, account.IsLeaf)
	return err
}

// GetAccountByCode 根据编码获取会计科目
func GetAccountByCode(code string) (*model.Account, error) {
	db := GetDB()
	query := `SELECT id, code, name, account_type, balance_direction, parent_code, is_leaf, created_at, updated_at 
	          FROM accounts WHERE code = $1`
	account := &model.Account{}
	err := db.Get(account, query, code)
	if err != nil {
		log.Printf("GetAccountByCode error: %v", err)
		return nil, err
	}
	return account, nil
}

// ListAccounts 获取所有会计科目
func ListAccounts() ([]*model.Account, error) {
	db := GetDB()
	query := `SELECT id, code, name, account_type, balance_direction, parent_code, is_leaf, created_at, updated_at 
	          FROM accounts ORDER BY code`
	accounts := []*model.Account{}
	err := db.Select(&accounts, query)
	if err != nil {
		log.Printf("ListAccounts error: %v", err)
		return nil, err
	}
	return accounts, nil
}

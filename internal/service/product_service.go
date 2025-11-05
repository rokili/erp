package service

import (
	"erp/internal/model"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

// CreateCategory 创建商品分类
func (s *ProductService) CreateCategory(category *model.Category) error {
	// TODO: 实现创建商品分类逻辑
	return nil
}

// GetCategory 获取商品分类
func (s *ProductService) GetCategory(id int) (*model.Category, error) {
	// TODO: 实现获取商品分类逻辑
	return &model.Category{}, nil
}

// ListCategories 获取所有商品分类
func (s *ProductService) ListCategories() ([]*model.Category, error) {
	// TODO: 实现获取所有商品分类逻辑
	return []*model.Category{}, nil
}

// CreateProduct 创建商品
func (s *ProductService) CreateProduct(product *model.Product) error {
	// TODO: 实现创建商品逻辑
	return nil
}

// GetProduct 获取商品
func (s *ProductService) GetProduct(id int) (*model.Product, error) {
	// TODO: 实现获取商品逻辑
	return &model.Product{}, nil
}

// ListProducts 获取所有商品
func (s *ProductService) ListProducts() ([]*model.Product, error) {
	// TODO: 实现获取所有商品逻辑
	return []*model.Product{}, nil
}

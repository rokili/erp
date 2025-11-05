package service

import (
	"erp/internal/model"
)

type SystemService struct{}

func NewSystemService() *SystemService {
	return &SystemService{}
}

// GetSystemStatus 获取系统状态
func (s *SystemService) GetSystemStatus() *model.SystemStatusResponse {
	// TODO: 实现获取系统状态逻辑
	return &model.SystemStatusResponse{
		IsOpened: false,
		Message:  "系统未开账",
	}
}

// OpenSystem 系统开账
func (s *SystemService) OpenSystem() error {
	// TODO: 实现系统开账逻辑
	return nil
}

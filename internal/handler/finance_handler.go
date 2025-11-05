package handler

import (
	"erp/internal/model"
	"erp/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAccount 创建会计科目
func CreateAccount(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	// TODO: 实现创建会计科目逻辑
	// financeService := service.NewFinanceService()
	// err := financeService.CreateAccount(&account)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "创建会计科目失败",
	// 	})
	// 	return
	// }
	c.JSON(http.StatusCreated, gin.H{
		"message": "会计科目创建成功",
		"account": account,
	})
}

// GetAccount 获取会计科目
func GetAccount(c *gin.Context) {
	code := c.Param("code")
	// TODO: 实现获取会计科目逻辑
	financeService := service.NewFinanceService()
	account, err := financeService.GetAccount(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "会计科目不存在",
		})
		return
	}
	c.JSON(http.StatusOK, account)
}

// ListAccounts 获取所有会计科目
func ListAccounts(c *gin.Context) {
	// TODO: 实现获取所有会计科目逻辑
	financeService := service.NewFinanceService()
	accounts, err := financeService.ListAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取会计科目列表失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

// CreateVoucher 创建凭证
func CreateVoucher(c *gin.Context) {
	// TODO: 实现创建凭证逻辑
	c.JSON(http.StatusCreated, gin.H{
		"message": "凭证创建成功",
	})
}

// GetVoucher 获取凭证
func GetVoucher(c *gin.Context) {
	id := c.Param("id")
	// TODO: 实现获取凭证逻辑
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// ListVouchers 获取凭证列表
func ListVouchers(c *gin.Context) {
	// TODO: 实现获取凭证列表逻辑
	c.JSON(http.StatusOK, gin.H{
		"vouchers": []string{},
	})
}

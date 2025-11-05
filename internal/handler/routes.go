package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 系统管理路由
	system := r.Group("/api/system")
	{
		system.GET("/status", GetSystemStatus)
		system.POST("/open", OpenSystem)
	}

	// 财务管理路由
	finance := r.Group("/api/finance")
	{
		// 会计科目
		accounts := finance.Group("/accounts")
		{
			accounts.POST("", CreateAccount)
			accounts.GET("/:code", GetAccount)
			accounts.GET("", ListAccounts)
		}

		// 凭证
		vouchers := finance.Group("/vouchers")
		{
			vouchers.POST("", CreateVoucher)
			vouchers.GET("/:id", GetVoucher)
			vouchers.GET("", ListVouchers)
		}
	}

	// 商品管理路由
	products := r.Group("/api")
	{
		// 商品分类
		categories := products.Group("/categories")
		{
			categories.POST("", CreateCategory)
			categories.GET("/:id", GetCategory)
			categories.GET("", ListCategories)
		}

		// 商品
		products := products.Group("/products")
		{
			products.POST("", CreateProduct)
			products.GET("/:id", GetProduct)
			products.GET("", ListProducts)
		}
	}

	// 采购管理路由
	purchase := r.Group("/api/purchase")
	{
		// 供应商
		suppliers := purchase.Group("/suppliers")
		{
			suppliers.POST("", CreateSupplier)
			suppliers.GET("/:id", GetSupplier)
			suppliers.GET("", ListSuppliers)
		}

		// 采购订单
		orders := purchase.Group("/orders")
		{
			orders.POST("", CreatePurchaseOrder)
			orders.GET("/:id", GetPurchaseOrder)
			orders.GET("", ListPurchaseOrders)
			orders.PUT("/:id/approve", ApprovePurchaseOrder)
			orders.PUT("/:id/close", ClosePurchaseOrder)
		}
	}

	// 销售管理路由
	sales := r.Group("/api/sales")
	{
		// 客户
		customers := sales.Group("/customers")
		{
			customers.POST("", CreateCustomer)
			customers.GET("/:id", GetCustomer)
			customers.GET("", ListCustomers)
		}

		// 销售订单
		orders := sales.Group("/orders")
		{
			orders.POST("", CreateSalesOrder)
			orders.GET("/:id", GetSalesOrder)
			orders.GET("", ListSalesOrders)
			orders.PUT("/:id/approve", ApproveSalesOrder)
			orders.PUT("/:id/close", CloseSalesOrder)
		}
	}

	// 用户权限管理路由
	usersGroup := r.Group("/api/users")
	{
		usersGroup.POST("", CreateUser)
		usersGroup.GET("/:id", GetUser)
		usersGroup.GET("", ListUsers)
	}

	// 用户角色管理路由
	userRoles := r.Group("/api/user-roles")
	{
		userRoles.GET("/:user_id/roles", GetUserRoles)
		userRoles.POST("/:user_id/roles/:role_id", AssignRoleToUser)
		userRoles.DELETE("/:user_id/roles/:role_id", RemoveRoleFromUser)
		userRoles.GET("/:user_id/permissions", GetUserPermissions)
	}

	// 角色管理路由
	rolesGroup := r.Group("/api/roles")
	{
		rolesGroup.POST("", CreateRole)
		rolesGroup.GET("/:id", GetRole)
		rolesGroup.GET("", ListRoles)
	}

	// 角色权限管理路由
	rolePermissions := r.Group("/api/role-permissions")
	{
		rolePermissions.GET("/:role_id/permissions", GetRolePermissions)
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}

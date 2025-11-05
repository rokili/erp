# ERP系统使用说明

## 启动系统

1. 确保PostgreSQL数据库已安装并运行
2. 初始化数据库（参考README.md中的数据库设置部分）
3. 修改配置文件[config.yaml](file:///d:/soft/pg/erp/config.yaml)以匹配您的数据库设置
4. 运行系统：
   ```
   go run main.go
   ```

## 访问系统

系统启动后，可以通过以下方式访问：

1. **API接口**：http://localhost:8080
2. **Web界面**：http://localhost:8080/web/index.html

## 主要功能模块

### 1. 系统管理
- 查看系统状态
- 执行系统开账操作

### 2. 财务管理
- 管理会计科目
- 录入和查询财务凭证

### 3. 采购管理
- 管理供应商信息
- 创建和管理采购订单

### 4. 销售管理
- 管理客户信息
- 创建和管理销售订单

### 5. 商品管理
- 管理商品分类
- 管理商品信息

### 6. 用户权限管理
- 管理用户和角色
- 分配权限

## API测试

可以使用以下工具测试API：

1. **curl命令**（Linux/Mac）
2. **Postman**
3. **浏览器开发者工具**
4. **系统自带的测试页面**（http://localhost:8080/web/index.html）

## 常见问题

### 1. 数据库连接失败
- 检查PostgreSQL是否正在运行
- 检查[config.yaml](file:///d:/soft/pg/erp/config.yaml)中的数据库连接信息是否正确
- 确保数据库用户具有足够的权限

### 2. 端口被占用
- 修改[config.yaml](file:///d:/soft/pg/erp/config.yaml)中的端口设置
- 或者停止占用8080端口的其他程序

### 3. 编译错误
- 确保已安装Go 1.21或更高版本
- 运行 `go mod tidy` 安装依赖
// 全局变量
let currentTab = 'finance';

// DOM加载完成后执行
document.addEventListener('DOMContentLoaded', function() {
    // 初始化页面
    initPage();
    
    // 绑定事件监听器
    bindEventListeners();
    
    // 检查系统状态
    checkSystemStatus();
});

// 初始化页面
function initPage() {
    // 设置当前日期为默认凭证日期
    const today = new Date().toISOString().split('T')[0];
    document.getElementById('voucherDate').value = today;
}

// 绑定事件监听器
function bindEventListeners() {
    // 导航标签切换
    document.querySelectorAll('.nav-link').forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            switchTab(this.getAttribute('data-tab'));
        });
    });

    // 凭证表单提交
    document.getElementById('voucherForm').addEventListener('submit', function(e) {
        e.preventDefault();
        createVoucher();
    });

    // 添加分录按钮
    document.getElementById('addEntryBtn').addEventListener('click', function() {
        addEntryRow();
    });

    // 删除分录按钮事件委托
    document.getElementById('entriesContainer').addEventListener('click', function(e) {
        if (e.target.classList.contains('remove-entry')) {
            removeEntryRow(e.target);
        }
    });

    // 科目表单提交
    document.getElementById('accountForm').addEventListener('submit', function(e) {
        e.preventDefault();
        createAccount();
    });

    // 商品分类表单提交
    document.getElementById('categoryForm').addEventListener('submit', function(e) {
        e.preventDefault();
        createProductCategory();
    });

    // 商品表单提交
    document.getElementById('productForm').addEventListener('submit', function(e) {
        e.preventDefault();
        createProduct();
    });

    // 供应商表单提交
    document.getElementById('supplierForm').addEventListener('submit', function(e) {
        e.preventDefault();
        createSupplier();
    });

    // 客户表单提交
    document.getElementById('customerForm').addEventListener('submit', function(e) {
        e.preventDefault();
        createCustomer();
    });

    // 开账按钮
    document.getElementById('openSystemBtn').addEventListener('click', function() {
        openSystem();
    });

    // 设置页面开账按钮
    document.getElementById('settingsOpenSystemBtn').addEventListener('click', function() {
        openSystem();
    });
}

// 切换标签页
function switchTab(tabName) {
    // 更新当前标签
    currentTab = tabName;
    
    // 隐藏所有标签内容
    document.querySelectorAll('.tab-content').forEach(section => {
        section.classList.remove('active');
    });
    
    // 显示选中的标签内容
    document.getElementById(tabName).classList.add('active');
    
    // 更新导航链接状态
    document.querySelectorAll('.nav-link').forEach(link => {
        link.classList.remove('active');
    });
    
    document.querySelector(`.nav-link[data-tab="${tabName}"]`).classList.add('active');
}

// 检查系统状态
async function checkSystemStatus() {
    try {
        const response = await fetch('/api/system/status');
        const data = await response.json();
        
        const statusText = data.opened ? '已开账' : '未开账';
        document.getElementById('systemStatus').textContent = `系统状态: ${statusText}`;
        document.getElementById('settingsSystemStatus').textContent = statusText;
        
        // 根据系统状态启用/禁用开账按钮
        const openSystemBtn = document.getElementById('openSystemBtn');
        const settingsOpenSystemBtn = document.getElementById('settingsOpenSystemBtn');
        
        if (data.opened) {
            openSystemBtn.disabled = true;
            settingsOpenSystemBtn.disabled = true;
            openSystemBtn.textContent = '已开账';
            settingsOpenSystemBtn.textContent = '已开账';
        } else {
            openSystemBtn.disabled = false;
            settingsOpenSystemBtn.disabled = false;
            openSystemBtn.textContent = '开账';
            settingsOpenSystemBtn.textContent = '开账';
        }
    } catch (error) {
        console.error('检查系统状态失败:', error);
        showMessage('检查系统状态失败', 'error');
    }
}

// 开账
async function openSystem() {
    if (!confirm('确定要开账吗？开账后将无法再进行单方面调整！')) {
        return;
    }
    
    try {
        const response = await fetch('/api/system/open', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        
        if (response.ok) {
            showMessage('系统开账成功', 'success');
            checkSystemStatus();
        } else {
            const error = await response.json();
            showMessage(`开账失败: ${error.error}`, 'error');
        }
    } catch (error) {
        console.error('开账失败:', error);
        showMessage('开账失败，请稍后重试', 'error');
    }
}

// 创建凭证
async function createVoucher() {
    // 收集表单数据
    const voucherData = {
        voucher_date: document.getElementById('voucherDate').value,
        description: document.getElementById('voucherDescription').value,
        entries: []
    };
    
    // 收集分录数据
    const entryRows = document.querySelectorAll('.entry-row');
    for (let row of entryRows) {
        const accountCode = row.querySelector('.account-code').value;
        const accountName = row.querySelector('.account-name').value;
        const debitAmount = parseFloat(row.querySelector('.debit-amount').value) || 0;
        const creditAmount = parseFloat(row.querySelector('.credit-amount').value) || 0;
        
        if (accountCode && accountName) {
            voucherData.entries.push({
                account_code: accountCode,
                debit_amount: debitAmount,
                credit_amount: creditAmount,
                description: accountName
            });
        }
    }
    
    // 验证数据
    if (voucherData.entries.length === 0) {
        showMessage('至少需要一个分录', 'error');
        return;
    }
    
    try {
        const response = await fetch('/api/finance/vouchers', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(voucherData)
        });
        
        if (response.ok) {
            showMessage('凭证创建成功', 'success');
            document.getElementById('voucherForm').reset();
            document.getElementById('voucherDate').value = new Date().toISOString().split('T')[0];
            // 保留第一个分录行
            const container = document.getElementById('entriesContainer');
            container.innerHTML = '';
            addEntryRow();
        } else {
            const error = await response.json();
            showMessage(`凭证创建失败: ${error.error}`, 'error');
        }
    } catch (error) {
        console.error('创建凭证失败:', error);
        showMessage('创建凭证失败，请稍后重试', 'error');
    }
}

// 添加分录行
function addEntryRow() {
    const container = document.getElementById('entriesContainer');
    const newRow = document.createElement('div');
    newRow.className = 'entry-row';
    newRow.innerHTML = `
        <input type="text" class="account-code" placeholder="科目代码" required>
        <input type="text" class="account-name" placeholder="科目名称" required>
        <input type="number" class="debit-amount" placeholder="借方金额" step="0.01">
        <input type="number" class="credit-amount" placeholder="贷方金额" step="0.01">
        <button type="button" class="btn btn-danger remove-entry">删除</button>
    `;
    container.appendChild(newRow);
}

// 删除分录行
function removeEntryRow(button) {
    const row = button.parentElement;
    // 确保至少保留一行
    if (document.querySelectorAll('.entry-row').length > 1) {
        row.remove();
    } else {
        showMessage('至少需要保留一个分录行', 'warning');
    }
}

// 创建科目
async function createAccount() {
    const accountData = {
        code: document.getElementById('accountCode').value,
        name: document.getElementById('accountName').value,
        account_type: document.getElementById('accountType').value,
        balance_direction: document.getElementById('balanceDirection').value,
        parent_code: document.getElementById('parentAccount').value,
        is_leaf: document.getElementById('isLeaf').checked
    };
    
    // 验证必填字段
    if (!accountData.code || !accountData.name || !accountData.account_type || !accountData.balance_direction) {
        showMessage('请填写所有必填字段', 'error');
        return;
    }
    
    try {
        const response = await fetch('/api/finance/accounts', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(accountData)
        });
        
        if (response.ok) {
            showMessage('科目创建成功', 'success');
            document.getElementById('accountForm').reset();
        } else {
            const error = await response.json();
            showMessage(`科目创建失败: ${error.error}`, 'error');
        }
    } catch (error) {
        console.error('创建科目失败:', error);
        showMessage('创建科目失败，请稍后重试', 'error');
    }
}

// 创建商品分类
async function createProductCategory() {
    const categoryData = {
        code: document.getElementById('categoryCode').value,
        name: document.getElementById('categoryName').value,
        parent_id: document.getElementById('categoryParent').value || null,
        description: document.getElementById('categoryDescription').value
    };
    
    // 验证必填字段
    if (!categoryData.name) {
        showMessage('分类名称不能为空', 'error');
        return;
    }
    
    try {
        const response = await fetch('/api/categories', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(categoryData)
        });
        
        if (response.ok) {
            showMessage('商品分类创建成功', 'success');
            document.getElementById('categoryForm').reset();
        } else {
            const error = await response.json();
            showMessage(`商品分类创建失败: ${error.error}`, 'error');
        }
    } catch (error) {
        console.error('创建商品分类失败:', error);
        showMessage('创建商品分类失败，请稍后重试', 'error');
    }
}

// 创建商品
async function createProduct() {
    const productData = {
        code: document.getElementById('productCode').value,
        name: document.getElementById('productName').value,
        category_id: document.getElementById('productCategory').value || null,
        unit: document.getElementById('productUnit').value,
        specification: document.getElementById('productSpec').value,
        description: document.getElementById('productDescription').value
    };
    
    // 验证必填字段
    if (!productData.name || !productData.unit) {
        showMessage('商品名称和单位不能为空', 'error');
        return;
    }
    
    try {
        const response = await fetch('/api/products', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(productData)
        });
        
        if (response.ok) {
            showMessage('商品创建成功', 'success');
            document.getElementById('productForm').reset();
        } else {
            const error = await response.json();
            showMessage(`商品创建失败: ${error.error}`, 'error');
        }
    } catch (error) {
        console.error('创建商品失败:', error);
        showMessage('创建商品失败，请稍后重试', 'error');
    }
}

// 创建供应商
async function createSupplier() {
    const supplierData = {
        code: document.getElementById('supplierCode').value,
        name: document.getElementById('supplierName').value,
        contact_person: document.getElementById('supplierContact').value,
        phone: document.getElementById('supplierPhone').value,
        email: document.getElementById('supplierEmail').value,
        address: document.getElementById('supplierAddress').value,
        tax_number: document.getElementById('supplierTax').value,
        bank_name: document.getElementById('supplierBank').value,
        bank_account: document.getElementById('supplierBankAccount').value
    };
    
    // 验证必填字段
    if (!supplierData.name) {
        showMessage('供应商名称不能为空', 'error');
        return;
    }
    
    try {
        const response = await fetch('/api/suppliers', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(supplierData)
        });
        
        if (response.ok) {
            showMessage('供应商创建成功', 'success');
            document.getElementById('supplierForm').reset();
        } else {
            const error = await response.json();
            showMessage(`供应商创建失败: ${error.error}`, 'error');
        }
    } catch (error) {
        console.error('创建供应商失败:', error);
        showMessage('创建供应商失败，请稍后重试', 'error');
    }
}

// 创建客户
async function createCustomer() {
    const customerData = {
        code: document.getElementById('customerCode').value,
        name: document.getElementById('customerName').value,
        contact_person: document.getElementById('customerContact').value,
        phone: document.getElementById('customerPhone').value,
        email: document.getElementById('customerEmail').value,
        address: document.getElementById('customerAddress').value,
        tax_number: document.getElementById('customerTax').value,
        bank_name: document.getElementById('customerBank').value,
        bank_account: document.getElementById('customerBankAccount').value
    };
    
    // 验证必填字段
    if (!customerData.name) {
        showMessage('客户名称不能为空', 'error');
        return;
    }
    
    try {
        const response = await fetch('/api/customers', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(customerData)
        });
        
        if (response.ok) {
            showMessage('客户创建成功', 'success');
            document.getElementById('customerForm').reset();
        } else {
            const error = await response.json();
            showMessage(`客户创建失败: ${error.error}`, 'error');
        }
    } catch (error) {
        console.error('创建客户失败:', error);
        showMessage('创建客户失败，请稍后重试', 'error');
    }
}

// 显示消息
function showMessage(message, type) {
    // 创建消息元素
    const messageEl = document.createElement('div');
    messageEl.className = `message message-${type}`;
    messageEl.textContent = message;
    
    // 添加到页面
    document.body.appendChild(messageEl);
    
    // 3秒后自动移除
    setTimeout(() => {
        messageEl.remove();
    }, 3000);
}

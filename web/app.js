// API 基础 URL
const API_BASE = 'http://localhost:8080';

// DOM 元素
const accountsSection = document.getElementById('accounts-section');
const vouchersSection = document.getElementById('vouchers-section');
const purchaseOrdersSection = document.getElementById('purchase-orders-section');
const salesOrdersSection = document.getElementById('sales-orders-section');

const accountsBtn = document.getElementById('accounts-btn');
const vouchersBtn = document.getElementById('vouchers-btn');
const purchaseOrdersBtn = document.getElementById('purchase-orders-btn');
const salesOrdersBtn = document.getElementById('sales-orders-btn');

// 导航按钮事件监听器
accountsBtn.addEventListener('click', () => showSection(accountsSection));
vouchersBtn.addEventListener('click', () => showSection(vouchersSection));
purchaseOrdersBtn.addEventListener('click', () => showSection(purchaseOrdersSection));
salesOrdersBtn.addEventListener('click', () => showSection(salesOrdersSection));

// 显示指定部分并隐藏其他部分
function showSection(section) {
    // 隐藏所有部分
    accountsSection.style.display = 'none';
    vouchersSection.style.display = 'none';
    purchaseOrdersSection.style.display = 'none';
    salesOrdersSection.style.display = 'none';
    
    // 显示指定部分
    section.style.display = 'block';
    
    // 更新活动按钮样式
    accountsBtn.classList.remove('active');
    vouchersBtn.classList.remove('active');
    purchaseOrdersBtn.classList.remove('active');
    salesOrdersBtn.classList.remove('active');
    
    if (section === accountsSection) accountsBtn.classList.add('active');
    if (section === vouchersSection) vouchersBtn.classList.add('active');
    if (section === purchaseOrdersSection) purchaseOrdersBtn.classList.add('active');
    if (section === salesOrdersSection) salesOrdersBtn.classList.add('active');
    
    // 加载数据
    if (section === accountsSection) loadAccounts();
    if (section === vouchersSection) loadVouchers();
    if (section === purchaseOrdersSection) loadPurchaseOrders();
    if (section === salesOrdersSection) loadSalesOrders();
}

// 初始化
showSection(accountsSection);

// 账户相关功能
document.getElementById('account-form').addEventListener('submit', createAccount);
document.getElementById('add-entry').addEventListener('click', addEntry);
document.getElementById('add-po-item').addEventListener('click', addPOItem);
document.getElementById('add-so-item').addEventListener('click', addSOItem);

// 采购订单表单提交
document.getElementById('purchase-order-form').addEventListener('submit', createPurchaseOrder);

// 销售订单表单提交
document.getElementById('sales-order-form').addEventListener('submit', createSalesOrder);

// 凭证表单提交
document.getElementById('voucher-form').addEventListener('submit', createVoucher);

// 创建账户
async function createAccount(e) {
    e.preventDefault();
    
    const accountData = {
        code: document.getElementById('account-code').value,
        name: document.getElementById('account-name').value,
        account_type: document.getElementById('account-type').value,
        balance_direction: 'DEBIT', // 简化处理
        is_leaf: true
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/finance/accounts`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(accountData)
        });
        
        if (response.ok) {
            alert('账户创建成功');
            document.getElementById('account-form').reset();
            loadAccounts();
        } else {
            const error = await response.json();
            alert('创建账户失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载账户列表
async function loadAccounts() {
    try {
        const response = await fetch(`${API_BASE}/api/finance/accounts`);
        if (response.ok) {
            const accounts = await response.json();
            displayAccounts(accounts);
        }
    } catch (error) {
        console.error('加载账户失败:', error);
    }
}

// 显示账户列表
function displayAccounts(accounts) {
    const accountsList = document.getElementById('accounts-list');
    accountsList.innerHTML = '<h3>账户列表</h3>';
    
    accounts.forEach(account => {
        const accountDiv = document.createElement('div');
        accountDiv.className = 'account-item';
        accountDiv.innerHTML = `
            <h3>${account.code} - ${account.name}</h3>
            <p>类型: ${account.account_type}</p>
        `;
        accountsList.appendChild(accountDiv);
    });
}

// 添加凭证分录
function addEntry() {
    const container = document.getElementById('entries-container');
    const entryDiv = document.createElement('div');
    entryDiv.className = 'entry';
    entryDiv.innerHTML = `
        <input type="text" class="entry-account" placeholder="科目代码" required>
        <input type="number" class="entry-debit" placeholder="借方金额" step="0.01">
        <input type="number" class="entry-credit" placeholder="贷方金额" step="0.01">
    `;
    container.appendChild(entryDiv);
}

// 创建凭证
async function createVoucher(e) {
    e.preventDefault();
    
    // 收集凭证数据
    const entries = [];
    const entryElements = document.querySelectorAll('.entry');
    
    for (let entryEl of entryElements) {
        const accountCode = entryEl.querySelector('.entry-account').value;
        const debit = parseFloat(entryEl.querySelector('.entry-debit').value) || 0;
        const credit = parseFloat(entryEl.querySelector('.entry-credit').value) || 0;
        
        entries.push({
            account_code: accountCode,
            debit_amount: debit,
            credit_amount: credit
        });
    }
    
    const voucherData = {
        voucher_date: document.getElementById('voucher-date').value,
        description: document.getElementById('voucher-description').value,
        entries: entries
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/finance/vouchers`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(voucherData)
        });
        
        if (response.ok) {
            alert('凭证创建成功');
            document.getElementById('voucher-form').reset();
            loadVouchers();
        } else {
            const error = await response.json();
            alert('创建凭证失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载凭证列表
async function loadVouchers() {
    try {
        const response = await fetch(`${API_BASE}/api/finance/vouchers`);
        if (response.ok) {
            const vouchers = await response.json();
            displayVouchers(vouchers);
        }
    } catch (error) {
        console.error('加载凭证失败:', error);
    }
}

// 显示凭证列表
function displayVouchers(vouchers) {
    const vouchersList = document.getElementById('vouchers-list');
    vouchersList.innerHTML = '<h3>凭证列表</h3>';
    
    vouchers.forEach(voucher => {
        const voucherDiv = document.createElement('div');
        voucherDiv.className = 'voucher-item';
        voucherDiv.innerHTML = `
            <h3>凭证 #${voucher.id}</h3>
            <p>日期: ${voucher.voucher_date}</p>
            <p>描述: ${voucher.description}</p>
        `;
        vouchersList.appendChild(voucherDiv);
    });
}

// 添加采购订单商品
function addPOItem() {
    const container = document.getElementById('po-items-container');
    const itemDiv = document.createElement('div');
    itemDiv.className = 'po-item';
    itemDiv.innerHTML = `
        <input type="text" class="item-code" placeholder="商品代码" required>
        <input type="text" class="item-name" placeholder="商品名称" required>
        <input type="number" class="item-quantity" placeholder="数量" required>
        <input type="number" class="item-price" placeholder="单价" step="0.01" required>
    `;
    container.appendChild(itemDiv);
}

// 创建采购订单
async function createPurchaseOrder(e) {
    e.preventDefault();
    
    // 收集订单商品数据
    const items = [];
    const itemElements = document.querySelectorAll('.po-item');
    
    for (let itemEl of itemElements) {
        const code = itemEl.querySelector('.item-code').value;
        const name = itemEl.querySelector('.item-name').value;
        const quantity = parseInt(itemEl.querySelector('.item-quantity').value);
        const price = parseFloat(itemEl.querySelector('.item-price').value);
        
        items.push({
            product_code: code,
            product_name: name,
            unit: '件', // 简化处理
            quantity: quantity,
            unit_price: price,
            amount: quantity * price
        });
    }
    
    const orderData = {
        supplier_code: document.getElementById('po-supplier-code').value,
        supplier_name: document.getElementById('po-supplier-name').value,
        order_date: document.getElementById('po-date').value,
        delivery_date: document.getElementById('po-delivery-date').value,
        items: items
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/purchase/orders`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(orderData)
        });
        
        if (response.ok) {
            alert('采购订单创建成功');
            document.getElementById('purchase-order-form').reset();
            loadPurchaseOrders();
        } else {
            const error = await response.json();
            alert('创建采购订单失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载采购订单列表
async function loadPurchaseOrders() {
    try {
        const response = await fetch(`${API_BASE}/api/purchase/orders`);
        if (response.ok) {
            const orders = await response.json();
            displayPurchaseOrders(orders);
        }
    } catch (error) {
        console.error('加载采购订单失败:', error);
    }
}

// 显示采购订单列表
function displayPurchaseOrders(orders) {
    const ordersList = document.getElementById('purchase-orders-list');
    ordersList.innerHTML = '<h3>采购订单列表</h3>';
    
    orders.forEach(order => {
        const orderDiv = document.createElement('div');
        orderDiv.className = 'order-item';
        orderDiv.innerHTML = `
            <h3>订单 #${order.id}</h3>
            <p>供应商: ${order.supplier_name}</p>
            <p>日期: ${order.order_date}</p>
            <p>状态: <span class="${order.status === 'APPROVED' ? 'status-approved' : order.status === 'CLOSED' ? 'status-closed' : ''}">${order.status}</span></p>
        `;
        ordersList.appendChild(orderDiv);
    });
}

// 添加销售订单商品
function addSOItem() {
    const container = document.getElementById('so-items-container');
    const itemDiv = document.createElement('div');
    itemDiv.className = 'so-item';
    itemDiv.innerHTML = `
        <input type="text" class="item-code" placeholder="商品代码" required>
        <input type="text" class="item-name" placeholder="商品名称" required>
        <input type="number" class="item-quantity" placeholder="数量" required>
        <input type="number" class="item-price" placeholder="单价" step="0.01" required>
    `;
    container.appendChild(itemDiv);
}

// 创建销售订单
async function createSalesOrder(e) {
    e.preventDefault();
    
    // 收集订单商品数据
    const items = [];
    const itemElements = document.querySelectorAll('.so-item');
    
    for (let itemEl of itemElements) {
        const code = itemEl.querySelector('.item-code').value;
        const name = itemEl.querySelector('.item-name').value;
        const quantity = parseInt(itemEl.querySelector('.item-quantity').value);
        const price = parseFloat(itemEl.querySelector('.item-price').value);
        
        items.push({
            product_code: code,
            product_name: name,
            unit: '件', // 简化处理
            quantity: quantity,
            unit_price: price,
            amount: quantity * price
        });
    }
    
    const orderData = {
        customer_code: document.getElementById('so-customer-code').value,
        customer_name: document.getElementById('so-customer-name').value,
        order_date: document.getElementById('so-date').value,
        delivery_date: document.getElementById('so-delivery-date').value,
        items: items
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/sales/orders`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(orderData)
        });
        
        if (response.ok) {
            alert('销售订单创建成功');
            document.getElementById('sales-order-form').reset();
            loadSalesOrders();
        } else {
            const error = await response.json();
            alert('创建销售订单失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载销售订单列表
async function loadSalesOrders() {
    try {
        const response = await fetch(`${API_BASE}/api/sales/orders`);
        if (response.ok) {
            const orders = await response.json();
            displaySalesOrders(orders);
        }
    } catch (error) {
        console.error('加载销售订单失败:', error);
    }
}

// 显示销售订单列表
function displaySalesOrders(orders) {
    const ordersList = document.getElementById('sales-orders-list');
    ordersList.innerHTML = '<h3>销售订单列表</h3>';
    
    orders.forEach(order => {
        const orderDiv = document.createElement('div');
        orderDiv.className = 'order-item';
        orderDiv.innerHTML = `
            <h3>订单 #${order.id}</h3>
            <p>客户: ${order.customer_name}</p>
            <p>日期: ${order.order_date}</p>
            <p>状态: <span class="${order.status === 'APPROVED' ? 'status-approved' : order.status === 'CLOSED' ? 'status-closed' : ''}">${order.status}</span></p>
        `;
        ordersList.appendChild(orderDiv);
    });
}
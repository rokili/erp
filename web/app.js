// API 基础 URL
const API_BASE = 'http://localhost:8080';

// DOM 元素
const dashboardSection = document.getElementById('dashboard-section');
const accountsSection = document.getElementById('accounts-section');
const vouchersSection = document.getElementById('vouchers-section');
const productsSection = document.getElementById('products-section');
const categoriesSection = document.getElementById('categories-section');
const suppliersSection = document.getElementById('suppliers-section');
const customersSection = document.getElementById('customers-section');
const inventorySection = document.getElementById('inventory-section');
const purchaseOrdersSection = document.getElementById('purchase-orders-section');
const salesOrdersSection = document.getElementById('sales-orders-section');

const dashboardBtn = document.getElementById('dashboard-btn');
const accountsBtn = document.getElementById('accounts-btn');
const vouchersBtn = document.getElementById('vouchers-btn');
const productsBtn = document.getElementById('products-btn');
const categoriesBtn = document.getElementById('categories-btn');
const suppliersBtn = document.getElementById('suppliers-btn');
const customersBtn = document.getElementById('customers-btn');
const inventoryBtn = document.getElementById('inventory-btn');
const purchaseOrdersBtn = document.getElementById('purchase-orders-btn');
const salesOrdersBtn = document.getElementById('sales-orders-btn');

// 导航按钮事件监听器
dashboardBtn.addEventListener('click', () => showSection(dashboardSection));
accountsBtn.addEventListener('click', () => showSection(accountsSection));
vouchersBtn.addEventListener('click', () => showSection(vouchersSection));
productsBtn.addEventListener('click', () => showSection(productsSection));
categoriesBtn.addEventListener('click', () => showSection(categoriesSection));
suppliersBtn.addEventListener('click', () => showSection(suppliersSection));
customersBtn.addEventListener('click', () => showSection(customersSection));
inventoryBtn.addEventListener('click', () => showSection(inventorySection));
purchaseOrdersBtn.addEventListener('click', () => showSection(purchaseOrdersSection));
salesOrdersBtn.addEventListener('click', () => showSection(salesOrdersSection));

// 显示指定部分并隐藏其他部分
function showSection(section) {
    // 隐藏所有部分
    dashboardSection.style.display = 'none';
    accountsSection.style.display = 'none';
    vouchersSection.style.display = 'none';
    productsSection.style.display = 'none';
    categoriesSection.style.display = 'none';
    suppliersSection.style.display = 'none';
    customersSection.style.display = 'none';
    inventorySection.style.display = 'none';
    purchaseOrdersSection.style.display = 'none';
    salesOrdersSection.style.display = 'none';
    
    // 显示指定部分
    section.style.display = 'block';
    
    // 更新活动按钮样式
    dashboardBtn.classList.remove('active');
    accountsBtn.classList.remove('active');
    vouchersBtn.classList.remove('active');
    productsBtn.classList.remove('active');
    categoriesBtn.classList.remove('active');
    suppliersBtn.classList.remove('active');
    customersBtn.classList.remove('active');
    inventoryBtn.classList.remove('active');
    purchaseOrdersBtn.classList.remove('active');
    salesOrdersBtn.classList.remove('active');
    
    if (section === dashboardSection) dashboardBtn.classList.add('active');
    if (section === accountsSection) accountsBtn.classList.add('active');
    if (section === vouchersSection) vouchersBtn.classList.add('active');
    if (section === productsSection) productsBtn.classList.add('active');
    if (section === categoriesSection) categoriesBtn.classList.add('active');
    if (section === suppliersSection) suppliersBtn.classList.add('active');
    if (section === customersSection) customersBtn.classList.add('active');
    if (section === inventorySection) inventoryBtn.classList.add('active');
    if (section === purchaseOrdersSection) purchaseOrdersBtn.classList.add('active');
    if (section === salesOrdersSection) salesOrdersBtn.classList.add('active');
    
    // 加载数据
    if (section === dashboardSection) loadDashboard();
    if (section === accountsSection) loadAccounts();
    if (section === vouchersSection) loadVouchers();
    if (section === productsSection) loadProducts();
    if (section === categoriesSection) loadCategories();
    if (section === suppliersSection) loadSuppliers();
    if (section === customersSection) loadCustomers();
    if (section === inventorySection) loadInventory();
    if (section === purchaseOrdersSection) {
        loadSuppliersForPO();
        loadProductsForPO();
        loadPurchaseOrders();
    }
    if (section === salesOrdersSection) {
        loadCustomersForSO();
        loadProductsForSO();
        loadSalesOrders();
    }
}

// 初始化
showSection(dashboardSection);

// 仪表板功能
async function loadDashboard() {
    try {
        // 加载统计数据
        const productsResponse = await fetch(`${API_BASE}/api/products`);
        const suppliersResponse = await fetch(`${API_BASE}/api/suppliers`);
        const customersResponse = await fetch(`${API_BASE}/api/customers`);
        
        if (productsResponse.ok && suppliersResponse.ok && customersResponse.ok) {
            const products = await productsResponse.json();
            const suppliers = await suppliersResponse.json();
            const customers = await customersResponse.json();
            
            document.getElementById('total-products').textContent = products.length;
            document.getElementById('total-suppliers').textContent = suppliers.length;
            document.getElementById('total-customers').textContent = customers.length;
        }
    } catch (error) {
        console.error('加载仪表板数据失败:', error);
    }
}

// 账户相关功能
document.getElementById('account-form').addEventListener('submit', createAccount);
document.getElementById('add-entry').addEventListener('click', addEntry);

// 商品相关功能
document.getElementById('product-form').addEventListener('submit', createProduct);

// 商品分类相关功能
document.getElementById('category-form').addEventListener('submit', createCategory);

// 供应商相关功能
document.getElementById('supplier-form').addEventListener('submit', createSupplier);

// 客户相关功能
document.getElementById('customer-form').addEventListener('submit', createCustomer);

// 采购订单相关功能
document.getElementById('add-po-item').addEventListener('click', addPOItem);
document.getElementById('purchase-order-form').addEventListener('submit', createPurchaseOrder);

// 销售订单相关功能
document.getElementById('add-so-item').addEventListener('click', addSOItem);
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

// 创建商品
async function createProduct(e) {
    e.preventDefault();
    
    const productData = {
        code: document.getElementById('product-code').value,
        name: document.getElementById('product-name').value,
        category_id: parseInt(document.getElementById('product-category').value) || null,
        unit: document.getElementById('product-unit').value,
        specification: document.getElementById('product-spec').value
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/products`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(productData)
        });
        
        if (response.ok) {
            alert('商品创建成功');
            document.getElementById('product-form').reset();
            loadProducts();
        } else {
            const error = await response.json();
            alert('创建商品失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载商品列表
async function loadProducts() {
    try {
        const response = await fetch(`${API_BASE}/api/products`);
        if (response.ok) {
            const products = await response.json();
            displayProducts(products);
            updateProductSelectors(products);
        }
    } catch (error) {
        console.error('加载商品失败:', error);
    }
}

// 显示商品列表
function displayProducts(products) {
    const productsList = document.getElementById('products-list');
    productsList.innerHTML = '<h3>商品列表</h3>';
    
    products.forEach(product => {
        const productDiv = document.createElement('div');
        productDiv.className = 'product-item';
        productDiv.innerHTML = `
            <h3>${product.code} - ${product.name}</h3>
            <p>单位: ${product.unit}</p>
            <p>规格: ${product.specification || '无'}</p>
        `;
        productsList.appendChild(productDiv);
    });
}

// 更新商品选择器
function updateProductSelectors(products) {
    const productSelectors = document.querySelectorAll('.item-product');
    productSelectors.forEach(selector => {
        selector.innerHTML = '<option value="">选择商品</option>';
        products.forEach(product => {
            const option = document.createElement('option');
            option.value = product.id;
            option.textContent = `${product.code} - ${product.name}`;
            selector.appendChild(option);
        });
    });
}

// 创建商品分类
async function createCategory(e) {
    e.preventDefault();
    
    const categoryData = {
        code: document.getElementById('category-code').value,
        name: document.getElementById('category-name').value,
        parent_id: parseInt(document.getElementById('category-parent').value) || null,
        description: document.getElementById('category-desc').value
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/categories`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(categoryData)
        });
        
        if (response.ok) {
            alert('分类创建成功');
            document.getElementById('category-form').reset();
            loadCategories();
        } else {
            const error = await response.json();
            alert('创建分类失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载商品分类列表
async function loadCategories() {
    try {
        const response = await fetch(`${API_BASE}/api/categories`);
        if (response.ok) {
            const categories = await response.json();
            displayCategories(categories);
            updateCategorySelectors(categories);
        }
    } catch (error) {
        console.error('加载分类失败:', error);
    }
}

// 显示商品分类列表
function displayCategories(categories) {
    const categoriesList = document.getElementById('categories-list');
    categoriesList.innerHTML = '<h3>商品分类列表</h3>';
    
    categories.forEach(category => {
        const categoryDiv = document.createElement('div');
        categoryDiv.className = 'category-item';
        categoryDiv.innerHTML = `
            <h3>${category.code} - ${category.name}</h3>
            <p>描述: ${category.description || '无'}</p>
        `;
        categoriesList.appendChild(categoryDiv);
    });
}

// 更新分类选择器
function updateCategorySelectors(categories) {
    const categorySelect = document.getElementById('product-category');
    const parentSelect = document.getElementById('category-parent');
    
    categorySelect.innerHTML = '<option value="">选择分类</option>';
    parentSelect.innerHTML = '<option value="">无上级分类</option>';
    
    categories.forEach(category => {
        const option1 = document.createElement('option');
        option1.value = category.id;
        option1.textContent = `${category.code} - ${category.name}`;
        categorySelect.appendChild(option1);
        
        const option2 = document.createElement('option');
        option2.value = category.id;
        option2.textContent = `${category.code} - ${category.name}`;
        parentSelect.appendChild(option2);
    });
}

// 创建供应商
async function createSupplier(e) {
    e.preventDefault();
    
    const supplierData = {
        code: document.getElementById('supplier-code').value,
        name: document.getElementById('supplier-name').value,
        contact_person: document.getElementById('supplier-contact').value,
        phone: document.getElementById('supplier-phone').value,
        email: document.getElementById('supplier-email').value,
        address: document.getElementById('supplier-address').value
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/suppliers`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(supplierData)
        });
        
        if (response.ok) {
            alert('供应商创建成功');
            document.getElementById('supplier-form').reset();
            loadSuppliers();
        } else {
            const error = await response.json();
            alert('创建供应商失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载供应商列表
async function loadSuppliers() {
    try {
        const response = await fetch(`${API_BASE}/api/suppliers`);
        if (response.ok) {
            const suppliers = await response.json();
            displaySuppliers(suppliers);
            updateSupplierSelectors(suppliers);
        }
    } catch (error) {
        console.error('加载供应商失败:', error);
    }
}

// 显示供应商列表
function displaySuppliers(suppliers) {
    const suppliersList = document.getElementById('suppliers-list');
    suppliersList.innerHTML = '<h3>供应商列表</h3>';
    
    suppliers.forEach(supplier => {
        const supplierDiv = document.createElement('div');
        supplierDiv.className = 'supplier-item';
        supplierDiv.innerHTML = `
            <h3>${supplier.code} - ${supplier.name}</h3>
            <p>联系人: ${supplier.contact_person || '无'}</p>
            <p>电话: ${supplier.phone || '无'}</p>
        `;
        suppliersList.appendChild(supplierDiv);
    });
}

// 更新供应商选择器
function updateSupplierSelectors(suppliers) {
    const supplierSelect = document.getElementById('po-supplier');
    supplierSelect.innerHTML = '<option value="">选择供应商</option>';
    
    suppliers.forEach(supplier => {
        const option = document.createElement('option');
        option.value = supplier.id;
        option.textContent = `${supplier.code} - ${supplier.name}`;
        supplierSelect.appendChild(option);
    });
}

// 加载供应商（用于采购订单）
async function loadSuppliersForPO() {
    try {
        const response = await fetch(`${API_BASE}/api/suppliers`);
        if (response.ok) {
            const suppliers = await response.json();
            updateSupplierSelectors(suppliers);
        }
    } catch (error) {
        console.error('加载供应商失败:', error);
    }
}

// 创建客户
async function createCustomer(e) {
    e.preventDefault();
    
    const customerData = {
        code: document.getElementById('customer-code').value,
        name: document.getElementById('customer-name').value,
        contact_person: document.getElementById('customer-contact').value,
        phone: document.getElementById('customer-phone').value,
        email: document.getElementById('customer-email').value,
        address: document.getElementById('customer-address').value
    };
    
    try {
        const response = await fetch(`${API_BASE}/api/customers`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(customerData)
        });
        
        if (response.ok) {
            alert('客户创建成功');
            document.getElementById('customer-form').reset();
            loadCustomers();
        } else {
            const error = await response.json();
            alert('创建客户失败: ' + error.error);
        }
    } catch (error) {
        alert('网络错误: ' + error.message);
    }
}

// 加载客户列表
async function loadCustomers() {
    try {
        const response = await fetch(`${API_BASE}/api/customers`);
        if (response.ok) {
            const customers = await response.json();
            displayCustomers(customers);
            updateCustomerSelectors(customers);
        }
    } catch (error) {
        console.error('加载客户失败:', error);
    }
}

// 显示客户列表
function displayCustomers(customers) {
    const customersList = document.getElementById('customers-list');
    customersList.innerHTML = '<h3>客户列表</h3>';
    
    customers.forEach(customer => {
        const customerDiv = document.createElement('div');
        customerDiv.className = 'customer-item';
        customerDiv.innerHTML = `
            <h3>${customer.code} - ${customer.name}</h3>
            <p>联系人: ${customer.contact_person || '无'}</p>
            <p>电话: ${customer.phone || '无'}</p>
        `;
        customersList.appendChild(customerDiv);
    });
}

// 更新客户选择器
function updateCustomerSelectors(customers) {
    const customerSelect = document.getElementById('so-customer');
    customerSelect.innerHTML = '<option value="">选择客户</option>';
    
    customers.forEach(customer => {
        const option = document.createElement('option');
        option.value = customer.id;
        option.textContent = `${customer.code} - ${customer.name}`;
        customerSelect.appendChild(option);
    });
}

// 加载客户（用于销售订单）
async function loadCustomersForSO() {
    try {
        const response = await fetch(`${API_BASE}/api/customers`);
        if (response.ok) {
            const customers = await response.json();
            updateCustomerSelectors(customers);
        }
    } catch (error) {
        console.error('加载客户失败:', error);
    }
}

// 加载商品（用于采购订单）
async function loadProductsForPO() {
    try {
        const response = await fetch(`${API_BASE}/api/products`);
        if (response.ok) {
            const products = await response.json();
            updateProductSelectors(products);
        }
    } catch (error) {
        console.error('加载商品失败:', error);
    }
}

// 加载商品（用于销售订单）
async function loadProductsForSO() {
    try {
        const response = await fetch(`${API_BASE}/api/products`);
        if (response.ok) {
            const products = await response.json();
            updateProductSelectors(products);
        }
    } catch (error) {
        console.error('加载商品失败:', error);
    }
}

// 添加采购订单商品
function addPOItem() {
    const container = document.getElementById('po-items-container');
    const itemDiv = document.createElement('div');
    itemDiv.className = 'po-item';
    itemDiv.innerHTML = `
        <select class="item-product" required>
            <option value="">选择商品</option>
        </select>
        <input type="number" class="item-quantity" placeholder="数量" required>
        <input type="number" class="item-price" placeholder="单价" step="0.01" required>
    `;
    container.appendChild(itemDiv);
    
    // 更新商品选择器
    loadProductsForPO();
}

// 创建采购订单
async function createPurchaseOrder(e) {
    e.preventDefault();
    
    // 收集订单商品数据
    const items = [];
    const itemElements = document.querySelectorAll('.po-item');
    
    for (let itemEl of itemElements) {
        const productId = parseInt(itemEl.querySelector('.item-product').value);
        const quantity = parseFloat(itemEl.querySelector('.item-quantity').value);
        const price = parseFloat(itemEl.querySelector('.item-price').value);
        
        // 获取商品信息
        const productResponse = await fetch(`${API_BASE}/api/products/${productId}`);
        if (!productResponse.ok) continue;
        
        const product = await productResponse.json();
        
        items.push({
            product_code: product.code,
            product_name: product.name,
            unit: product.unit,
            quantity: quantity,
            unit_price: price,
            amount: quantity * price
        });
    }
    
    // 获取供应商信息
    const supplierId = document.getElementById('po-supplier').value;
    const supplierResponse = await fetch(`${API_BASE}/api/suppliers/${supplierId}`);
    if (!supplierResponse.ok) {
        alert('获取供应商信息失败');
        return;
    }
    
    const supplier = await supplierResponse.json();
    
    const orderData = {
        supplier_code: supplier.code,
        supplier_name: supplier.name,
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
        <select class="item-product" required>
            <option value="">选择商品</option>
        </select>
        <input type="number" class="item-quantity" placeholder="数量" required>
        <input type="number" class="item-price" placeholder="单价" step="0.01" required>
    `;
    container.appendChild(itemDiv);
    
    // 更新商品选择器
    loadProductsForSO();
}

// 创建销售订单
async function createSalesOrder(e) {
    e.preventDefault();
    
    // 收集订单商品数据
    const items = [];
    const itemElements = document.querySelectorAll('.so-item');
    
    for (let itemEl of itemElements) {
        const productId = parseInt(itemEl.querySelector('.item-product').value);
        const quantity = parseFloat(itemEl.querySelector('.item-quantity').value);
        const price = parseFloat(itemEl.querySelector('.item-price').value);
        
        // 获取商品信息
        const productResponse = await fetch(`${API_BASE}/api/products/${productId}`);
        if (!productResponse.ok) continue;
        
        const product = await productResponse.json();
        
        items.push({
            product_code: product.code,
            product_name: product.name,
            unit: product.unit,
            quantity: quantity,
            unit_price: price,
            amount: quantity * price
        });
    }
    
    // 获取客户信息
    const customerId = document.getElementById('so-customer').value;
    const customerResponse = await fetch(`${API_BASE}/api/customers/${customerId}`);
    if (!customerResponse.ok) {
        alert('获取客户信息失败');
        return;
    }
    
    const customer = await customerResponse.json();
    
    const orderData = {
        customer_code: customer.code,
        customer_name: customer.name,
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

// 加载库存信息
async function loadInventory() {
    try {
        const response = await fetch(`${API_BASE}/api/inventory`);
        if (response.ok) {
            const inventory = await response.json();
            displayInventory(inventory);
        }
    } catch (error) {
        console.error('加载库存信息失败:', error);
    }
}

// 显示库存信息
function displayInventory(inventory) {
    const inventoryList = document.getElementById('inventory-list');
    inventoryList.innerHTML = '<h3>库存列表</h3>';
    
    inventory.forEach(item => {
        const itemDiv = document.createElement('div');
        itemDiv.className = 'inventory-item';
        itemDiv.innerHTML = `
            <h3>${item.product_code} - ${item.product_name}</h3>
            <p>库存数量: ${item.quantity} ${item.unit}</p>
            <p>单位成本: ¥${item.unit_cost.toFixed(2)}</p>
            <p>总成本: ¥${item.total_cost.toFixed(2)}</p>
        `;
        inventoryList.appendChild(itemDiv);
    });
}
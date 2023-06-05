the app is: GoMagentoAPIs

the files we have decided to generate are: 
1. go_magento_apis.go
2. go_magento_apis_test.go
3. java_deskera_sync.java
4. java_deskera_sync_test.java
5. README.md

Shared dependencies:
1. Auth token variable: authToken
2. API base URLs: magentoBaseURL, deskeraBaseURL
3. Data schemas: Product, Customer, Order, Payment, DeskeraProduct, DeskeraContact, DeskeraSalesInvoice
4. Filter variables: createdDate, updatedDate, productName, productCode, customerName, customerCode, orderId
5. Function names: getAllProducts, getAllCustomers, getAllOrders, getAllPayments, getProductById, getCustomerById, getOrderById, syncProducts, syncCustomers, syncOrders
6. Message names: productsResponse, customersResponse, ordersResponse, paymentsResponse, productDetailsResponse, customerDetailsResponse, orderDetailsResponse
7. Cron job names: syncProductsJob, syncCustomersJob, syncOrdersJob
8. ID names for DOM elements (if applicable): N/A
# GoMagentoAPIs

This project provides APIs to interact with Magento community latest edition in Golang. No backend needed.

## Installation

Clone the repository:

```
git clone https://github.com/juniorvish/GoMagentoAPIs.git
```

## Usage

1. Set up your Magento and Deskera API credentials in the `auth/auth.go` file.
2. Run the Golang API server:

```
go run GoMagentoAPIs/main.go
```

3. Run the Java cron jobs for syncing data with Deskera:

```
javac GoMagentoAPIs/cronjobs/JavaSync/main.java
java GoMagentoAPIs/cronjobs/JavaSync/main
```

## API Endpoints

- `GET /products`: Get all products with filters and pagination. Filters: created date, updated date, product name, product code
- `GET /customers`: Get all customers with filters and pagination. Filters: created date, updated date, customer name, customer code
- `GET /orders`: Get all orders with filters and pagination. Filters: created date, updated date, order id, customer code, customer name
- `GET /payments`: Get all payments against orders with filters and pagination
- `GET /products/:id`: Get details on a product by id
- `GET /customers/:id`: Get details on a customer by id
- `GET /orders/:id`: Get details on an order by id

## Java Cron Jobs

- `syncProductsJob`: Sync every 30 minutes the newly created products from Magento to Deskera Products
- `syncCustomersJob`: Sync every 30 minutes the newly created customers from Magento to Deskera Contacts
- `syncOrdersJob`: Sync every 30 minutes the newly created order from Magento to Deskera Sales Invoice

## Dependencies

- Golang for the API server
- Java for the cron jobs
- Magento community latest edition
- Deskera API

## License

MIT License
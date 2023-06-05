# GoMagentoAPIs

This project provides APIs to interact with Magento community latest edition using Golang. No backend needed.

## Installation

Clone the repository:

```
git clone https://github.com/juniorvish/GoMagentoAPIs.git
```

Navigate to the project directory:

```
cd GoMagentoAPIs
```

## Usage

1. Set up your Magento store and obtain an auth token.
2. Replace the placeholder auth token in `auth/auth.go` with your actual auth token.
3. Run the project:

```
go run main.go
```

## API Endpoints

- `GET /api/products`: Get all products with filters and pagination
- `GET /api/customers`: Get all customers with filters and pagination
- `GET /api/orders`: Get all orders with filters and pagination
- `GET /api/payments`: Get all payments against orders with filters and pagination
- `GET /api/products/:id`: Get details on a product by id
- `GET /api/customers/:id`: Get details on a customer by id
- `GET /api/orders/:id`: Get details on an order by id

## Filters and Pagination

You can apply filters and pagination to the `GET` requests for products, customers, orders, and payments by adding query parameters to the request URL.

Example:

```
GET /api/products?page=1&limit=10&filter[name]=example
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
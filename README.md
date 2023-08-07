# btc-tx-manager

# Transaction Service

This service provides CRUD functionalities for Bitcoin blockchain transactions. It is written in Go and leverages the Go Fiber web framework and MongoDB for database operations.

## Features

- **Create**: Add a new transaction to the database by 'blockchain.com' API.
- **Read**: Retrieve one or multiple transactions based on filtering criteria.
- **Update**: Update specific details of a transaction.
- **Delete**: Remove a transaction.

## Endpoints

- **POST** `/transactions/:hashId`: Add a new transaction.
- **GET** `/transactions`: Retrieve a list of transactions. Supports filtering by date range, input address, and output address.
- **GET** `/transactions/:hashId`: Retrieve details of a specific transaction using its hash.
- **PUT** `/transactions/:hashId`: Update details of a specific transaction.
- **DELETE** `/transactions/:hashId`: Delete a transaction.

## Supported Filters

- **Date Range**: Filter transactions between two dates using the `start_date` and `end_date` parameters.
- **Input Address**: Filter transactions that have a specific address as an input using the `input_address` parameter.
- **Output Address**: Filter transactions that have a specific address as an output using the `output_address` parameter.

## How to Run

1. Install the dependencies using `go mod download`.
2. Set up the MongoDB database running `docker compose up`
3. Update the credentials in the envirolment file.
4. Run the service using `go run main.go`.

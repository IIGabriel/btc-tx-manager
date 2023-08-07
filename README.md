# btc-tx-manager

# Transaction Service

This service provides CRUD functionalities for Bitcoin blockchain transactions. It is written in Go and leverages the Go Fiber web framework and MongoDB for database operations.
The Blockchain.com API was used to fetch the data

## Features

- **Create**: Add a new transaction to the database using data fetched from the blockchain.
- **Read**: Retrieve one or multiple transactions based on filtering criteria.
- **Update**: Update specific details of a transaction. Fields that can be updated include:
    - `time (RFC3339)` 
    - `fee`
    - `inputs`
    - `outputs`
    - `confirmations`
- **Delete**: Remove a transaction.

## Endpoints

- **POST** `/transactions/:hash_or_id`: Add a new transaction using data fetched from the blockchain.
- **GET** `/transactions`: Retrieve a list of transactions. Supports filtering by date range, input address, and output address.
  - Filtering by date range, input address, and output address.
  - Sorting by passing a `sort_field` and optionally an `asc` parameter (set to `true` for ascending order).
  - Pagination using the `page` and `perPage` parameters.
- **GET** `/transactions/:hash_or_id`: Retrieve details of a specific transaction using its hash.
- **PUT** `/transactions/:hash_or_id`: Update details of a specific transaction.
- **PUT** `/transactions/blockchain/:hash`: Update details of a specific transaction using data fetched from the blockchain.
- **DELETE** `/transactions/:hash_or_id`: Delete a transaction.

## Supported Filters

- **Date Range**: Filter transactions between two dates using the `start_date` and `end_date` parameters.
- **Input Address**: Filter transactions that have a specific address as an input using the `input_address` parameter.
- **Output Address**: Filter transactions that have a specific address as an output using the `output_address` parameter.

## How to Run

1. Install the dependencies using `go mod download`.
2. Set up the MongoDB database running `docker compose up`
3. Update the credentials in the envirolment file.
4. Run the service using `go run main.go`.

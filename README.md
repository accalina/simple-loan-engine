# simple-loan-engine

This project is a Loan Engine API built using Go and the Fiber framework. The API allows creating, approving, investing in, and disbursing loans. The application uses PostgreSQL as the database.

## Prerequisites

- Go 1.16+
- PostgreSQL
- Git

## Installation

### Clone the Repository

```sh
$ git clone https://github.com/accalina/simple-loan-engine
$ cd simple-loan-engine
```

### Set Up Environment Variables
Create a `.env` file in the root directory of the project and fill it with the following values:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=loan
```

### Set Up the Database
1. Install PostgreSQL if it's not already installed. Follow the instructions on the [PostgreSQL website](https://www.postgresql.org/download/).
2. Create a new PostgreSQL database:
```sh
$ psql -U postgres
```

```sql
CREATE DATABASE loan;
CREATE USER your_db_user WITH ENCRYPTED PASSWORD 'your_db_password';
GRANT ALL PRIVILEGES ON DATABASE loan TO your_db_user;
```

### Run the Application
1. Install Go dependencies:
```sh
$ go mod tidy
```

2. Run the application:
```sh
$ go run main.go
```

The application should now be running on `http://localhost:8081`.
___
## Usage

### Simple flow for Loan
1. Create new Investor
```sh
curl -X POST http://localhost:8081/investor -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "email": "john.doe@example.com"
}'
```

2. Create New Loan
```sh
curl -X POST http://localhost:8081/loan -H "Content-Type: application/json" -d '{
  "borrower_id": "123456",
  "principal_amount": 10000,
  "rate": 5,
  "roi": 10,
  "agreement_letter_link": "http://example.com/agreement.pdf"
}'
```

2. Approve a Loan
```sh
curl -X PUT http://localhost:8081/loan/approve/{loan_id} -H "Content-Type: application/json" -d '{
  "proof_photo_url": "http://example.com/proof.jpg",
  "validator_id": "987654321",
  "approval_date": "2024-07-26T00:00:00Z"
}'
```

3. Invest in a loan
```sh
curl -X PUT http://localhost:8081/loan/invest/{loan_id} -H "Content-Type: application/json" -d '{
  "investor_id": "ab7e34f1-c4e8-4f8c-812d-b72c4e4c8ab3",
  "amount": 5000
}'
```

4. Disburse a loan
```sh
curl -X PUT http://localhost:8081/loan/disburse/{loan_id} -H "Content-Type: application/json" -d '{
  "signed_agreement_url": "http://example.com/signed_agreement.pdf",
  "officer_id": "123456789",
  "disbursement_date": "2024-07-26T00:00:00Z"
}'
```

### Retrieving Data
- Get Loan by ID
```sh
curl -X GET http://localhost:8081/loan/{loan_id}
```

- Get Investor by ID
```sh
curl -X GET http://localhost:8081/investor/{investor_id}
```
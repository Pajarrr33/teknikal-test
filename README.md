
# Teknikal Test
This repository contains a simple API developed in Golang to facilitate interactions between merchants and banks.




## Tech Stack

- Golang : https://github.com/golang/go
- PostgreSQL (Database) : https://www.postgresql.org/


## Framework & Library
- Gin Gonic (HTTP Framework) : https://github.com/gin-gonic/gin
- Godotenv (Configuration) : https://github.com/joho/godotenv
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Golang JWT : https://github.com/golang-jwt/jwt
## Installation

1. Clone this repository
```bash
    https://github.com/Pajarrr33/teknikal-test.git
```
2. Create a database
```bash
    CREATE DATABASE example_name
```

3. Run this DDL query or copy it from DDL.sql file
```bash
    CREATE TABLE customers (
        id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier for the customer
        name VARCHAR(255) NOT NULL, -- Name of the customer
        email VARCHAR(255) NOT NULL UNIQUE, -- Email of the customer, must be unique
        password TEXT NOT NULL, -- Hashed password of the customer
        balance NUMERIC(15, 2) DEFAULT 0.00, -- Balance with two decimal points
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- Record creation time with timezone
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW() -- Record update time with timezone
    );

    CREATE TABLE merchant (
        id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier for the merchant
        name VARCHAR(255) NOT NULL, -- Name of the merchant
        category VARCHAR(255) NOT NULL, -- Category of the merchant
        contact VARCHAR(255) NOT NULL, -- Contact information of the merchant
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

    CREATE TABLE transaction (
        id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier for the transaction
        customer_id INT NOT NULL REFERENCES customers(id), -- ID of the customer involved in the transaction
        merchant_id INT NOT NULL REFERENCES merchant(id), -- ID of the merchant involved in the transaction
        amount NUMERIC(15, 2) NOT NULL, -- Amount of the transaction with two decimal points
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- Record creation time with timezone
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW() -- Record update time with timezone
    );

    CREATE TABLE expired_token (
        id SERIAL PRIMARY KEY, -- Auto-incrementing unique identifier for the expired token
        token TEXT NOT NULL -- Token value
    );

```
4. Run this DML query or copy it from DML.sql file
```bash
    INSERT INTO customers (name, email, password, balance)
    VALUES 
        -- All of the password was 12345678
        ('fajar', 'fajarshidik@gmail.com', '$2a$10$isRQ8HTQ8hRl.UWZ/KEXr.b.gGOuWNYiLDBCQH.2NciPnOxmYGSHu',0.00),
        ('ucup', 'ucup@example.com', '$2a$10$isRQ8HTQ8hRl.UWZ/KEXr.b.gGOuWNYiLDBCQH.2NciPnOxmYGSHu', 1000.00);

    INSERT INTO merchant (name, category, contact)
    VALUES
        ('Seonjana Bakery', 'Bakery', '08123456789'),
        ('Mie Abang Adek', 'Fast Food', '08123456790');

    INSERT INTO transaction (customer_id, merchant_id, amount)
    VALUES
        (1, 1, 100.00),
        (2, 2, 200.00);

    INSERT INTO expired_token (token)
    VALUES
        ('eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJpYXQiOjE2NjMzODA0NjAsImV4cCI6MTY2MzQxNjA2MH0.DYH1G2E0oV3Q2VxV4M2oIzQ5OeQ3wZoZC1rGc9Oa8C8');
    
```

5. Configure your database in env_example file and change the file name to .env_example
```bash
    DB_HOST= your_database_host
    DB_PORT= your_database_port
    DB_USER= your_database_username
    DB_PASSWORD= your_database_password
    DB_NAME= your_database_name
    DB_DRIVER= your_database_driver
    API_PORT= your_api_port
    TOKEN_ISSUE= application_name / your_name
    TOKEN_SECRET= your_secret_token
    TOKEN_EXPIRE= token_expire_times_in_minutes
```

6. Navigate to the project directory
```bash
    cd teknikal-test
```

7. Install necessary dependencies
```bash
    go mod tidy
```

8. Run the application
```bash
    go run main.go
```



    
## API Spec
You can import a postman colletion in `Teknikal Test.postman_collection.json` files

### Login

Request :

- Method : POST
- Endpoint : `api/v1/login`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

```json
    {
        "email" : "testingapp@gmail.com",
        "password" : "12345678"
    }
```

Response :

- Status : 200 OK
- Body :
```json
    {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtbmNfdGVrbmlrYWxfdGVzdCIsImV4cCI6MTczMjUzNzYxNSwiaWF0IjoxNzMyNTM1ODE1LCJpZCI6IjMiLCJ1c2VySWQiOiJmYWphcnNoaWRpa0BnbWFpbC5jb20ifQ.dxr14vdv4YOrG9I60TRa-yV44qzCt98FmxCV8gWyH5o"
    }
```

### Logout

Request :
- Method : POST
- Endpoint : `api/v1/logout`
- Header :
    - Content-Type : application/json
    - Accept : application/json
    - Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :
```json
    {
        "message": "Logout successful"
    }
```

### Payment

Request :

- Method : POST
- Endpoint : `api/v1/payment`
- Header :
    - Content-Type : application/json
    - Accept : application/json
    - Authorization : Bearer Token
- Body :
```json
    {
        "customer_id" : "3",
        "merchant_id" : "1",
        "amount"  : 200
    }
```

Response :

- Status : 200 OK
- Body :
```json
    {
        "id": "4",
        "customer_id": "3",
        "merchant_id": "1",
        "amount": 200,
        "created_at": "2024-11-25T18:57:09.805388+07:00",
        "updated_at": "2024-11-25T18:57:09.805388+07:00"
    }
```


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

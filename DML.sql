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
    
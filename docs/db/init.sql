CREATE TABLE merchant_order (
    merchant_id VARCHAR(255) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    status VARCHAR(100),
    last_modified_date TIMESTAMP,
    created_date TIMESTAMP,
    PRIMARY KEY (merchant_id, order_id)
);


INSERT INTO merchant_order (merchant_id, order_id, status, created_date, last_modified_date)
VALUES
('a7eaf5d8-ac11-4768-948c-2ede10a881ff', 101, 'in_preparation', '2024-04-01 08:00:00', '2024-04-01 08:00:00');

INSERT INTO merchant_order (merchant_id, order_id, status, created_date, last_modified_date)
VALUES
('a7eaf5d8-ac11-4768-948c-2ede10a881ff', 102, 'finish', '2024-04-01 09:00:00', '2024-04-01 09:00:00');

INSERT INTO merchant_order (merchant_id, order_id, status, created_date, last_modified_date)
VALUES
('a3bf7fdd-bb5f-41bc-8422-cc979736fc01', 201, 'in_preparation', '2024-04-02 10:30:00', '2024-04-02 10:30:00');

INSERT INTO merchant_order (merchant_id, order_id, status, created_date, last_modified_date)
VALUES
('64cd07b4-9730-4ef0-94db-d6807e5e23a2', 301, 'finish', '2024-04-03 11:15:00', '2024-04-03 11:15:00');

INSERT INTO merchant_order (merchant_id, order_id, status, created_date, last_modified_date)
VALUES
('64cd07b4-9730-4ef0-94db-d6807e5e23a2', 302, 'in_preparation', '2024-04-03 12:45:00', '2024-04-03 12:45:00');

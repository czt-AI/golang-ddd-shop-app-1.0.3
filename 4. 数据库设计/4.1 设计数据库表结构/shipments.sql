CREATE TABLE shippings (
    shipping_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    logistics_company VARCHAR(100),
    tracking_number VARCHAR(100),
    estimated_delivery DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
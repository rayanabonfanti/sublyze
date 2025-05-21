CREATE DATABASE IF NOT EXISTS subscriptions;

USE subscriptions;

CREATE TABLE IF NOT EXISTS subscriptions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    plan VARCHAR(255) NOT NULL,
    status VARCHAR(100) NOT NULL
);

SELECT * FROM subscriptions;
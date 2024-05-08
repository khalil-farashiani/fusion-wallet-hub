-- +migrate Up
CREATE TABLE redeem (
                       id INT AUTO_INCREMENT,
                       title VARCHAR(191) NOT NULL UNIQUE,
                       amount INT NOT NULL,
                       quantity INT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       PRIMARY KEY (id)
);
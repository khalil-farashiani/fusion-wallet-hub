-- +migrate Up
CREATE TABLE transaction (
                        `id` INT AUTO_INCREMENT,
                        `type` VARCHAR(191) NOT NULL,
                        account_id VARCHAR(191) NOT NULL,
                        amount INT NOT NULL,

                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        PRIMARY KEY (id)
);
CREATE TABLE balance (
                             `id` INT AUTO_INCREMENT,
                             account_id VARCHAR(191) NOT NULL UNIQUE,
                             amount INT NOT NULL,

                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             PRIMARY KEY (id)
);
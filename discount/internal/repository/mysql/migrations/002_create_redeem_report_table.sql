-- +migrate Up
CREATE TABLE redeem_report (
                        id INT AUTO_INCREMENT,
                        title VARCHAR(191) NOT NULL UNIQUE,
                        user_id VARCHAR(191) NOT NULL,
                        amount INT NOT NULL,
                        status ENUM('USED', 'NEW') NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        PRIMARY KEY (id)
);
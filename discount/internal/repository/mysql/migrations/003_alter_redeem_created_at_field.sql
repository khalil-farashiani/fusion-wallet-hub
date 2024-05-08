-- +migrate Up
ALTER TABLE redeem
    CHANGE created_at
        created_at TIMESTAMP NOT NULL
            DEFAULT CURRENT_TIMESTAMP
            ON UPDATE CURRENT_TIMESTAMP;
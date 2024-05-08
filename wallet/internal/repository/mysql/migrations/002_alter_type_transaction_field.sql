-- +migrate Up
ALTER TABLE transaction
    CHANGE `type`
    `type` ENUM('IN', 'OUT') NOT NULL;
-- +migrate Up
ALTER TABLE redeem_report
    CHANGE title
    title VARCHAR(191) NULL;


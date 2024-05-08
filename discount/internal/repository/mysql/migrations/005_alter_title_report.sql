-- +migrate Up
ALTER TABLE redeem_report
    MODIFY title VARCHAR(191) null;
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD COLUMN account_locked BOOLEAN;

ALTER TABLE results ADD COLUMN template_id bigint;
ALTER TABLE mail_logs ADD COLUMN template_id integer;
CREATE TABLE IF NOT EXISTS "template_groups" ("id" integer primary key autoincrement,"campaign_id" bigint,"template_id" varchar(255),"group_id" text);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

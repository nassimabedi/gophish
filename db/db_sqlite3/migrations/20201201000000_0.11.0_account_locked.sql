-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD COLUMN account_locked BOOLEAN;

ALTER TABLE results ADD COLUMN template_id bigint;
ALTER TABLE mail_logs ADD COLUMN template_id integer;
ALTER TABLE mail_logs ADD COLUMN profile_id integer;
CREATE TABLE IF NOT EXISTS "template_groups" ("id" integer primary key autoincrement,"campaign_id" bigint, "profile" varchar(255),"template" varchar(255),"groups" text);
CREATE TABLE IF NOT EXISTS "campaign_setting" ("id" integer primary key autoincrement,"duration" bigint);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

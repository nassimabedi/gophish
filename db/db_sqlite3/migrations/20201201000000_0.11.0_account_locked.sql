-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD COLUMN account_locked BOOLEAN;

ALTER TABLE results ADD COLUMN template_id bigint;
ALTER TABLE mail_logs ADD COLUMN template_id integer;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

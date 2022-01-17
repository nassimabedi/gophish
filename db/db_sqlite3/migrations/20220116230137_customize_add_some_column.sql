-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEn-d

ALTER TABLE results ADD COLUMN template_id bigint;
ALTER TABLE mail_logs ADD COLUMN template_id integer;
ALTER TABLE campaignttts ADD COLUMN template_id bigint;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

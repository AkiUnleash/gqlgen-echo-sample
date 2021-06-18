-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts (
  id SERIAL NOT NULL,
  name varchar(255) DEFAULT NULL,
  password varchar(255) DEFAULT NULL,
  completed integer DEFAULT 0,
  created_at TIMESTAMP DEFAULT NULL,
  updated_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY(id)
);
CREATE INDEX account_id on accounts (id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX account_id;
DROP TABLE accounts;
-- +goose StatementEnd
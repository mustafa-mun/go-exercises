-- +goose Up
CREATE TABLE phone_numbers (
  phone_number VARCHAR(15) PRIMARY KEY
);

-- +goose Down
DROP TABLE phone_numbers;
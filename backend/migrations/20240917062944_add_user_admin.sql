-- +goose Up
INSERT INTO users (id, username, password, role)
VALUES (1, 'admin', '$2a$14$qHAl0aORLjeLySkhRTDJvOfMsLm4hWC438g.TLldpV0NcxmQomH/m', 'ADMIN');
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DELETE FROM items
WHERE owner_id = 1;

DELETE FROM users
WHERE username = 'admin';
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

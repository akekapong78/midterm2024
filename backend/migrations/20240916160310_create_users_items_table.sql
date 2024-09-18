-- +goose Up
-- goose postgres "postgres://postgres:password@localhost:5432/task" up
-- สร้างตาราง users เพื่อจัดเก็บข้อมูลผู้ใช้งาน
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
);

-- สร้างตาราง items เพื่อจัดเก็บข้อมูลคำร้องขอเบิก
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    owner_id INT,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- ลบตาราง items ก่อนเพราะมี foreign key เชื่อมโยงกับ users
DROP TABLE IF EXISTS items;

-- ลบตาราง users
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
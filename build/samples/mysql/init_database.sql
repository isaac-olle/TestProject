CREATE TABLE users (
    id BINARY(16) PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    surname VARCHAR(64) NOT NULL,
    birthDate DATE,
    email VARCHAR(256) UNIQUE NOT NULL,
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    isDeleted BOOL NOT NULL DEFAULT 0
);

CREATE TRIGGER users_id
    BEFORE INSERT ON users
    FOR EACH ROW
    BEGIN
        SET NEW.id = UNHEX(REPLACE(UUID(), '-', ''));
    END;

CREATE VIEW users_view AS
SELECT
    CONCAT_WS('-',
        SUBSTR(HEX(id), 1, 8),
        SUBSTR(HEX(id), 9, 4),
        SUBSTR(HEX(id), 13, 4),
        SUBSTR(HEX(id), 17, 4),
        SUBSTR(HEX(id), 21)
    ) AS uuid,
    name,
    email,
    birthDate,
    createdAt,
    updatedAt,
    isDeleted
FROM users;
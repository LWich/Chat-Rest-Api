CREATE TABLE users (
    Id BIGSERIAL NOT NULL PRIMARY KEY,
    Email VARCHAR(40) NOT NULL UNIQUE,
    Encrypted_Password VARCHAR(30) NOT NULL
);
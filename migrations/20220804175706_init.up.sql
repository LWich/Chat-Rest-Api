CREATE TABLE users (
    Id BIGSERIAL NOT NULL PRIMARY KEY,
    Email VARCHAR(40) NOT NULL UNIQUE,
    Encrypted_Password VARCHAR NOT NULL,
    Expires_In BIGINT,
    Refresh_Token UUID
);
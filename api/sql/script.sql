IF OBJECT_ID(N'devbook') IS NULL
BEGIN
    CREATE DATABASE devbook;
END;
GO

USE devbook;

GO

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT NOT NULL IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at datetime2 DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime2  DEFAULT CURRENT_TIMESTAMP,    
);



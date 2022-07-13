IF OBJECT_ID(N'devbook') IS NULL
BEGIN
    CREATE DATABASE devbook;
END;
GO

USE devbook;

GO

DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    id INT NOT NULL IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at datetime2 DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime2 DEFAULT CURRENT_TIMESTAMP,
);

GO

CREATE TABLE followers
(
    user_id INT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users (id) ON DELETE CASCADE,

    follower_id INT NOT NULL,
    FOREIGN KEY(follower_id) REFERENCES users (id),

    created_at datetime2 DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime2 DEFAULT CURRENT_TIMESTAMP,

    PRIMARY key (user_id, follower_id)
);


--DROP DATABASE devbook 
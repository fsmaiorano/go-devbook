IF OBJECT_ID(N'devbook') IS NULL
BEGIN
    CREATE DATABASE devbook;
END;
GO

USE devbook;

GO

DROP TABLE IF EXISTS posts;
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

GO

CREATE TABLE posts
(
    id INT NOT NULL IDENTITY PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY(author_id) REFERENCES users (id) ON DELETE CASCADE,
    likes INT NOT NULL DEFAULT 0,
    created_at datetime2 DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime2 DEFAULT CURRENT_TIMESTAMP,
);



-- drop database devbook
-- select * from users
-- select * from followers
select *
from posts
-- -- followers
-- SELECT u.id, u.name, u.nickname, u.email, u.created_at, u.updated_at FROM users u INNER JOIN followers f on u.id = f.follower_id where f.user_id = 1

-- -- following
-- SELECT u.id, u.name, u.nickname, u.email, u.created_at, u.updated_at FROM users u INNER JOIN followers f on u.id = f.user_id where f.follower_id = 2

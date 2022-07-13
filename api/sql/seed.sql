USE devbook;

GO

INSERT INTO users
    (name, nickname, email, password)
VALUES
    ('John Doe', 'jdoe', 'jdoe@jdoe.com', 'password123'),
    ('Jane Doe', 'jdoe', 'jdoe@jdoe.com', 'password123'),
    ('John Smith', 'jsmith', 'jsmith@jsmith.com', 'password123'),
    ('Jane Smith', 'jsmith', 'jsmith@jsmith.com', 'password123');

GO

INSERT INTO followers
    (user_id, follower_id, created_at, updated_at)
VALUES
    (1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (1, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
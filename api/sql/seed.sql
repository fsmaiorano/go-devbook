USE devbook;

GO

INSERT INTO devbook.dbo.users
    (name, nickname, email, password)
VALUES
    ('Fábio Maiorano', 'fsmaiorano', 'fsmaiorano@gmail.com', '$2a$10$BbRXTKFW0/3XUDWs94t/nOdJqWozgDhh9zRJ6PaxPLPRSpCRtNJtq'),
    ('John Doe', 'jdoe', 'jdoe@jdoe.com', '$2a$10$BbRXTKFW0/3XUDWs94t/nOdJqWozgDhh9zRJ6PaxPLPRSpCRtNJtq'),
    ('Jane Doe', 'JaneDoe', 'Jane@jdoe.com', '$2a$10$BbRXTKFW0/3XUDWs94t/nOdJqWozgDhh9zRJ6PaxPLPRSpCRtNJtq'),
    ('John Smith', 'jsmith', 'John@jsmith.com', '$2a$10$BbRXTKFW0/3XUDWs94t/nOdJqWozgDhh9zRJ6PaxPLPRSpCRtNJtq'),
    ('John Smith', 'jsmith', 'John@jsmith.com', '$2a$10$BbRXTKFW0/3XUDWs94t/nOdJqWozgDhh9zRJ6PaxPLPRSpCRtNJtq');
   

-- seed PASSWORD = 'password'    

GO

INSERT INTO devbook.dbo.followers
    (user_id, follower_id, created_at, updated_at)
VALUES
    (1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (1, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

GO

INSERT INTO devbook.dbo.posts
    (title, content, author_id)
VALUES
    ('Post 1', 'This is the first post', 1),
    ('Post 2', 'This is the second post', 1),
    ('Post 3', 'This is the third post', 3),
    ('Post 4', 'This is the fourth post', 4),
    ('Post 5', 'This is the fifth post', 5);

CREATE TABLE IF NOT EXISTS comment_likes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    commentId INT NOT NULL,
    userId INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_comments_likes FOREIGN KEY (userId) REFERENCES users(id),
    CONSTRAINT fk_comment_id_comment_likes FOREIGN KEY (commentId) REFERENCES comments(id)
)
--& create the db if not exist
    CREATE DATABASE IF NOT EXISTS chat_app;
    USE chat_app;
--& create the user table structure
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
--& create the message table structure
    CREATE TABLE IF NOT EXISTS messages (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        username VARCHAR(50) NOT NULL,
        message TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
        CONSTRAINT `fk_the_message_far_user` 
            FOREIGN KEY (`user_id`) 
            REFERENCES `users` (`id`) 
            ON DELETE CASCADE,
        INDEX idx_created_at (created_at)
    );
--& seed user as a placeholder data
    INSERT INTO users (username) VALUES ('Badr'), ('Ali');
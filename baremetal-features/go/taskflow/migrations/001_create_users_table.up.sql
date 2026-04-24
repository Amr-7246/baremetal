--& User table
    CREATE TABLE IF NOT EXISTS users(
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- from where that function coming and what it do
        email VARCHAR(255) UNIQUE NOT NULL,
        hashed_password VARCHAR(255) UNIQUE NOT NULL,
        full_name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    );

--& index for faster email lookups
    CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);
--& The function to automatically update updated_at
    --? Explain me that function, both syntax and logic
    CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
    BEGIN 
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
    $$ language 'plpgsql'
--& The trigger to update updated_at on row update
    CREATE TRIGGER update_users_updated_at 
        BEFORE UPDATE ON users
        FOR EACH ROW 
        EXISTS FUNCTION update_updated_at_column()
--& seed a test user data 
    INSERT INTO users (id, email, hashed_password, full_name) VALUES (
        '123e4567-e89b-12d3-a456-426614174000',
        'test@example.com',
        '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mr/.jRpF9lFtM3C3VXJQkOqXZvF5Rji',
        'Test User'
    ) ON CONFLICT (email) DO NOTHING; --? what does that line do 

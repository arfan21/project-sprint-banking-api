CREATE TABLE
    IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        createdAt TIMESTAMP DEFAULT now (),
        updatedAt TIMESTAMP DEFAULT now ()
    );

CREATE INDEX IF NOT EXISTS users_id_hash ON users USING HASH (id);

CREATE INDEX IF NOT EXISTS users_email_hash ON users USING HASH (email);

CREATE TRIGGER update_users_updated_at BEFORE
UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated ();
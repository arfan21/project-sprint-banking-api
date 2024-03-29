CREATE TABLE
    IF NOT EXISTS balances (
        userId UUID NOT NULL,
        currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
        balance NUMERIC(20, 2) NOT NULL DEFAULT 0,
        createdAt TIMESTAMP DEFAULT now (),
        updatedAt TIMESTAMP DEFAULT now (),
        PRIMARY KEY (userId, currency),
        CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TRIGGER update_balances_updated_at BEFORE
UPDATE ON balances FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated ();
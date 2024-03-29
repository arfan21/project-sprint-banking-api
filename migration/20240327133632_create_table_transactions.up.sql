CREATE TABLE
    IF NOT EXISTS transactions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        userId UUID NOT NULL,
        currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
        amount NUMERIC(20, 2) NOT NULL,
        bankName VARCHAR(255) NOT NULL,
        bankAccountNumber VARCHAR(255) NOT NULL,
        transferProofImg VARCHAR(255) NULL,
        createdAt TIMESTAMP DEFAULT now (),
        updatedAt TIMESTAMP DEFAULT now (),
        CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TRIGGER update_transactions_updated_at BEFORE
UPDATE ON transactions FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated ();
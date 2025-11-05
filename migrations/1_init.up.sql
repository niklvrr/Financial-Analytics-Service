BEGIN;

CREATE TABLE bank_accounts (
                               id SERIAL PRIMARY KEY,
                               name TEXT NOT NULL,
                               balance NUMERIC(18,2) NOT NULL DEFAULT 0
);

CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            kind TEXT NOT NULL,
                            name TEXT NOT NULL
);

CREATE TABLE operations (
                            id SERIAL PRIMARY KEY,
                            kind TEXT NOT NULL,
                            bank_account_id INT NOT NULL
                                REFERENCES bank_accounts(id) ON UPDATE CASCADE ON DELETE RESTRICT,
                            amount NUMERIC(18,2) NOT NULL,
                            date TIMESTAMP NOT NULL,
                            description TEXT,
                            category_id INT NOT NULL
                                REFERENCES categories(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE INDEX idx_operations_date ON operations (date);
CREATE INDEX idx_operations_bank_account_id ON operations (bank_account_id);
CREATE INDEX idx_operations_category_id ON operations (category_id);

COMMIT;

ALTER TABLE transactions
DROP COLUMN IF EXISTS transaction_date,
DROP COLUMN IF EXISTS virtual_account;
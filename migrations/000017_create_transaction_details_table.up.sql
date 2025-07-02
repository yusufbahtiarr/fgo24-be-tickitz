CREATE TABLE transaction_details(
  id_transaction INT NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
  seat VARCHAR(255) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);
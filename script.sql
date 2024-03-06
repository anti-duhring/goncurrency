CREATE TABLE
  clients (
    id INT PRIMARY KEY,
    account_limit INT NOT NULL,
    balance INT NOT NULL DEFAULT 0
  );

CREATE UNLOGGED TABLE
  transactions (
    id SERIAL PRIMARY KEY,
    client_id SMALLINT NOT NULL,
    amount INT NOT NULL,
    operation CHAR(1) NOT NULL,
    description VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (client_id) REFERENCES clients (id)
  );

ALTER TABLE
  transactions
SET
  (autovacuum_enabled = false);

DO
  $$ BEGIN INSERT INTO clients (id, account_limit) 
VALUES 
  (1, 100000), 
  (2, 80000), 
  (3, 1000000), 
  (4, 10000000), 
  (5, 500000);
END;
$$

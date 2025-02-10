-- 001_create_tables.up.sql

CREATE TABLE IF NOT EXISTS blocks (
  number BIGINT PRIMARY KEY,
  hash TEXT NOT NULL,
  parent_hash TEXT NOT NULL,
  timestamp BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    hash TEXT PRIMARY KEY,
    "from" TEXT NOT NULL,
    "to" TEXT NOT NULL,
    value TEXT NOT NULL,
    block_number BIGINT NOT NULL,
    CONSTRAINT fk_block FOREIGN KEY (block_number) REFERENCES blocks(number)
);

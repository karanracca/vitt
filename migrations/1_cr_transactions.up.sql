CREATE TABLE IF NOT EXISTS transactions (
	id TEXT PRIMARY KEY,
	hash TEXT,
	date TEXT,
	description TEXT,
	acc_type TEXT,
  source TEXT,
	dest TEXT,
  direction TEXT,
	amount REAL,
	comments TEXT
);
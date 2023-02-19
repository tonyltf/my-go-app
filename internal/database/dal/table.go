package dal

const RateTable = `
CREATE TABLE IF NOT EXISTS rate (
	id VARCHAR(36) NOT NULL,
	currency_pair VARCHAR(100) NOT NULL,
	exchange_rate REAL NOT NULL,
	created_at DATETIME,
	PRIMARY KEY (id)
)`

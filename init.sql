CREATE TABLE uptime_logs (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255),
    url TEXT,
    status VARCHAR(50),
    last_down TIMESTAMP,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

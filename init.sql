CREATE TABLE uptime_logs (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255),
    url TEXT,
    status VARCHAR(50),
    last_down TIMESTAMP,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subscribers (
    email_id VARCHAR(255) PRIMARY KEY
);

INSERT INTO uptime_logs (service_name, url, status, last_down)
VALUES
    ('jaypatel', 'https://jaypatel.link', 'UP', '2024-11-15 10:30:00'),
    ('dev', 'https://dev.jaypatel.link', 'UP', '2024-12-10 10:30:00'),
    ('app', 'https://app.jaypatel.link', 'UP', '2024-12-05 14:00:00'),
    ('res', 'https://res.jaypatel.link', 'UP', '2024-12-15 14:00:00'),
    ('magicdot', 'https://magicdot.jaypatel.link', 'UP', '2024-12-02 10:30:00'),
    ('uptime', 'https://uptime.jaypatel.link', 'DOWN', '2024-12-16 14:00:00');
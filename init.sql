CREATE TABLE uptime_logs (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255),
    url TEXT,
    status VARCHAR(50),
    last_down TIMESTAMP,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CREATE TABLE subscribers (
--     email_id VARCHAR(255) PRIMARY KEY
-- );

CREATE TABLE subscribers (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    token VARCHAR(255) NOT NULL UNIQUE,
    is_verified BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    link1 VARCHAR(255),
    link2 VARCHAR(255),
    link3 VARCHAR(255)
);

INSERT INTO users (email, username, link1) VALUES('test@gmail.com','test','https://test.com');

INSERT INTO uptime_logs (service_name, url, status, last_down)
VALUES
    ('jaypatel', 'https://jaypatel.link', 'UP', '2024-11-15 10:30:00'),
    ('dev', 'https://dev.jaypatel.link', 'UP', '2024-12-10 10:30:00'),
    ('app', 'https://app.jaypatel.link', 'UP', '2024-12-05 14:00:00'),
    ('res', 'https://res.jaypatel.link', 'UP', '2024-12-15 14:00:00'),
    ('magicdot', 'https://magicdot.jaypatel.link', 'UP', '2024-12-02 10:30:00'),
    ('uptime', 'https://uptime.jaypatel.link', 'DOWN', '2024-12-16 14:00:00');
CREATE TABLE scheduler (
    id INT IDENTITY(1,1) PRIMARY KEY,
    date CHAR(8) NOT NULL DEFAULT '',
    comment TEXT NOT NULL DEFAULT '',
    title VARCHAR NOT NULL DEFAULT '',
    repeat VARCHAR NOT NULL DEFAULT ''
);
CREATE INDEX idx_scheduler_date ON scheduler(date);
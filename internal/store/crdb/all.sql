CREATE TABLE namespace (
    name VARCHAR(64),
    config JSONB,
    commit_time TIMESTAMP,
    PRIMARY KEY (name, commit_time)
);
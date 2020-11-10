CREATE TABLE namespace_config (
    namespace VARCHAR,
    config JSONB,
    commit_time TIMESTAMP,
    PRIMARY KEY (namespace, commit_time)
);
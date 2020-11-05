CREATE TABLE relation_tuples (
    shard_id VARCHAR,
    object_id VARCHAR,
    relation VARCHAR,
    user VARCHAR,
    commit_time TIMESTAMP,
    PRIMARY KEY (shard_id, object_id, relation, user, commit_time)
);
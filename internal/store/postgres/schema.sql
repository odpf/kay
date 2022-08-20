CREATE TABLE IF NOT EXISTS clusters
(
    id                 BIGSERIAL NOT NULL PRIMARY KEY,
    urn                TEXT      NOT NULL UNIQUE,
    name               TEXT      NOT NULL,
    created_at         timestamp NOT NULL DEFAULT current_timestamp,
    updated_at         timestamp NOT NULL DEFAULT current_timestamp,
    bootstrap_servers  bytea     NOT NULL
);
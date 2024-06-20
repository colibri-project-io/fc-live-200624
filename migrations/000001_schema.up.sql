-- CREATE UUID EXTENSION
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CREATE SCHEMA
CREATE TABLE IF NOT EXISTS cities (
    id         UUID      NOT NULL DEFAULT uuid_generate_v1mc(),
    name       TEXT      NOT NULL,
    uf         TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT cities_pk PRIMARY KEY (id),
    CONSTRAINT cities_name_uf_un UNIQUE (name, uf)
);

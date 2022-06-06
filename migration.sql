
CREATE EXTENSION IF NOT EXISTS earthdistance CASCADE;

DROP TABLE places;

CREATE TABLE IF NOT EXISTS places (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(200) NOT NULL,
    description VARCHAR(500) NOT NULL,
    resource_id BIGINT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    lat double precision NOT NULL,
    lng double precision NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX places_is_deleted_idx ON places USING hash (is_deleted);
CREATE INDEX places_location_idx ON places USING gist (ll_to_earth(lat, lng));

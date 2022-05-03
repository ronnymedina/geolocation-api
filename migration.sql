
CREATE EXTENSION IF NOT EXISTS earthdistance CASCADE;

DROP TABLE places;

CREATE TABLE IF NOT EXISTS places (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(200) NOT NULL,
    description VARCHAR(500) NOT NULL,
    resource_id BIGINT NOT NULL,
    lat double precision NOT NULL,
    lng double precision NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX places_name_low_idx ON places USING btree (lower((name)::text));
CREATE INDEX places_location_idx ON places USING gist (ll_to_earth(lat, lng));

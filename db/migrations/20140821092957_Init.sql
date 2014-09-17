-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE event (
	id serial PRIMARY KEY,
	name varchar(255),
	auth_method varchar(255),
	auth_method_config json
);
CREATE UNIQUE INDEX event_name ON event (name);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE event;

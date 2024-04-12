-- +migrate Up

-- Create the car table
CREATE TABLE IF NOT EXISTS car (
    id SERIAL PRIMARY KEY,
    reg_nums TEXT[],
    mark VARCHAR(255),
    model VARCHAR(255),
    year INT,
    owner_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the index for owner_id if needed
-- CREATE INDEX IF NOT EXISTS car_owner_id_idx ON car (owner_id);

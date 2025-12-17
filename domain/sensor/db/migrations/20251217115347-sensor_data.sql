-- +migrate Up
CREATE TABLE IF NOT EXISTS sensor_data (
    id BIGSERIAL PRIMARY KEY,
    device_id VARCHAR(100) NOT NULL,        -- ID sensor
    temperature DOUBLE PRECISION NOT NULL,
    humidity DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index untuk query per device dan time-series
CREATE INDEX IF NOT EXISTS idx_sensor_data_device_id
ON sensor_data (device_id);

CREATE INDEX IF NOT EXISTS idx_sensor_data_created_at
ON sensor_data (created_at DESC);

-- +migrate Down
DROP TABLE IF EXISTS sensor_data;
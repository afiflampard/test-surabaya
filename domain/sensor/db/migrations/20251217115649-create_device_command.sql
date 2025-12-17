-- +migrate Up
CREATE TABLE IF NOT EXISTS device_commands (
    id BIGSERIAL PRIMARY KEY,
    device_id VARCHAR(100) NOT NULL,         -- ID device
    command VARCHAR(10) NOT NULL,            -- ON / OFF
    status VARCHAR(20) NOT NULL DEFAULT 'SENT',  -- SENT / DELIVERED / FAILED
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index untuk query per device dan time-series
CREATE INDEX IF NOT EXISTS idx_device_commands_device_id
ON device_commands (device_id);

CREATE INDEX IF NOT EXISTS idx_device_commands_created_at
ON device_commands (created_at DESC);

-- +migrate Down
DROP TABLE IF EXISTS device_commands;

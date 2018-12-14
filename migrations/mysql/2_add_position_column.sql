-- +migrate Up
ALTER TABLE congestion_log ADD COLUMN position INT NOT NULL;
-- +migrate Up
ALTER TABLE images CHANGE path body BLOB NOT NULL;

-- +migrate Up

INSERT INTO buses(id, name) VALUES(1, '1'),(2, '2'),(3, '3'),(4, '4');

-- +migrate Down

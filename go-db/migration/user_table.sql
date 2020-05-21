CREATE TABLE IF NOT EXISTS db_user (
                            id INT NOT NULL PRIMARY KEY,
                            name varchar(20) NOT NULL,
                            contacts varchar[] NOT NULL,
                            createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO db_user(id, name, contacts) VALUES(1, 'user-1', '{"1234567890A", "2345678901A"}');
INSERT INTO db_user(id, name, contacts) VALUES(2, 'user-2', '{"3456789012A", "4567890123A"}');
INSERT INTO db_user(id, name, contacts) VALUES(3, 'user-3', '{"5678901234A", "6789012345A"}');

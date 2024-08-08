CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  text TEXT NOT NULL
);

SELECT * FROM messages;

INSERT INTO messages(text) VALUES ('test');
INSERT INTO messages(text) VALUES ('test2');
INSERT INTO messages(text) VALUES ('test3');

DROP TABLE messages;
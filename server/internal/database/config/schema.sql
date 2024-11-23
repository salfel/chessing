CREATE TABLE users (
    id          INTEGER PRIMARY KEY,
    username    text NOT NULL,
    password    text NOT NULL
);

CREATE TABLE games (
    id      INTEGER PRIMARY KEY,
    white   INTEGER,
    black   INTEGER,
    FOREIGN KEY (black) REFERENCES users(id),
    FOREIGN KEY (white) REFERENCES users(id)
);

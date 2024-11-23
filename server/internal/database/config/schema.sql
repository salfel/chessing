DROP TABLE users;
DROP TABLE games;
DROP TABLE user_games;

CREATE TABLE users (
    id      INTEGER PRIMARY KEY,
    name    text    NOT NULL
);

CREATE TABLE games (
    id      INTEGER PRIMARY KEY
);

CREATE TABLE user_games (
    user_id INTEGER NOT NULL,
    game_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, game_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (game_id) REFERENCES games(id)
);

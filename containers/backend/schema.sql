CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS player;

CREATE TABLE faction (
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    name VARCHAR(20) NOT NULL UNIQUE,
    color INT NOT NULL UNIQUE,
    PRIMARY KEY (uuid)
);

CREATE TABLE player (
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    nickname VARCHAR(16) NOT NULL UNIQUE,
    registered BOOLEAN NOT NULL DEFAULT false,
    faction UUID DEFAULT uuid_nil(), 
    PRIMARY KEY (uuid),
    FOREIGN KEY (faction) REFERENCES faction(uuid)
);

INSERT INTO faction (uuid, name, color) VALUES
    (uuid_nil(), 'nil', 0);

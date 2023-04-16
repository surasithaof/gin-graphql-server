CREATE TABLE
    IF NOT EXISTS teams (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        country VARCHAR(255) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS players (
        id serial PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        rating NUMERIC(3, 2) NOT NULL,
        team_id INTEGER NOT NULL,
        CONSTRAINT fk_team FOREIGN KEY(team_id) REFERENCES teams(id)
    );
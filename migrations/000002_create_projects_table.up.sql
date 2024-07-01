CREATE TABLE projects(

        id              SERIAL PRIMARY KEY,
        name     VARCHAR(30) NOT NULL UNIQUE,
        user_id          INTEGER NOT NULL,
        created_at      TIMESTAMP,
        updated_at      TIMESTAMP

);
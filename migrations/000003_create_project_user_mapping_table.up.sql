CREATE TABLE project_user_mappings(
        id              SERIAL PRIMARY KEY,
        project_id       INTEGER NOT NULL,
        user_id          INTEGER NOT NULL
);
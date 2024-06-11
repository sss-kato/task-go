CREATE TABLE projects(

        id              SERIAL PRIMARY KEY,
        projectname     VARCHAR(30) NOT NULL UNIQUE,
        userid          INTEGER NOT NULL,
        created_at      TIMESTAMP,
        updated_at      TIMESTAMP

);
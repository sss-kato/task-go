CREATE TABLE project_user_mappings(

        id              SERIAL PRIMARY KEY,
        projectid       INTEGER NOT NULL,
        userid          INTEGER NOT NULL
);
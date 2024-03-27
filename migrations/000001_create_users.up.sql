CREATE TABLE users(

        id SERIAL PRIMARY KEY,
        name VARCHAR(15) NOT NULL UNIQUE,
        mailadress VARCHAR(30) NOT NULL UNIQUE,
        password VARCHAR(15) NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP

);
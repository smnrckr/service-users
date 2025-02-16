CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) not null,
    lastname VARCHAR(255) not null,
    profile_photo VARCHAR(255)
);
CREATE TABLE users(
    id serial primary key,
    first_name text not null,
    last_name text not null
);

INSERT INTO users(first_name, last_name) VALUES ('Grace', 'Hopper');
INSERT INTO users(first_name, last_name) VALUES ('Ada', 'Lovelace');
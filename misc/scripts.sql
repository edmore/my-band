CREATE DATABASE band;

\c band

CREATE TABLE members (  
  id SERIAL PRIMARY KEY,
  name TEXT,
  surname TEXT,
  speciality TEXT
);

\dt

INSERT INTO members (name, surname, speciality)  
VALUES ('Edmore', 'Moyo', 'vocalist');
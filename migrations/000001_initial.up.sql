CREATE TABLE IF NOT EXISTS authors
  (
     created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     id             SERIAL PRIMARY KEY,
     name           VARCHAR NOT NULL,
     lastname       VARCHAR NOT NULL,
     username       VARCHAR NOT NULL,
     specialization VARCHAR NOT NULL
  );

CREATE TABLE IF NOT EXISTS books
  (
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     id         SERIAL PRIMARY KEY,
     title      VARCHAR NOT NULL,
     genre      VARCHAR NOT NULL,
     id_author_books VARCHAR NOT NULL,
     code_isbn  INTEGER NOT NULL
  );


CREATE TABLE IF NOT EXISTS members
  (
     created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     id             SERIAL PRIMARY KEY,
     name           VARCHAR NOT NULL,
     lastname       VARCHAR NOT NULL,
     borrowed_books VARCHAR NOT NULL,
     member_id_books VARCHAR NOT NULL
  );
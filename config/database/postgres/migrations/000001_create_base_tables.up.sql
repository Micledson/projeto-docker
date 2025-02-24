--- Creating Extensions and Functions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE todo (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  description VARCHAR(30) NOT NULL,
  is_active          BOOLEAN default true,
  created_at         TIMESTAMP        DEFAULT NOW(),
  updated_at         TIMESTAMP        DEFAULT NOW()
);

COPY todo(id, description, is_active)
    FROM '/fixtures/000001/todo.csv'
    DELIMITER ';' csv header;
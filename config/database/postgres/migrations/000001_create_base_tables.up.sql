--- Creating Extensions and Functions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE todo (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  description VARCHAR(30) NOT NULL,
  is_active          BOOLEAN,
  created_at         TIMESTAMP        DEFAULT NOW(),
  updated_at         TIMESTAMP        DEFAULT NOW()
);

set search_path = gym, public;
CREATE SCHEMA gym;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE gym."results" ("id" uuid UNIQUE DEFAULT uuid_generate_v4(),"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"text" text,"tags" text,PRIMARY KEY ("id"));

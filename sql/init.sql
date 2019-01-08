-- TABLE "users" --
CREATE TABLE IF NOT EXISTS users (
  "id"        BIGINT UNIQUE NOT NULL PRIMARY KEY,
  "firstname" CITEXT NOT NULL,
  "lastname"  CITEXT NOT NULL,
  "gender"    CHAR[1],
  "age"       INT
);


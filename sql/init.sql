-- TABLE "users" --
CREATE TABLE IF NOT EXISTS users (
  "id"        BIGINT UNIQUE NOT NULL PRIMARY KEY,
  "firstname" CITEXT NOT NULL,
  "lastname"  CITEXT NOT NULL,
  "sex"       CHAR[1],
  "age"       INT
);

-- TABLE "orders" --
CREATE TABLE IF NOT EXISTS orders (
  "id"          SERIAL UNIQUE PRIMARY KEY,
  "user_id"     BIGINT NOT NULL,
  "firstname"   CITEXT DEFAULT NULL,
  "lastname"    CITEXT DEFAULT NULL,
  "phone"       CITEXT DEFAULT NULL,
  "company"     CITEXT DEFAULT NULL,
  "address"     CITEXT DEFAULT NULL,
  "date"        TIMESTAMPTZ(3) DEFAULT NULL,
  "state"       INT NOT NULL DEFAULT 0,
  "is_finished" BOOLEAN DEFAULT FALSE
)
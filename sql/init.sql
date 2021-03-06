-- TABLE "users" --
CREATE TABLE IF NOT EXISTS users (
  "id"        INT PRIMARY KEY,
  "firstname" CITEXT NOT NULL,
  "lastname"  CITEXT NOT NULL,
  "age"       INT,
  "sex"       INT
);

-- TABLE "orders" --
CREATE TABLE IF NOT EXISTS orders (
  "id"            SERIAL PRIMARY KEY,
  "user_id"       INT NOT NULL,
  "firstname"     CITEXT DEFAULT NULL,
  "lastname"      CITEXT DEFAULT NULL,
  "phone"         CITEXT DEFAULT NULL,
  "company"       CITEXT DEFAULT NULL,
  "address"       CITEXT DEFAULT NULL,
  "delivery_date" CITEXT DEFAULT NULL,
  "order_date"    TIMESTAMPTZ(3) DEFAULT NULL
);

-- TABLE "states" --
CREATE TABLE IF NOT EXISTS states (
  "user_id"     INT PRIMARY KEY,
  "state"       INT DEFAULT 0 NOT NULL
);
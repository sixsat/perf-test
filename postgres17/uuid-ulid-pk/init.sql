-- UUID v4
CREATE TABLE
    test1 (id uuid PRIMARY KEY DEFAULT gen_random_uuid ());

-- UUID v7
CREATE EXTENSION pg_uuidv7;

CREATE TABLE
    test2 (id uuid PRIMARY KEY DEFAULT uuid_generate_v7 ());

-- ULID
CREATE EXTENSION ulid;

CREATE TABLE
    test3 (id ulid PRIMARY KEY DEFAULT gen_ulid ());
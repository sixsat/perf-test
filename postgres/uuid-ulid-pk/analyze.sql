EXPLAIN ANALYSE
INSERT INTO
    test1 (id)
SELECT
    gen_random_uuid ()
FROM
    generate_series (1, 1000000);

EXPLAIN ANALYSE
INSERT INTO
    test2 (id)
SELECT
    uuid_generate_v7 ()
FROM
    generate_series (1, 1000000);

EXPLAIN ANALYSE
INSERT INTO
    test3 (id)
SELECT
    gen_ulid ()
FROM
    generate_series (1, 1000000);
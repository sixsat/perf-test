SELECT
    pg_size_pretty (pg_total_relation_size ('test1')) AS uuidv4;

SELECT
    pg_size_pretty (pg_total_relation_size ('test2')) AS uuidv7;

SELECT
    pg_size_pretty (pg_total_relation_size ('test3')) AS ulid;
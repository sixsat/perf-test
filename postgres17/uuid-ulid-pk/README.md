# Postgres 17 Primary Key Benchmark: UUIDv4 vs UUIDv7 vs ULID

## 💻 System Information

| Component | Detail       |
| --------- | ------------ |
| OS        | macOS 15.4.1 |
| CPU       | Apple M4     |
| RAM       | 16 GB        |

## 🐳 Docker Desktop Resources

| Resource | Allocation |
| -------- | ---------- |
| CPU      | 2          |
| Memory   | 2 GB       |
| Swap     | 0          |
| Disk     | 128 GB     |

## 📋 Prerequisites

```sh
make build
make start
```

## 📊 Results

- Execution time: `make analyze`

```sql
QUERY PLAN
----------------------------------------------------------------------------------------------------------------------------------------
 Insert on test1  (cost=0.00..12500.00 rows=0 width=0) (actual time=1796.996..1796.997 rows=0 loops=1)
   ->  Function Scan on generate_series  (cost=0.00..12500.00 rows=1000000 width=16) (actual time=31.954..664.695 rows=1000000 loops=1)
 Planning Time: 0.034 ms
 Execution Time: 1797.557 ms
(4 rows)

QUERY PLAN
----------------------------------------------------------------------------------------------------------------------------------------
 Insert on test2  (cost=0.00..12500.00 rows=0 width=0) (actual time=1546.000..1546.000 rows=0 loops=1)
   ->  Function Scan on generate_series  (cost=0.00..12500.00 rows=1000000 width=16) (actual time=27.914..678.474 rows=1000000 loops=1)
 Planning Time: 0.022 ms
 Execution Time: 1546.595 ms
(4 rows)

QUERY PLAN
----------------------------------------------------------------------------------------------------------------------------------------
 Insert on test3  (cost=0.00..12500.00 rows=0 width=0) (actual time=1520.148..1520.149 rows=0 loops=1)
   ->  Function Scan on generate_series  (cost=0.00..12500.00 rows=1000000 width=16) (actual time=25.126..670.035 rows=1000000 loops=1)
 Planning Time: 0.023 ms
 Execution Time: 1520.624 ms
(4 rows)
```

- Total disk usage: `make size`

```sql
 uuidv4
--------
 80 MB
(1 row)

 uuidv7
--------
 81 MB
(1 row)

 ulid
-------
 82 MB
(1 row)
```

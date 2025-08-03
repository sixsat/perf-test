FROM postgres:17 AS base

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    git \
    postgresql-server-dev-17

# install pg_uuidv7 extension
WORKDIR /src/pg_uuidv7
RUN git clone https://github.com/fboulnois/pg_uuidv7.git .
RUN make && make install

# install pg-ulid extension
WORKDIR /src/pg-ulid
RUN git clone https://github.com/andrielfn/pg-ulid.git .
RUN make && make install

FROM postgres:17

COPY --from=base /usr/lib/postgresql/17/lib/pg_uuidv7.so /usr/lib/postgresql/17/lib/
COPY --from=base /usr/share/postgresql/17/extension/pg_uuidv7.control /usr/share/postgresql/17/extension/
COPY --from=base /usr/share/postgresql/17/extension/pg_uuidv7--1.6.sql /usr/share/postgresql/17/extension/

COPY --from=base /usr/lib/postgresql/17/lib/ulid.so /usr/lib/postgresql/17/lib/
COPY --from=base /usr/share/postgresql/17/extension/ulid.control /usr/share/postgresql/17/extension/
COPY --from=base /usr/share/postgresql/17/extension/ulid--0.0.1.sql /usr/share/postgresql/17/extension/
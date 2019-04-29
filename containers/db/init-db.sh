#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  grant all privileges on database itpdb to itp;
  create table vm (
    id serial primary key,
    vm_id varchar(50) not null,
    vm_name varchar(50) not null,
    vm_os varchar(20) not null
  );
  create table vm_spec (
    vm_id integer not null,
    cpus integer,
    memory varchar(50),
    storage varchar(50),
    extras varchar
  );
EOSQL
#!/bin/bash
docker run --name a-postgres \
       -v /Users/mrykov/tmp/postgres-experiments/a-data:/var/lib/postgresql/data \
       -v /Users/mrykov/tmp/postgres-experiments/a-config/repmgr.d:/etc/repmgr.d \
       -v /Users/mrykov/tmp/postgres-experiments/a-config/barman.d:/etc/barman.d \
       -e POSTGRES_PASSWORD=mysecretpassword \
       -d drunk-pg


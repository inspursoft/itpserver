FROM postgres

COPY "containers/db/init-db.sh" "/docker-entrypoint-initdb.d"
RUN chmod u+x /docker-entrypoint-initdb.d/init-db.sh
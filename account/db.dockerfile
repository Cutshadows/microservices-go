FROM postgres:15.15-trixie

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]
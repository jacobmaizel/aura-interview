FROM alpine:latest

WORKDIR /app
COPY ../scripts/run_migrations.sh /usr/local/bin/
COPY ../migrations/init.sql /app/migrations/init.sql
RUN chmod +x /usr/local/bin/run_migrations.sh

CMD ["/usr/local/bin/run_migrations.sh"]

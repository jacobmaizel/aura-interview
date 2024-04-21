#!/bin/bash

# Wait for the database to be ready
until pg_isready -U postgres -h db -p 5432; do
  echo "Waiting for the database to be ready..."
  sleep 5
done

export PGPASSWORD=postgres
# Run migrations
echo "Running migrations..."
psql -U postgres -h db -d postgres -a -f /app/migrations/init.sql


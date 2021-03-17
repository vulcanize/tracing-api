#!/bin/sh
sleep 1
# Runs the db migrations and starts the service
set +x

# Construct the connection string for postgres
VDB_PG_CONNECT=postgresql://$CACHE_DATABASE_USER:$CACHE_DATABASE_PASSWORD@$CACHE_DATABASE_HOSTNAME:$CACHE_DATABASE_PORT/$CACHE_DATABASE_NAME?sslmode=disable

until PGPASSWORD=$CACHE_DATABASE_PASSWORD psql -h "$CACHE_DATABASE_HOSTNAME" -U "$CACHE_DATABASE_USER" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 10
done

>&2 echo "Postgres is up - executing command"

# Run the DB migrations
echo "Connecting with: $VDB_PG_CONNECT"
echo "Running database migrations"
goose -table goose_db_version_trace -dir migrations postgres "$VDB_PG_CONNECT" up

# If the db migrations ran without err
if [[ $? -ne 0 ]]; then
    echo "Could not run migrations. Are the database details correct?"
    exit 1
fi

echo "Running the Tracing-API process"
exec tracer "$@"
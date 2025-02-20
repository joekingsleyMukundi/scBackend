#!/bin/sh

set -e

echo "db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "start app"
exec "$@"
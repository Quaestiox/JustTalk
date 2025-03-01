#!/bin/sh

set -e

echo "run db migration"
/usr/app/migrate -path /usr/app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
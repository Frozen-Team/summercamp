#!/usr/bin/env bash
export PGPASSWORD=q1w2e3r4
until psql -h "db" -U "postgres" -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
godep get
./migrate.sh "" -m test
./prepare_version.sh
go get ./...
go test ./tests/... || exit 1
bee pack || exit 2
mkdir -p /tmp/build/
mv summercamp.tar.gz /tmp/build || exit 4
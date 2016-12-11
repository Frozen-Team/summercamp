#!/usr/bin/env bash
export PGPASSWORD=q1w2e3r4
until psql -h "db" -U "postgres" -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
godep get
./prepare_version.sh
go get ./...
case "$SUMMERCAMP_CONTAINER_OBJECTIVE" in
 "test_and_build" )
    export PGPASSWORD=q1w2e3r4
    psql -U postgres -c "CREATE DATABASE summercamp_test" -h db
    ./migrate.sh "" -m test
    go test ./tests/... || exit 1
    bee pack || exit 2
    mkdir -p /tmp/build/
    mv summercamp.tar.gz /tmp/build || exit 4
 ;;
 "run")
    ./migrate.sh "" -m dev
    bee run || exit 1
 ;;
esac
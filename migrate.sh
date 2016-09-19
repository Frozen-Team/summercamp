#!/bin/bash

RED='\033[0;31m'
NC='\033[0m'

source database/migrate.conf
MODE="dev"

function error {
    >&2 echo -e $RED$1$NC
    exit
}
case $1 in
    rollback)
        SUBCMD="rollback"
        ;;
    reset)
        SUBCMD="reset"
        ;;
    refresh)
        SUBCMD="refresh"
        ;;
    *)
        # Checking if it's not a key
        if [[ $1 != \-* ]];
        then
          error "Unknown command '$1'"
        fi
esac
while [[ $# -gt 1 ]]
    do
    key="$1"

    case $key in
        -m|--mode)
        MODE="$2"
        shift # past argument
        ;;
    esac
    shift
done
case $MODE in
    dev)
        DSN=$dsn_dev
        ;;
    prod)
        DSN=$dsn_prod
        ;;
    test)
        DSN=$dsn_test
        ;;
    *)
        error "Unknown mode"
        ;;
esac
if [ -z "$GOPATH" ]; then
    error "GOPATH variable not set. Migration aborted."
fi

if [ ! -x $GOPATH/bin/bee ]; then
    echo -e "${RED}Bee not found in GOPATH.${NC}"
    echo -e "Installing bee to continue..."
    go get github.com/beego/bee
    if [ ! -x $GOPATH/bin/bee ]; then
        echo -e "${RED}Unable to install github.com/beego/bee${NC}"
        echo -e "You should install bee manually: go get github.com/beego/bee to run migration via bee${NC}"
        exit 1
    fi
fi

echo "Running bee migrate $SUBCMD -driver="$driver" -conn="$driver://$DSN""
bee migrate $SUBCMD -driver="$driver" -conn="$driver://$DSN"
#!/bin/sh

RED='\033[0;31m'
NC='\033[0m'

source database/migrate.conf

if [ -z "$GOPATH" ]; then
    echo "${RED}GOPATH variable not set. Migration aborted.${NC}"
    exit 1
fi

if [ ! -x $GOPATH/bin/bee ]; then
    echo -e "${RED}Bee not found in GOPATH.${NC}"
    echo -e "Installing bee to continue..."
    $(go get github.com/beego/bee)
    if [ ! -x $GOPATH/bin/bee ]; then
        echo -e "${RED}Unable to install github.com/beego/bee${NC}"
        echo -e "You should install bee manually: go get github.com/beego/bee to run migration via bee${NC}"
        exit 1
    fi
fi

bee migrate $1 -driver="$driver" -conn="$driver://$dsn"

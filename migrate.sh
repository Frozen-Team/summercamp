#!/bin/bash

RED='\033[0;31m'
NC='\033[0m'

source database/migrate.conf

databases=( summercamp summercamp_test )

containsDb () {
  local e
  for e in "${databases[@]}"; do
    if [ "$e" == "$1" ] ; then
     return 0
    fi
  done
  return 1
}

if ! containsDb "$1" ; then
    echo "invalid database chosen. Choose any of: [ ${databases[@]} ]"
    exit 1
fi

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

bee migrate $2 -driver="$driver" -conn="$driver://$dsn"

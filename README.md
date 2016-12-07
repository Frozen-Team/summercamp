[![Build Status](https://travis-ci.org/Frozen-Team/summercamp.svg?branch=mvp)](https://travis-ci.org/Frozen-Team/summercamp)
## Requirements

The following instructions are only applicable if the **$GOPATH** is set correctly and you have 
**$GOPATH/bin** in your **$PATH**

1. Install **bee** by running: `go get github.com/beego/bee`
2. Install **godep** by running: `go get github.com/tools/godep`
3. Run `godep get` to install missing external dependencies described in a Godeps/Godeps.json file

## Running application

Use command `bee run`

## Building application
To build application you should have `docker-compose` installed on your machine.
This commands will start all necessary containers and will built project 


    cd summercamp
    docker-compose up
    

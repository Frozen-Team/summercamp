[![Build Status](https://travis-ci.org/Frozen-Team/summercamp.svg?branch=mvp)](https://travis-ci.org/Frozen-Team/summercamp)
## Requirements

The following instructions are only applicable if the **$GOPATH** is set correctly and you have 
**$GOPATH/bin** in your **$PATH**

1. Install **bee** by running: `go get github.com/beego/bee`
2. Install **godep** by running: `go get github.com/tools/godep`
3. Run `godep get` to install missing external dependencies described in a Godeps/Godeps.json file

## Running application

Use command `bee run`

## Running and building application
To build application or run application you should have `docker-compose` installed on your machine.
 
### Starting application
These commands will start all necessary containers and will run the project

    cd summercamp_dev
    docker-compose up web-run
### Building application
These commands will start all necessary containers and will build the project 


    cd summercamp_dev
    docker-compose up web-build
    

#!/usr/bin/env bash

echo "package main
const version = \"`git --git-dir=./.git describe --tags --long`\"" > version.go
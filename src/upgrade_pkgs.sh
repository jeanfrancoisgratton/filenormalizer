#!/usr/bin/env sh

egrep "gotest|golang|github" go.mod|awk '{print "go get "$1}'|sh

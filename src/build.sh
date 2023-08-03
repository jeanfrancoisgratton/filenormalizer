#!/usr/bin/env sh

OUTPUT=/opt/bin
#MANPAGE=README.md

if [ "$#" -gt 0 ]; then
    OUTPUT=$1
fi

if [ "$#" -gt 1 ]; then
  MANPAGE=$2
fi

go build -o ${OUTPUT}/fnormalizer .

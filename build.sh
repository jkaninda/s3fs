#!/usr/bin/env bash
if [ $# -eq 0 ]
  then
    tag='latest'
  else
    tag=$1
fi

docker build  -t jkaninda/s3fs:$tag .

docker-compose up -d
#!/bin/sh

MONGO_HOST=dev1.local
MONGO_PORT=27017
LANG="de"
COUNTRY="Deutschland"


mongoexport --uri=mongodb://$MONGO_HOST:$MONGO_PORT \
  --collection=products \
  --query="{\"lang\":\"${LANG}\",\"countries\":\"${COUNTRY}\"}" \
  --out products.$LANG.json
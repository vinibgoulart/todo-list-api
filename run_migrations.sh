#!/bin/bash

# $1 -> up/down
# $2 -> environment

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

OPTIONS="-config=./dbconfig.yml -env $2"

set -ex

sql-migrate $1 $OPTIONS
sql-migrate status $OPTIONS
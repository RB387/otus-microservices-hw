#!/bin/bash

sleep ${SLEEP}
exec shmig -t postgresql -d ${POSTGRES_DB} -l ${POSTGRES_USER} -H ${POSTGRES_HOST} -p ${POSTGRES_PASSWORD} -m /migrations "$1" "$2"
#!/usr/bin/env sh
set -e



echo "=== Insert Test Data ==="

curl -X POST localhost:8080 -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

curl -X POST localhost:8080 -d \
'{"activity": {"description": "cross country skiing is horrible and cold", "time":"2021-12-09T16:56:12Z"}}'

curl -X POST localhost:8080 -d \
'{"activity": {"description": "sledding with nephew", "time":"2021-12-09T16:56:23Z"}}'



echo "=== Test Descriptions ==="

curl -X GET localhost:8080 -d '{"id": 0}' 2>/dev/null | grep -q 'christmas eve bike class'
[ $? == 0 ] && echo success
curl -X GET localhost:8080 -d '{"id": 1}' 2>/dev/null | grep -q 'cross country skiing'
[ $? == 0 ] && echo success
curl -X GET localhost:8080 -d '{"id": 2}' 2>/dev/null | grep -q 'sledding'
[ $? == 0 ] && echo success
# check=skip=SecretsUsedInArgOrEnv

# Grocery Inventory Demo
# Copyright ⓒ 2026 Athaariq Ardhiansyah
#
# This program comes with ABSOLUTELY NO WARRANTY and you are welcome to redistribute it under
# certain conditions. See LICENSE file.

FROM postgres:18-alpine

COPY schema.sql /docker-entrypoint-initdb.d/init.sql

ENV POSTGRES_DB=grocery_inventory \
	POSTGRES_PASSWORD=s8vEGfvyzWfM3g1G

RUN /usr/bin/env bash /usr/local/bin/docker-enforce-initdb.sh
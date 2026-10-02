# Grocery Inventory Demo
# Copyright ⓒ 2026 Athaariq Ardhiansyah
#
# This program comes with ABSOLUTELY NO WARRANTY and you are welcome to redistribute it under
# certain conditions. See LICENSE file.

FROM python:3-slim-trixie

RUN apt update && apt install -y --auto-remove --purge python3-psycopg && rm -rf /var/cache/apt

COPY data.csv /usr/local/share/data.csv
COPY load_data.py /usr/local/share/load_data.py

ENTRYPOINT [ "/usr/bin/python3" ]
CMD [ "/usr/local/share/load_data.py" ]
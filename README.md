# Grocery Inventory REST API for Business Intelligence assignment

Dataset source: https://www.kaggle.com/datasets/willianoliveiragibin/grocery-inventory

Required programs:

- [Docker](https://www.docker.com/get-started/)
- [docker-compose](https://docs.docker.com/compose/install/)

## How to use

1. Open terminal in this directory (or folder in Windows)
2. Enter: `docker-compose up`
3. Open this on browser: http://localhost:8000/best_sales

## Backend Endpoints

- http://localhost:8000/best_sales = Sort 10 highest products for re-stocking priority
- http://localhost:8000/worst_sales = Sort 10 lowest products for promotion planning
- http://localhost:8000/warehouses = Get all warehouse IDs and names for dropdown list
- http://localhost:8000/best_sales/(warehouse_id) = Sort 10 highest products on specific warehouse
- http://localhost:8000/worst_sales/(warehouse_id) = Sort 10 lowest products on specific warehouse

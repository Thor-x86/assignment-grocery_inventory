#!/usr/bin/python3

"""
Grocery Inventory Demo
Copyright ⓒ 2026 Athaariq Ardhiansyah

This program comes with ABSOLUTELY NO WARRANTY and you are welcome to redistribute it under certain
conditions. See LICENSE file.
"""

from dataclasses import dataclass
import csv
from datetime import datetime
from time import sleep
import psycopg

@dataclass
class WarehouseModel:
	id: int
	code: int

@dataclass
class ProductModel:
	supplier_id: int
	warehouse_id: int
	received_at: str # ISO 8601
	last_order_at: str # ISO 8601
	expired_at: str # ISO 8601
	name: str
	status: str
	unit_price_usd: float
	percentage: float
	stock_qty: int
	reorder_level: int
	reorder_qty: int
	sales_volume: int
	inventory_turnover_ratio: int

# Key is ID
supplier_names: dict[int, str] = {}

# Key is name
warehouses: dict[str, WarehouseModel] = {}

# Key is ID
products: dict[int, ProductModel] = {}

last_warehouse_id = 0

# Column order:
# [0] Product_Name, [1] Category, [2] Supplier_Name, [3] Warehouse_Location, [4] Status,
# [5] Product_ID, [6] Supplier_ID, [7] Date_Received, [8] Last_Order_Date, [9] Expiration_Date,
# [10] Stock_Quantity, [11] Reorder_Level, [12] Reorder_Quantity, [13] Unit_Price,
# [14] Sales_Volume, [15] Inventory_Turnover_Rate, [16] percentage

with open("/usr/local/share/data.csv") as data_file:
	csv_reader = csv.reader(data_file)
	csv_reader.__next__() # Skip header
	for each_line in csv_reader:
		warehouse_code_name = each_line[3]
		[warehouse_code_str, warehouse_name] = warehouse_code_name.split(" ", 1)
		if warehouse_name in warehouses.keys():
			warehouse_id = warehouses[warehouse_name].id
		else:
			last_warehouse_id += 1
			warehouse_id = last_warehouse_id
			warehouse_code = int(warehouse_code_str)
			new_warehouse = WarehouseModel(warehouse_id, warehouse_code)
			warehouses[warehouse_name] = new_warehouse

		supplier_id = int(each_line[6].replace('-', ''))
		if supplier_id in supplier_names.keys():
			supplier_name = supplier_names[supplier_id]
		else:
			supplier_name = each_line[2]
			supplier_names[supplier_id] = supplier_name

		product_id = int(each_line[5].replace('-', ''))
		received_at = datetime.strptime(each_line[7], "%m/%d/%y").isoformat()
		last_order_at = datetime.strptime(each_line[8], "%m/%d/%y").isoformat()
		expired_at = datetime.strptime(each_line[9], "%m/%d/%y").isoformat()
		name = each_line[0]
		status = each_line[4].lower()
		unit_price_usd = float(each_line[13].strip('$'))
		percentage = float(each_line[16].strip('%'))
		stock_qty = int(each_line[10])
		reorder_level = int(each_line[11])
		reorder_qty = int(each_line[12])
		sales_volume = int(each_line[14])
		inventory_turnover_ratio = int(each_line[15])

		new_product = ProductModel(supplier_id, warehouse_id, received_at, last_order_at,
		                           expired_at, name, status, unit_price_usd, percentage, stock_qty,
								   reorder_level, reorder_qty, sales_volume,
								   inventory_turnover_ratio)

		products[product_id] = new_product

attempt_count = 0

while True:
	try:
		conn = psycopg.connect(
			dbname="grocery_inventory",
			user="postgres",
			password="s8vEGfvyzWfM3g1G",
			host="database.grocery_inventory.internal",
			port="5432"
		)
		break
	except psycopg.OperationalError:
		if attempt_count > 10:
			raise TimeoutError("Failed to connect PostgreSQL in 10 seconds")
		sleep(1)
		attempt_count += 1

with conn.cursor() as cur:
	cur.execute("TRUNCATE public.warehouse CASCADE")
	warehouse_query = "INSERT INTO public.warehouse (id, name, code) VALUES (%s, %s, %s)"
	warehouse_data: list[tuple[int, str, int]] = []
	for each_name in warehouses.keys():
		each_warehouse = warehouses[each_name]
		warehouse_data.append((each_warehouse.id, each_name, each_warehouse.code))
	cur.executemany(warehouse_query, warehouse_data)

	cur.execute("TRUNCATE public.supplier CASCADE")
	supplier_query = "INSERT INTO public.supplier (id, name) VALUES (%s, %s)"
	supplier_data: list[tuple[int, str]] = []
	for each_id in supplier_names.keys():
		each_name = supplier_names[each_id]
		supplier_data.append((each_id, each_name))
	cur.executemany(supplier_query, supplier_data)

	cur.execute("TRUNCATE public.product CASCADE")
	product_query = "INSERT INTO public.product (id, supplier_id, warehouse_id, received_at, last_order_at, expired_at, name, status, unit_price_usd, percentage, stock_qty, reorder_level, reorder_qty, sales_volume, inventory_turnover_ratio) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"
	product_data: list[tuple[int, int, int, str, str, str, str, str, float, float, int, int, int, int, int]] = []
	for each_id in products.keys():
		each_product = products[each_id]
		product_data.append((
			each_id,
			each_product.supplier_id,
			each_product.warehouse_id,
			each_product.received_at,
			each_product.last_order_at,
			each_product.expired_at,
			each_product.name,
			each_product.status,
			each_product.unit_price_usd,
			each_product.percentage,
			each_product.stock_qty,
			each_product.reorder_level,
			each_product.reorder_qty,
			each_product.sales_volume,
			each_product.inventory_turnover_ratio
		))
	cur.executemany(product_query, product_data)

	conn.commit()

conn.close()
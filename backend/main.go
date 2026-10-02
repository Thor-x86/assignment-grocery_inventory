package main

import (
	"database/sql"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const connStr = "postgres://app:hWPyWuUgmvfTzPUJ@database.grocery_inventory.internal:5432/grocery_inventory?sslmode=disable"

type SalesEntry struct {
	ID                     uint64    `json:"id"`
	Supplier               string    `json:"supplier"`
	Warehouse              string    `json:"warehouse"`
	WarehouseCode          uint32    `json:"warehouse_code"`
	ReceivedAt             time.Time `json:"received_at"`
	LastOrderAt            time.Time `json:"last_order_at"`
	ExpiredAt              time.Time `json:"expired_at"`
	Name                   string    `json:"name"`
	Status                 string    `json:"status"`
	UnitPriceUSD           float32   `json:"unit_price_usd"`
	Percentage             float32   `json:"percentage"`
	StockQty               uint16    `json:"stock_qty"`
	ReorderLevel           uint16    `json:"reorder_level"`
	ReorderQty             uint16    `json:"reorder_qty"`
	SalesVolume            uint16    `json:"sales_volume"`
	InventoryTurnoverRatio uint16    `json:"inventory_turnover_ratio"`
}

type WarehouseEntry struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func bestSales(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.AbortWithStatusJSON(503, gin.H{"error": "PostgreSQL not ready", "debug": err.Error()})
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT p.id, p.name, s.name, w.name, w.code, p.received_at, p.last_order_at, p.expired_at, p.status, p.unit_price_usd, p.percentage, p.stock_qty, p.reorder_level, p.reorder_qty, p.sales_volume, p.inventory_turnover_ratio FROM product p INNER JOIN warehouse w ON p.warehouse_id = w.id INNER JOIN supplier s ON p.supplier_id = s.id WHERE p.status = 'active' ORDER BY p.reorder_level DESC LIMIT 10")
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": err.Error()})
		return
	}

	defer rows.Close()

	if rows.Err() != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": rows.Err().Error()})
		return
	}

	entries := []SalesEntry{}
	for rows.Next() {
		var entry SalesEntry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Supplier, &entry.Warehouse, &entry.WarehouseCode, &entry.ReceivedAt, &entry.LastOrderAt, &entry.ExpiredAt, &entry.Status, &entry.UnitPriceUSD, &entry.Percentage, &entry.StockQty, &entry.ReorderLevel, &entry.ReorderQty, &entry.SalesVolume, &entry.InventoryTurnoverRatio)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the entry", "debug": err.Error()})
			return
		}
		entries = append(entries, entry)
	}

	c.JSON(200, entries)
}

func worstSales(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.AbortWithStatusJSON(503, gin.H{"error": "PostgreSQL not ready"})
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT p.id, p.name, s.name, w.name, w.code, p.received_at, p.last_order_at, p.expired_at, p.status, p.unit_price_usd, p.percentage, p.stock_qty, p.reorder_level, p.reorder_qty, p.sales_volume, p.inventory_turnover_ratio FROM product p INNER JOIN warehouse w ON p.warehouse_id = w.id INNER JOIN supplier s ON p.supplier_id = s.id WHERE p.status = 'active' ORDER BY p.reorder_level ASC LIMIT 10")
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": err.Error()})
		return
	}

	defer rows.Close()

	if rows.Err() != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": rows.Err().Error()})
		return
	}

	entries := []SalesEntry{}
	for rows.Next() {
		var entry SalesEntry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Supplier, &entry.Warehouse, &entry.WarehouseCode, &entry.ReceivedAt, &entry.LastOrderAt, &entry.ExpiredAt, &entry.Status, &entry.UnitPriceUSD, &entry.Percentage, &entry.StockQty, &entry.ReorderLevel, &entry.ReorderQty, &entry.SalesVolume, &entry.InventoryTurnoverRatio)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the entry", "debug": err.Error()})
			return
		}
		entries = append(entries, entry)
	}

	c.JSON(200, entries)
}

func warehouses(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.AbortWithStatusJSON(503, gin.H{"error": "PostgreSQL not ready"})
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM warehouse")
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": err.Error()})
		return
	}

	defer rows.Close()

	if rows.Err() != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": rows.Err().Error()})
		return
	}

	entries := []WarehouseEntry{}
	for rows.Next() {
		var entry WarehouseEntry
		err := rows.Scan(&entry.ID, &entry.Name)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the entry", "debug": err.Error()})
			return
		}
		entries = append(entries, entry)
	}

	c.JSON(200, entries)
}

func bestSalesByWarehouse(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.AbortWithStatusJSON(503, gin.H{"error": "PostgreSQL not ready"})
		return
	}

	defer db.Close()

	warehouse_id := c.Param("warehouse_id")

	rows, err := db.Query("SELECT p.id, p.name, s.name, w.name, w.code, p.received_at, p.last_order_at, p.expired_at, p.status, p.unit_price_usd, p.percentage, p.stock_qty, p.reorder_level, p.reorder_qty, p.sales_volume, p.inventory_turnover_ratio FROM product p INNER JOIN warehouse w ON p.warehouse_id = w.id INNER JOIN supplier s ON p.supplier_id = s.id WHERE p.status = 'active' AND p.warehouse_id = $1 ORDER BY p.reorder_level DESC LIMIT 10", warehouse_id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": err.Error()})
		return
	}

	defer rows.Close()

	if rows.Err() != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": rows.Err().Error()})
		return
	}

	entries := []SalesEntry{}
	for rows.Next() {
		var entry SalesEntry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Supplier, &entry.Warehouse, &entry.WarehouseCode, &entry.ReceivedAt, &entry.LastOrderAt, &entry.ExpiredAt, &entry.Status, &entry.UnitPriceUSD, &entry.Percentage, &entry.StockQty, &entry.ReorderLevel, &entry.ReorderQty, &entry.SalesVolume, &entry.InventoryTurnoverRatio)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the entry", "debug": err.Error()})
			return
		}
		entries = append(entries, entry)
	}

	c.JSON(200, entries)
}

func worstSalesByWarehouse(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.AbortWithStatusJSON(503, gin.H{"error": "PostgreSQL not ready"})
		return
	}

	defer db.Close()

	warehouse_id := c.Param("warehouse_id")

	rows, err := db.Query("SELECT p.id, p.name, s.name, w.name, w.code, p.received_at, p.last_order_at, p.expired_at, p.status, p.unit_price_usd, p.percentage, p.stock_qty, p.reorder_level, p.reorder_qty, p.sales_volume, p.inventory_turnover_ratio FROM product p INNER JOIN warehouse w ON p.warehouse_id = w.id INNER JOIN supplier s ON p.supplier_id = s.id WHERE p.status = 'active' AND p.warehouse_id = $1 ORDER BY p.reorder_level ASC LIMIT 10", warehouse_id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": err.Error()})
		return
	}

	defer rows.Close()

	if rows.Err() != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the data", "debug": rows.Err().Error()})
		return
	}

	entries := []SalesEntry{}
	for rows.Next() {
		var entry SalesEntry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Supplier, &entry.Warehouse, &entry.WarehouseCode, &entry.ReceivedAt, &entry.LastOrderAt, &entry.ExpiredAt, &entry.Status, &entry.UnitPriceUSD, &entry.Percentage, &entry.StockQty, &entry.ReorderLevel, &entry.ReorderQty, &entry.SalesVolume, &entry.InventoryTurnoverRatio)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to get the entry", "debug": err.Error()})
			return
		}
		entries = append(entries, entry)
	}

	c.JSON(200, entries)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/best_sales", bestSales)
	r.GET("/worst_sales", worstSales)
	r.GET("/warehouses", warehouses)
	r.GET("/best_sales/:warehouse_id", bestSalesByWarehouse)
	r.GET("/worst_sales/:warehouse_id", worstSalesByWarehouse)

	r.Run(":8000")
}

package main

import (
	"log"
	"time"

	humanize "github.com/dustin/go-humanize"
	pp "github.com/k0kubun/pp/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CustomModel overrides gorm.Model
type CustomModel struct {
	ID        string         `gorm:"primary_key;auto_increment:false"`
	CreatedAt int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Unix timestamp support for DeletedAt
	// https://github.com/go-gorm/gorm/issues/6404
	//DeletedAt int64  `gorm:"autoDeleteTime:milli"`
}

// Product is the db model we experiment with below
type Product struct {
	CustomModel // drop-in convention replacement for gorm.Model
	Code        string
	Price       uint
}

func timeTest() {
	now := time.Now().UnixMilli() // returns int64
	// signed vs unisghed
	type model struct {
		x int64
		y uint64
	}
	m := model{
		x: now,
		y: uint64(now),
	}
	log.Println(humanize.Bytes(2141741))
	//log.Println(m.x, strconv.FormatUint(m.y, 10))
	pp.Println(m)
}

// example below borrowed from https://gorm.io/docs/index.html#Quick-Start
func main() {

	timeTest()

	log.Println("opening database")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	log.Println("migrate product schema")
	db.AutoMigrate(&Product{})

	// Create
	tx := db.Create(
		&Product{
			CustomModel: CustomModel{ID: "p_123"},
			Code:        "D42",
			Price:       100,
		},
	)
	log.Printf("rows created %d", tx.RowsAffected)

	// Read
	var product Product
	db.First(&product, "id = ?", "p_123") // find product with primary key
	pp.Printf("find by product by id: %v\n", product)
	db.First(&product, "code = ?", "D42") // find product with code D42
	pp.Printf("find product by code: %v\n", product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 188)
	pp.Printf("update price: %v\n", product)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	pp.Printf("update multiple fields: %v\n", product)

	// Delete - delete product
	db.Delete(&product, 1)
	pp.Printf("deleted %v\n", product)
}

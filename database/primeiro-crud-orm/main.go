package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int       `gorm:"primaryKey"`
	Name     string    `gorm:"not null"`
	Products []Product `gorm:"many2many:product_categories"`
}

type Product struct {
	ID         int        `gorm:"primaryKey"`
	Name       string     `gorm:"not null"`
	Price      float64    `gorm:"not null"`
	Categories []Category `gorm:"many2many:product_categories"`
	// SerialNumber SerialNumber
	gorm.Model
}

// type SerialNumber struct {
// 	ID        int    `gorm:"primaryKey"`
// 	Number    string `gorm:"not null"`
// 	ProductID int
// }

// type ProductORM struct {
// 	ID    int     `gorm:"primaryKey"`
// 	Name  string  `gorm:"not null"`
// 	Price float64 `gorm:"not null"`
// 	gorm.Model
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronico"}
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Product 1",
		Price:      100.0,
		Categories: []Category{category, category2},
	})

	var products []Product
	db.Preload("Categories").Find(&products)
	for _, product := range products {
		println("Product selected:", product.ID, product.Name, product.Price)
		for _, category := range product.Categories {
			println("Category selected:", category.ID, category.Name)
		}
	}

	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// })

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println("Product selected:", product.ID, product.Name, product.Price, product.CategoryId, product.Category, product.SerialNumber.Number)
	// }

	// var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	fmt.Println("Category selected:", category.ID, category.Name)
	// 	for _, product := range category.Products {
	// 		fmt.Println("Product selected:", product.ID, product.Name, product.Price, product.SerialNumber.Number)
	// 	}
	// }

	// db.Create(&ProductORM{
	// 	Name:  "Product 1",
	// 	Price: 100.0,
	// })

	// products := []ProductORM{
	// 	{Name: "Product 2", Price: 200.0},
	// 	{Name: "Product 3", Price: 300.0},
	// 	{Name: "Product 4", Price: 400.0},
	// }

	// db.Create(&products)

	// var product ProductORM
	// db.First(&product, 1)
	// println("Product selected:", product.ID, product.Name, product.Price)
	// db.First(&product, "name = ?", "Product 2")
	// println("Product selected:", product.ID, product.Name, product.Price)

	// var products []ProductORM
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	println("Product selected:", product.ID, product.Name, product.Price)
	// }

	// var product []ProductORM
	// db.Where("name LIKE ?", "%1").Find(&product)
	// for _, product := range product {
	// 	fmt.Println(product)
	// }

	// var p ProductORM
	// db.First(&p)
	// println("Product selected:", p.ID, p.Name, p.Price)
	// p.Name = "Product 1 updated"
	// db.Save(&p)
	// println("Product updated:", p.ID, p.Name, p.Price)

	// var p2 ProductORM
	// db.First(&p2, 1)
	// println("Product selected:", p2.ID, p2.Name, p2.Price)

	// db.Delete(&p2)
	// println("Product deleted:", p2.ID)

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	c.Name = "Cozinha updated"
	tx.Debug().Save(&c)
	tx.Commit()
	println("Category updated:", c.ID, c.Name)
}

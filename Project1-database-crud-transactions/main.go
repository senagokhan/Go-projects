package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ProductInsert struct {
	productID          int64
	productName        string
	productDescription string
	price              float64
	stock              int16
	category           string
	suppliers          string
	weight             float64
	dimensions         string
	color              string
}

type ProductUpdate struct {
	productID          int64
	productName        string
	productDescription string
	price              float64
	stock              int16
	category           string
	suppliers          string
	weight             float64
	dimensions         string
	color              string
}

var db *sql.DB

func main() {

	var err error
	db, err = connectMysql()
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	var choice int
	fmt.Println("1 : Insert Product")
	fmt.Println("2 : Delete Product")
	fmt.Println("3 : Update Product")
	fmt.Println("4 : Show All Product")
	fmt.Println("5 : Show Product Details")
	fmt.Printf("Please make your choice : ")
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		createTableProducts()
		data := ProductInsert{}

		fmt.Print("Product Name: ")
		fmt.Scanln(&data.productName)
		fmt.Print("Product Price: ")
		fmt.Scanln(&data.price)
		fmt.Print("Product Description: ")
		fmt.Scanln(&data.productDescription)
		fmt.Print("Product Stock: ")
		fmt.Scanln(&data.stock)
		fmt.Print("Product Category: ")
		fmt.Scanln(&data.category)
		fmt.Print("Product Suppliers: ")
		fmt.Scanln(&data.suppliers)
		fmt.Print("Product Weight: ")
		fmt.Scanln(&data.weight)
		fmt.Print("Product Dimensions: ")
		fmt.Scanln(&data.dimensions)
		fmt.Print("Product Color: ")
		fmt.Scanln(&data.color)

		record := AddProducts(data)
		if record != 0 {
			fmt.Println("Successfully added product")
		}

	case 2:
		var productIdToDelete int64
		fmt.Print("Enter Product ID to delete: ")
		fmt.Scanln(&productIdToDelete)

		DeleteProducts(productIdToDelete)

	case 3:
		var productIdToUpdate int64
		fmt.Print("Enter Product ID to update: ")
		fmt.Scanln(&productIdToUpdate)

		updateData := ProductUpdate{}

		fmt.Print("Updated Product Name: ")
		fmt.Scanln(&updateData.productName)
		fmt.Print("Updated Product Description: ")
		fmt.Scanln(&updateData.productDescription)
		fmt.Print("Updated Product Price: ")
		fmt.Scanln(&updateData.price)
		fmt.Print("Updated Product Stock: ")
		fmt.Scanln(&updateData.stock)
		fmt.Print("Updated Product Category: ")
		fmt.Scanln(&updateData.category)
		fmt.Print("Updated Product Suppliers: ")
		fmt.Scanln(&updateData.suppliers)
		fmt.Print("Updated Product Weight: ")
		fmt.Scanln(&updateData.weight)
		fmt.Print("Updated Product Dimensions: ")
		fmt.Scanln(&updateData.dimensions)
		fmt.Print("Updated Product Color: ")
		fmt.Scanln(&updateData.color)
		UpdateProducts(productIdToUpdate, updateData)

	case 4:
		ShowAllProducts()

	case 5:
		var productIdToGet int64
		fmt.Print("Enter Product ID to get details: ")
		fmt.Scanln(&productIdToGet)
		GetProductDetailsById(productIdToGet)

	default:
		fmt.Println("Invalid choice")
	}
}

func connectMysql() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "123456"
	dbName := "product_catalogue"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	fmt.Println("Successfully connected to MySQL!")
	return db, nil
}

func createTableProducts() {
	tableProduct := `
    CREATE TABLE IF NOT EXISTS products (
        productID INT AUTO_INCREMENT PRIMARY KEY,
        productName VARCHAR(50) NOT NULL,
        productDescription TEXT,
        price DECIMAL(10,2) NOT NULL,
        stock INT NOT NULL,
        category VARCHAR(50) NOT NULL,
        suppliers VARCHAR(50) NOT NULL,
        weight FLOAT NOT NULL,
        dimensions VARCHAR(50) NOT NULL,
        color VARCHAR(50) NOT NULL
    );`

	_, err := db.Exec(tableProduct)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Products table created successfully.")
}

func AddProducts(data ProductInsert) int64 {
	result, err := db.Exec(`
    INSERT INTO Products (productName, productDescription, price, stock, category, suppliers, weight, dimensions, color)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		data.productName, data.productDescription, data.price, data.stock, data.category, data.suppliers, data.weight, data.dimensions, data.color)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func ShowAllProducts() {
	rows, err := db.Query(`SELECT productID, productName, productDescription, price , stock, category, suppliers, weight, dimensions, color FROM products`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var productID int64
		var productName, productDescription, category, suppliers, dimensions, color string
		var price, weight float64
		var stock int16

		err := rows.Scan(&productID, &productName, &productDescription, &price, &stock, &category, &suppliers, &weight, &dimensions, &color)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Product Id: %d, Product Name: %s, Description: %s, Price: %.2f, Stock: %d\n, Category: %s, Supplier: %s, Weight: %.2f, Dimensions: %s, Color: %s\n",
			productID, productName, productDescription, price, stock, category, suppliers, weight, dimensions, color)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteProducts(productID int64) {
	_, err := db.Exec(`DELETE FROM products WHERE productID = ?`, productID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted product with ID:", productID)
}

func UpdateProducts(productID int64, data ProductUpdate) {
	_, err := db.Exec(`
    UPDATE Products SET productName = ?, productDescription = ?, price = ?, stock = ?, category = ?, suppliers = ?, weight = ?, dimensions = ?, color = ?
    WHERE ProductId = ?`,
		data.productName, data.productDescription, data.price, data.stock, data.category, data.suppliers, data.weight, data.dimensions, data.color, productID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully updated product with ID:", productID)
}

func GetProductDetailsById(productID int64) {
	var productName, productDescription, category, suppliers, dimensions, color string
	var price, weight float64
	var stock int16

	err := db.QueryRow(`SELECT productName, productDescription, price, stock, category, suppliers, weight, dimensions, color FROM products WHERE productID = ?`, productID).Scan(
		&productName, &productDescription, &price, &stock, &category, &suppliers, &weight, &dimensions, &color)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No product found with the given ID.")
		} else {
			log.Fatal(err)
		}
		return
	}
	fmt.Println("Product Details:")
	fmt.Println("Product Name: ", productName)
	fmt.Println("Product Description: ", productDescription)
	fmt.Println("Price: ", price)
	fmt.Println("Stock: ", stock)
	fmt.Println("Category: ", category)
	fmt.Println("Suppliers: ", suppliers)
	fmt.Println("Weight: ", weight)
	fmt.Println("Dimensions: ", dimensions)
	fmt.Println("Color: ", color)
}

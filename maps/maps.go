package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)

	// Add values
	product["Macbook"] = 50000
	product["iPad"] = 40000
	product["iPhone"] = 45000
	product["iPods"] = 10000
	fmt.Println("Add product =", product)

	// Update the value
	product["iPhone"] = 42000
	fmt.Println("Update product =", product)

	// Delete value
	delete(product, "iPad")
	fmt.Println("Delete product =", product)

	// Access map variable
	iPhone := product["iPhone"]
	fmt.Println("iPhone =", iPhone)

	// Create map variable and define values
	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("courseName =", courseName)
}

package main

// SearchProduct ...
func SearchProduct(search string, products []Product) []Product {
	return filter(products, shouldAdd, search)
}

func filter(products []Product, shouldAdd func(Product, string) bool, search string) []Product {
	filteredProducts := make([]Product, 0)

	for _, product := range products {
		if shouldAdd(product, search) {
			filteredProducts = append(filteredProducts, product)
		}
	}

	return filteredProducts
}

func shouldAdd(product Product, search string) bool {
	if product.Name == "VESTIDO TRANSPASSE BOW" {
		return true
	}
	return false
}

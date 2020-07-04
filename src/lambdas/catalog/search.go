package main

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// SearchProduct ...
func SearchProduct(search string, products []Product) []Product {
	return filter(products, shouldAdd, strings.ToLower(search))
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
	if findByName(product.Name, search) || findByID(product.Style, search) {
		return true
	}
	return false
}

func findByName(name string, search string) bool {
	name = strings.ToLower(name)
	name = removeDiacrits(name)
	return strings.Contains(name, search)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func removeDiacrits(value string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, value)
	return result
}

func findByID(id string, search string) bool {
	stylesParams := strings.Split(search, "_")

	return id == stylesParams[0]
}

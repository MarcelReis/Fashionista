package main

import (
	"strconv"
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

	stringID := strings.Split(search, "_")
	ID, err := strconv.Atoi(stringID[0])
	if err == nil {
		for _, product := range products {
			if findByID(product.Style, stringID[0]) {
				filteredProducts = append(filteredProducts, product)
			}
		}

	} else {
		for _, product := range products {
			if findByName(product.Name, search) {
				filteredProducts = append(filteredProducts, product)
			}
		}
	}

	return filteredProducts
}

func findByID(id string, search string) bool {
	return id == search
}

func findByName(name string, search string) bool {
	name = strings.ToLower(name)
	name = removeDiacrits(name)
	return strings.Contains(name, search)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func removeDiacrits(value string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, value)

	return result
}

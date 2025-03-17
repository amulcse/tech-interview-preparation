package main

import (
	"fmt"
	"strings"
)

// arrangeBreadcrumbs converts a list of S3 paths into a hierarchical tree structure.
func arrangeBreadcrumbs(paths []string) map[string]interface{} {
	tree := make(map[string]interface{})

	for _, path := range paths {
		parts := strings.Split(path, "/")
		node := tree

		for _, part := range parts {
			if _, exists := node[part]; !exists {
				node[part] = make(map[string]interface{})
			}
			node = node[part].(map[string]interface{})
		}
	}

	return tree
}

// displayBreadcrumbs prints the hierarchical tree structure.
func displayBreadcrumbs(node map[string]interface{}, level int) {
	for key, value := range node {
		fmt.Println(strings.Repeat("*", level+1), key)
		displayBreadcrumbs(value.(map[string]interface{}), level+1)
	}
}

func main() {
	paths := []string{
		"Electronics/Home/TV",
		"Furniture/Table1",
		"Furniture/Bedroom/Bed",
		"Electronics/Gadget/Table",
		"Electronics/Home/Bulb",
		"Electronics/Gadget/Mobile",
		"Gifts/Festival/Holi/Discounts",
	}

	tree := arrangeBreadcrumbs(paths)
	displayBreadcrumbs(tree, 0)
}

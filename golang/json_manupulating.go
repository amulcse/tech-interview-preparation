package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	input := []byte(`{
		"name" : "Amul",
		"email" : "amul@gmail.com",
		"favorite_colors" : ["blue","green"],
		"schools" : {
			"primary" : "transwad",
			"secondary" : "visnagar",
			"colleges" : {
				"b" : "GEC",
				"m"	: "Windsor"
			}
		}	
	}`)

	update := []byte(`{
		"name" : "Patel",
		"email" : "patel@gmail.com",
		"schools" : {
			"colleges" : {
				"b" : "GEC",
				"m"	: "Windsor University"
			}
		}	
	}`)

	var data map[string]interface{}

	err := json.Unmarshal(input, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	var updateData map[string]interface{}

	err = json.Unmarshal(update, &updateData)

	_ = countKeys(data, updateData)
	fmt.Println("Total keys including nested ones:", data)
}

// Recursive function to count all keys
func countKeys(m map[string]interface{}, update map[string]interface{}) int {
	count := 0
	for k, v := range m {
		// If value is a nested object, count its keys too
		if nestedMap, ok := v.(map[string]interface{}); ok {
			var nestedUpdate map[string]interface{}
			if nestedUpdate, ok = update[k].(map[string]interface{}); !ok {
				nestedUpdate = nil
			}
			count += countKeys(nestedMap, nestedUpdate)
		} else if newValue, ok := update[k]; ok {
			m[k] = newValue
			count++
		} else {
			count++
		}
	}
	return count
}

package main

import (
	"fmt"
	"reflect"
)
//memoization
 var mt = make(map[string]map[string]interface{})
// helper function searches for a key recursively in a nested map
func helper(key string, m map[string]interface{}) (map[string]interface{}, error) {
	// Check if the key exists in the current map
	if _, ok := m[key]; ok {
		return m, nil // Key found, return the map
	}

	// Iterate over the values in the map
	for _, val := range m {
		// Use reflection to determine the type of the value
		a := reflect.ValueOf(val)

		switch a.Kind() {

		// If the value is a nested map, call the helper function recursively
		case reflect.Map:
			nestedMap := a.Interface().(map[string]interface{})
			if found, err := helper(key, nestedMap); err == nil {
				return found, nil // Key found in the nested map, return the map
			}

		// If the value is a slice, iterate over its elements
		case reflect.Slice:
			nestedSlice := a.Interface().([]interface{})
			for _, value := range nestedSlice {
				a := reflect.ValueOf(value)
				// If the element is a map, call the helper function recursively
				if a.Kind() == reflect.Map {
					nestedMapInSlice := a.Interface().(map[string]interface{})
					if found, err := helper(key, nestedMapInSlice); err == nil {
						return found, nil // Key found in the nested map within the slice, return the map
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("key not found") // Key not found in the map
}

// update function updates the value of a key in the map
func update(key string, m map[string]interface{}, value interface{}) {
	// Call helper function to find the key in the map

	if mt[key]!=nil{
		nestmap:= mt[key]
		nestmap[key]=value
		fmt.Println("the value has been updated") 
	}else {
	if foundVal, err := helper(key, m); err != nil {
		fmt.Println(err) // Key not found, print error message
	} else {
		mt[key]=foundVal
		foundVal[key] = value // Update the value of the key
		fmt.Println("the value has been updated")
	}
}
}



func main() {

	// Example map containing various types of data
	var m = map[string]interface{}{
		"Name":          "rolex india pvt ltd",
		"landmark":      "India gate",
		"city":          "lucknow",
		"pincode":       226029,
		"state":         "Uttar Pradesh",
		"Address": []interface{}{
			map[string]interface{}{
				"street":     "rani jhasi marg",
				"land mark":  "Jhandewalan metro station",
				"city":       "Noida",
				"pincode":    226028,
				"state":      "uttar pradesh",
			},
			map[string]interface{}{
				"street":     "rani jhasi marg",
				"land mark":  "Jhandewalan metro station",
				"city":       "Noida",
				"pincode":    226028,
				"state":      "uttar pradesh",
			},
		},
		"Investors":     map[string]interface{}{"Name": "one"},
		"revenue":       "100 million$",
		"no_of_employee": 630,
		"str_text":      []interface{}{"one", "two"},
		"int_text":      []interface{}{1, 2, 3},
	}

	var value interface{}
	value = "pranav"

	// Update the value of the "pincode" key in the map
	update("pincode", m, value)
	fmt.Println(m)

	
	fmt.Println(m)
}


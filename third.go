package main

import (
	"fmt"
	"reflect"
)

// PopulateStruct function populates fields of a struct using data from a map.
func PopulateStruct(data map[string]interface{}, result interface{}) {
	// Get the reflection value of the result interface and navigate to its underlying struct
	resultValue := reflect.ValueOf(result).Elem()

	// Iterate over each key-value pair in the data map
	for key, value := range data {
		// Get the field in the struct corresponding to the key
		field := resultValue.FieldByName(key)

		// Check if the field is valid (exists in the struct)
		if field.IsValid() {
			// If the field is a struct itself, recursively populate its fields
			if field.Kind() == reflect.Struct {
				// Check if the value is a map, indicating nested struct data
				if nestedMap, ok := value.(map[string]interface{}); ok {
					// Create a new instance of the nested struct type
					nestedStruct := reflect.New(field.Type()).Interface()
					// Recursively populate the nested struct fields
					PopulateStruct(nestedMap, nestedStruct)
					// Set the populated nested struct back to the field
					field.Set(reflect.ValueOf(nestedStruct).Elem())
				}
			} else {
				// For non-struct fields, set the value directly
				field.Set(reflect.ValueOf(value))
			}
		}
	}
}

// Define a Person struct
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Define an Address struct
type Address struct {
	City  string
	State string
}

func main() {
	// Sample data to populate the Person struct
	data := map[string]interface{}{
		"Name":    "John",
		"Age":     30,
		"pincode": 226029, // This field won't be used as it doesn't exist in the struct
		"Address": map[string]interface{}{
			"City":  "New York",
			"State": "NY",
		},
	}

	// Create a pointer to a Person struct
	var personPtr *Person = &Person{}
	// Populate the Person struct fields using the data
	PopulateStruct(data, personPtr)

	// Print the populated Person struct
	fmt.Printf("%+v\n", *personPtr)
}

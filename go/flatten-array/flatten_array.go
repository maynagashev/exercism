package flatten

import "fmt"

// Flatten flattens nested array
func Flatten(input interface{}) (flattened []interface{}) {
	fmt.Printf("Input: %+v\n", input)
	for _, v := range input.([]interface{}) {
		flattened = append(flattened, v)
	}
	//fmt.Printf("Flattened: %+v\n", flattened)
	return flattened
}

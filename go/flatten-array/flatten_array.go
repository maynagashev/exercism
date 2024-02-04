package flatten

func Flatten(input interface{}) []interface{} {
	var result = make([]interface{}, 0)
	//fmt.Printf("Input: %+v\n", input)
	for _, value := range input.([]interface{}) {
		switch v := value.(type) {
		case []interface{}:
			// If the element is a slice, recursively flatten it and append the results.
			result = append(result, Flatten(v)...)
		case nil:
			// Ignore nil values.
			continue
		default:
			// Append non-nil, non-slice values directly to the result.
			result = append(result, v)
		}
	}

	return result
}
